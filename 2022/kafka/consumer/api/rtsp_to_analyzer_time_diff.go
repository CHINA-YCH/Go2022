package api

import (
	log "github.com/sirupsen/logrus"
	"time"
)

/*
 * @Author: veno
 * @File: rtsp_to_analyzer_time_diff
 * @Version: ...
 * @Date: 2023-02-28 16:22:07
 * @Description: analyzer 拉取 rtsp文件的时间差
 */
const formatTime = "2006-01-02 15:04:05"

var _ = formatTime
var BigMsgChan = make(chan MsgInfoSource, 8000)
var ChanClosedFlag = false
var LocalTime = time.Now()

// analyzer 产生事件  收到的各个摄像头的事件的第一 ntp_time（已和李俊确认，就算当前摄像头下无车辆，也会吐出数据）
var camAndInfo = make(map[string]MsgInfoSource, 0) // [camId, []信息]

func firstEventTime(m MsgInfoSource) {
	camId := m.CameraId
	// 如果存在该摄像头的数据 说明已经开始处理该视频
	if _, ok := camAndInfo[camId]; ok {
		log.Infof("already exists this cam:%v event", camId)
	} else {
		log.Infof("save cam:%v event", camId)
		camAndInfo[camId] = m
	}
}

func do() {
	// 7 秒之后看所有摄像头第一条数据的时间差 来判断analyzer拉流的时间差
	camArr := []string{"C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "C10", "C11"}
	for i, cam := range camArr {
		if i == len(camArr)-1 {
			return
		}
		preNtp := camAndInfo[cam].Snapshot[0].NtpTime
		preNow := camAndInfo[cam].Snapshot[0].Now
		preTopic := camAndInfo[cam].TopicTimeStamp
		nextNtp := camAndInfo[camArr[i+1]].Snapshot[0].NtpTime
		//nextNow := camAndInfo[camArr[i+1]].Snapshot[0].Now
		//nextTopic := camAndInfo[camArr[i+1]].TopicTimeStamp
		//log.Infof(`
		//		 pre摄像头: %v,  pre产生事件的时间戳: %v,   pre视频打上的时间戳:  %v,  pre topic的时间戳:  %v,
		//        next摄像头: %v, next产生事件的时间戳: %v,  next视频打上的时间戳:  %v, next topic的时间戳:  %v,
		//        ntpTimeDiff: %v ,
		//`, cam, preNow, preNtp, preTopic, camArr[i+1], nextNow, nextNtp, nextTopic, preNtp.Unix()-nextNtp.Unix())
		log.Infof(`摄像头:%v, 产生事件的时间戳:%v, 视频打上的时间戳:%v, send topic的时间戳:%v, ntpTimeDiff:%v, 当前时间:%v`, cam, preNow, preNtp, preTopic, preNtp.Unix()-nextNtp.Unix(), LocalTime)
	}
}

func BatchProcess() { // main concurrent call
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			do()
			ChanClosedFlag = true
			//close(BigMsgChan)
		default:
			select {
			case v, ok := <-BigMsgChan:
				if ok {
					log.Infof("- - -- - - - -")
					firstEventTime(v)
				}
			default:
			}
		}
	}
}

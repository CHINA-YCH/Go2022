package api

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

// MyConsumer 实现  github.com/Shopify/sarama/consumer_group.go/ConsumerGroupHandler  这个接口
type MyConsumer struct {
	File *os.File
}

func (MyConsumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (MyConsumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// checkFileIsExist 检查文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// ConsumeClaim 这个方法用来消费消息的
func (consumer MyConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 获取消息
	for msg := range claim.Messages() {
		//log.Infof("Message topic:%q partition:%d offset:%d, msg key: %v, msg value: %v\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		//log.Infof("\n - - - -consumer timestamp time =%v", time.Now())
		// 写文件
		//_ = readline.Do(string(msg.Value), consumer.File)
		bytes := msg.Value
		//calTimeDiff(bytes)
		calTimeDiffSameNow(bytes)
		// 将消息标记为已使用
		sess.MarkMessage(msg, "")
	}
	return nil
}

var df = make(map[string][]float64, 0)        // <camId,[ totalCount, >2 count, rate(a1/a0)]>
var df2 = make(map[string][]MsgInfoSource, 0) // <now,[msg]>
var df3 = make(map[string]float64, 0)         // <camId, [timedif]>
// 同一now时间 C2到C11 两两之间的时间差
// 1 同一now时间内 C2和C3之间的时间差平均值、C3和C4之间的时间差平均值、C4和C5之间的时间差平均值

func calTimeDiffSameNow(bytes []byte) {
	var msg = MsgInfoSource{}
	_ = json.Unmarshal(bytes, &msg)
	//camId := msg.CameraId
	//ntpTime := msg.Snapshot[0].NtpTime
	now := msg.Snapshot[0].Now.Format("2006-01-02 15:04:05") // now是数据推送出来的时间
	if v, ok := df2[now]; ok {
		df2[now] = append(v, msg)
	} else {
		df2[now] = []MsgInfoSource{msg}
	}

	if len(df2) == 500 { // 2000/10/10 = 20秒左右
		// 查看同一秒内 视频数据的时间差值

		var temp = make(map[string]int64, 0) // <camId, maxTime>
		for k, v := range df2 {              // now 时间戳 对应的msg集合
			for _, v2 := range v {
				cameraId := v2.CameraId
				ntpTime := v2.Snapshot[0].NtpTime
				if maxTimeV, okk := temp[cameraId]; okk {
					if maxTimeV-ntpTime.Unix() < 0 {
						temp[cameraId] = ntpTime.Unix()
					}
				} else {
					temp[cameraId] = ntpTime.Unix()
				}
			}
			// 同一now时间下的各个摄像头数据遍历完成
			var arr = []string{"C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "C10", "C11"}
			for i := 0; i < len(arr); i++ {
				if i+1 >= len(arr) {
					return
				}
				cam := arr[i]
				cam2 := arr[i+1]
				if temp[cam] != 0 && temp[cam2] != 0 {
					// C2 -C3
					dif := math.Abs(float64(temp[cam] - temp[cam2]))
					df3[k+cam+":"+cam2] = dif
				}
			}
		}
		log.Infof(" \n - - - -diff3: %+v", df3)
		time.Sleep(10 * time.Second)
		os.Exit(-1)
	}
	log.Infof(" \n - - - -diff2: %+v", len(df2))
}

func calTimeDiff(bytes []byte) {
	msg := &MsgInfoSource{}
	_ = json.Unmarshal(bytes, &msg)
	id := msg.CameraId
	ntpTime := msg.Snapshot[0].NtpTime
	now := msg.Snapshot[0].Now
	diff := now.Unix() - ntpTime.Unix() // 差值
	//lk.Lock()
	if v, ok := df[id]; ok {
		// 存在该id值
		temp1 := v[0] + 1.0
		temp2 := v[1]
		if diff > 2 {
			temp2 = v[1] + 1.0
		}
		df[id] = []float64{temp1, temp2, temp2 / temp1}
	} else {
		var temp2 float64 = 0.0
		if diff > 2 {
			temp2 = 1.0
		}
		df[id] = []float64{1.0, temp2, temp2 / 1.0}
	}
	//lk.Unlock()
	log.Infof(" \n - - - -diff: %+v", df)
}

// MsgInfoSource Analyzer 消息 消费topic数据结构
type MsgInfoSource struct {
	EventId                string     `json:"eventId"`
	TaskId                 string     `json:"taskId"`
	TaskType               string     `json:"taskType"`
	CameraId               string     `json:"cameraId"`
	Snapshot               []Snapshot `json:"snapshot"`
	OriginalViolationIndex int        `json:"originalViolationIndex"`
	StartTime              time.Time  `json:"startTime"`
	EndTime                time.Time  `json:"endTime"`
	CreatedAt              time.Time  `json:"createdAt"`
	UpdatedAt              time.Time  `json:"updatedAt"`
}

type Snapshot struct {
	SnapshotUriRaw     string    `json:"snapshotUriRaw"`
	NtpTime            time.Time `json:"ntp_time"`
	TrafficLightStatus string    `json:"trafficLightStatus"`
	Class              int64     `json:"class"`
	Now                time.Time `json:"now"`
}

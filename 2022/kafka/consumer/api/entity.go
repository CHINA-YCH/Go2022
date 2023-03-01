package api

import "time"

/*
 * @Author: ych
 * @Description: ...
 * @File: entity
 * @Version: ...
 * @Date: 2023-02-28 15:33:12
 */

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
	TopicTimeStamp         time.Time  `json:"topicTimeStamp"`
}

type Snapshot struct {
	SnapshotUriRaw     string    `json:"snapshotUriRaw"`
	NtpTime            time.Time `json:"ntp_time"`
	TrafficLightStatus string    `json:"trafficLightStatus"`
	Class              int64     `json:"class"`
	Now                time.Time `json:"now"`
	VideoNow           time.Time `json:"video_now"` // 引擎拉取到视频的时间
}

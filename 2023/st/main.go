package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
 * @Author: veno
 * @File: main
 * @Version: ...
 * @Date: 2023-08-17 14:31:36
 * @Description: ...
 */

type OutMsg struct {
	EventId    string          `json:"eventId" bson:"eventId"`     // 事件ID
	EventType  string          `json:"eventType" bson:"eventType"` // 事件类型
	TaskId     string          `json:"taskId" bson:"taskId"`       // 任务id
	TaskType   string          `json:"taskType" bson:"taskType"`   // 任务类型
	ObjectInfo []ObjectInfoStd `json:"object" bson:"object"`       // 轨迹结构
	StartTime  time.Time       `json:"startTime" bson:"startTime"` // 事件起始时间
	EndTime    time.Time       `json:"endTime" bson:"endTime"`     // 事件结束时间
}

type ObjectInfoStd struct {
	Uuid          string  `json:"uuid" bson:"uuid"` // 全局id
	Lane          int32   `json:"lane" bson:"lane"` // 如果是全局的输出，则从中间车道分别是1,2,3,4
	Timestamp     int64   `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Stake         string  `json:"stake" bson:"stake"`                             // 绝对位置，用桩号表示
	StakePosition float32 `json:"stake_position" bson:"stake_position"`           // 绝对位置，数值
	Longitude     float64 `json:"longitude,omitempty" bson:"longitude,omitempty"` // 经度
	Latitude      float64 `json:"latitude,omitempty" bson:"latitude,omitempty"`   // 纬度
	VehicleType   string  `json:"vehicle_type" bson:"vehicle_type"`               // 车辆类型
	Speed         float32 `json:"speed,omitempty" bson:"speed,omitempty"`         // 车辆速度
	CameraId      string  `json:"camera_id" bson:"camera_id"`                     // 图像id对应设备id
	DataSource    int32   `json:"data_source" bson:"data_source"`                 // 1 纯雷达、2 纯视频、3 雷视融合
	Direction     int32   `json:"direction" bson:"direction"`                     // 行车方向 0 上行 1下行
	VehiclePlate  string  `json:"vehicle_plate" bson:"vehicle_plate"`             // 车牌
	VehicleColor  string  `json:"vehicle_color" bson:"vehicle_color"`             // 车辆颜色
	// ----add
	Cog                           float32 `json:"cog" bson:"cog"`                                                           // 航向角（可选）
	SpecialVehicleType            string  `json:"special_vehicle_type" bson:"special_vehicle_type"`                         // 特种车辆类型（可选）
	HazardousChemicalsVehicleType string  `json:"hazardous_chemicals_vehicle_type" bson:"hazardous_chemicals_vehicle_type"` // 危化品车辆类型（可选）：易爆物品、易燃物品、惰性气体、毒性物品、氧化性物品、腐蚀性物品、感染性物品、放射性物品、杂项危险物品
	VehicleAxleType               string  `json:"vehicle_axle_type" bson:"vehicle_axle_type"`                               // 轴型（可选）
	VehiclePlateColor             string  `json:"vehicle_plate_color" bson:"vehicle_plate_color"`                           // 车牌颜色（可选）
	LiveStreamAddress             string  `json:"live_stream_address" bson:"live_stream_address"`                           // 实时流地址 可选
	VehiclePictureAddress         string  `json:"vehicle_picture_address" bson:"vehicle_picture_address"`                   // 车辆图片地址（可选）：可支持卡口抓拍车辆图片展示              //
}

type ObjectInfoIn struct {
	Uid           int64     `json:"uid" bson:"uid"`                       // 图像id
	Uuid          string    `json:"uuid" bson:"uuid"`                     // 全局id
	Lane          int32     `json:"lane" bson:"lane"`                     // 如果是全局的输出，则从中间车道分别是1,2,3,4
	RelPos        float64   `json:"rel_pos" bson:"rel_pos"`               // 隧道接力在使用该字段
	Stake         string    `json:"stake" bson:"stake"`                   //绝对位置，用桩号表示
	StakePosition float32   `json:"stake_position" bson:"stake_position"` //绝对位置，用桩号表示
	Timestamp     int64     `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Box           []float32 `json:"box" bson:"box"`                                 // box
	Box3d         []float32 `json:"box_3d" bson:"box_3d"`                           // 3d box
	Feature       []float32 `json:"feature" bson:"feature"`                         //特征
	VehicleType   string    `json:"vehicle_type" bson:"vehicle_type"`               // 车辆类型
	Location      []float32 `json:"location,omitempty" bson:"location,omitempty"`   //location[0]代表相对雷达的横向距离,location[1]代表相对雷达的纵向距离
	Speed         float32   `json:"speed,omitempty" bson:"speed,omitempty"`         // 车辆速度
	Longitude     float64   `json:"longitude,omitempty" bson:"longitude,omitempty"` // 经度
	Latitude      float64   `json:"latitude,omitempty" bson:"latitude,omitempty"`   // 纬度
	IsVirtual     bool      `json:"is_virtual" bson:"is_virtual"`                   // 是否虚拟
	Cog           float32   `json:"cog" bson:"cog"`
	DebugInfo     string    `json:"debug_info" bson:"debug_info"`   //用于页面debug
	CameraId      string    `json:"camera_id" bson:"camera_id"`     // 图像id对应设备id
	DataSource    int32     `json:"data_source" bson:"data_source"` // 1 纯雷达、2 纯视频、3 雷视融合
	Direction     int32     `json:"direction" bson:"direction"`     // 行车方向 0 上行 1下行
}

func main() {

	in := []ObjectInfoIn{
		{
			Uid:           101,
			Uuid:          "uuid_101",
			Lane:          1,
			Timestamp:     time.Now().Unix(),
			RelPos:        0.89,
			Stake:         "K180+101",
			StakePosition: 729.208,
			Longitude:     32.643287658691406,
			Latitude:      5,
			VehicleType:   "Car",
			Speed:         91.08,
			CameraId:      "camera_id_101",
			DataSource:    2,
			Direction:     0,
			Box:           nil,
		},
	}
	out := []ObjectInfoStd{
		ObjectInfoStd{
			Uuid:          "uuid_101",
			Lane:          1,
			Timestamp:     time.Now().Unix(),
			Stake:         "K180+101",
			StakePosition: 729.208,
			Longitude:     32.643287658691406,
			Latitude:      5,
			VehicleType:   "Car",
			Speed:         91.08,
			CameraId:      "camera_id_101",
			DataSource:    2,
			Direction:     0,
			VehiclePlate:  "浙A60008",
			VehicleColor:  "车辆颜色",
		},
	}

	outMsg := OutMsg{
		EventId:    "event_id_101",
		EventType:  "event_type_101",
		TaskId:     "task_id_101",
		TaskType:   "task_type_101",
		ObjectInfo: out,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
	}
	_ = out
	_ = in
	indent, _ := json.MarshalIndent(outMsg, "", "\t")
	fmt.Println(string(indent))
}

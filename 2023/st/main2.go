package main

import "fmt"

/*
 * @Author: veno
 * @File: main2
 * @Version: ...
 * @Date: 2023-08-22 14:31:58
 * @Description: ...
 */

type ObjectNew struct {
	Uid                           int64     `json:"-" bson:"-"`                                                               // 图像id
	Uuid                          string    `json:"-" bson:"-"`                                                               // 全局id
	Lane                          int32     `json:"lane" bson:"lane"`                                                         // 如果是全局的输出，则从中间车道分别是1,2,3,4
	RelPos                        float64   `json:"rel_pos" bson:"rel_pos"`                                                   // 全局相对位置 隧道接力在使用该字段
	Stake                         string    `json:"stake" bson:"stake"`                                                       // 绝对位置，用桩号表示
	StakePosition                 float32   `json:"stake_position" bson:"stake_position"`                                     // 绝对位置，用桩号表示
	Timestamp                     int64     `json:"timestamp,omitempty" bson:"timestamp,omitempty"`                           // 时间戳
	Box                           []float32 `json:"box" bson:"box"`                                                           // box
	Box3d                         []float32 `json:"box_3d" bson:"box_3d"`                                                     // 3d box
	Feature                       []float32 `json:"feature" bson:"feature"`                                                   // 特征
	VehicleType                   string    `json:"vehicle_type" bson:"vehicle_type"`                                         // 车辆类型
	Location                      []float32 `json:"location,omitempty" bson:"location,omitempty"`                             //location[0]代表相对雷达的横向距离,location[1]代表相对雷达的纵向距离
	Speed                         float32   `json:"speed,omitempty" bson:"speed,omitempty"`                                   // 车辆速度
	Longitude                     float64   `json:"longitude,omitempty" bson:"longitude,omitempty"`                           // 经度
	Latitude                      float64   `json:"latitude,omitempty" bson:"latitude,omitempty"`                             // 纬度
	IsVirtual                     bool      `json:"is_virtual" bson:"is_virtual"`                                             // 是否虚拟
	Cog                           float32   `json:"cog" bson:"cog"`                                                           // 航向
	DebugInfo                     string    `json:"debug_info" bson:"debug_info"`                                             //用于页面debug
	CameraId                      string    `json:"camera_id" bson:"camera_id"`                                               // 图像id对应设备id
	DataSource                    int32     `json:"data_source" bson:"data_source"`                                           // 1 纯雷达、2 纯视频、3 雷视融合
	Direction                     int32     `json:"direction" bson:"direction"`                                               // 行车方向 0 上行 1下行
	VehiclePlate                  string    `json:"vehicle_plate" bson:"vehicle_plate"`                                       // 车牌号
	VehicleColor                  string    `json:"vehicle_color" bson:"vehicle_color"`                                       // 车辆颜色
	SpecialVehicleType            string    `json:"special_vehicle_type" bson:"special_vehicle_type"`                         // 特种车辆类型
	HazardousChemicalsVehicleType string    `json:"hazardous_chemicals_vehicle_type" bson:"hazardous_chemicals_vehicle_type"` // 危化品车辆类型（可选）：易爆物品、易燃物品、惰性气体、毒性物品、氧化性物品、腐蚀性物品、感染性物品、放射性物品、杂项危险物品
	VehicleAxleType               string    `json:"vehicle_axle_type" bson:"vehicle_axle_type"`                               // 轴型
	VehiclePlateColor             string    `json:"vehicle_plate_color" bson:"vehicle_plate_color"`                           // 车牌颜色
	LiveStreamAddress             string    `json:"live_stream_address" bson:"live_stream_address"`                           // 实时流地址
	VehiclePictureAddress         string    `json:"vehicle_picture_address" bson:"vehicle_picture_address"`                   // 车辆图片地址
	VehiclePlatePictureAddress    string    `json:"vehicle_plate_picture_address" bson:"vehicle_plate_picture_address"`       // 车牌图片
}

func main() {
	en := &ObjectNew{
		Uid:  11,
		Uuid: "111",
		Lane: 2,
	}
	fmt.Println(en)
}

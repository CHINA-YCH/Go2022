package main

import (
	"context"
	"fmt"
	write2 "git.supremind.info/gobase/2022/io/read-line/write"
	api2 "git.supremind.info/gobase/2022/kafka/consumer/api"
	"git.supremind.info/gobase/2022/log-d"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"os"
)

type Param struct {
	Topic string
	Host  string
	Group string
}

func init() {
	log_d.SetLog()
}

// 接收数据
func main() {
	file := write2.DoWriteFile("/Users/hanchaoyue/Go2022/Go2022/kafka/data/2498-2/15-16.text")
	Consumer01(file, Param{
		// 10.4.0.113:9094 - tst
		// 10.4.0.113:9095 - log4j 10.4.0.110:21361
		Topic: "nb_2498", // "vehicle_track", BOX.EVENT_VEHICLE_MODEL 100.100.142.15:32449 ; 10.4.0.113:9094 jtsj_2498;
		Host:  "10.4.0.113:9094",
		Group: "anatracer-214",
	})
}

func Consumer01(file *os.File, cf Param) {
	// 先初始化 kafka
	config := sarama.NewConfig()
	// Version 必须大于等于  V0_10_2_0
	config.Version = sarama.V0_10_2_1
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	log.Infof("start connect kafka . . . . . . ")
	// 开始连接kafka服务器
	group, err := sarama.NewConsumerGroup([]string{cf.Host}, cf.Group, config)
	if err != nil {
		log.Errorf("connect kafka failed %v", err)
		return
	}
	// 检查错误
	go func() {
		for err := range group.Errors() {
			log.Errorf("group errors %v", err)
		}
	}()
	ctx := context.Background()
	log.Infof("start get msg . . . . . . ")
	// for 是应对 consumer rebalance
	for {
		// 需要监听的主题
		topics := []string{cf.Topic}
		handler := api2.MyConsumer{File: file}
		// 启动kafka消费组模式，消费的逻辑在上面的 ConsumeClaim 这个方法里
		err = group.Consume(ctx, topics, handler)
		if err != nil {
			fmt.Println("consume failed; err : ", err)
			return
		}
	}
}

package main

import (
	"context"
	"fmt"
	"git.supremind.info/gobase/kafka/consumer"
	logd "git.supremind.info/gobase/log-d"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type Param struct {
	Topic string
	Host  string
	Group string
}

func init() {
	logd.SetLog()
}

// 接收数据
func main() {
	Consumer01(Param{
		Topic: "BOX.EVENT_VEHICLE_MODEL", // "vehicle_track",
		Host:  "100.100.142.15:32449",
		Group: "track003",
	})
}

func Consumer01(cf Param) {
	// 先初始化 kafka
	config := sarama.NewConfig()
	// Version 必须大于等于  V0_10_2_0
	config.Version = sarama.V0_10_2_1
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
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
		handler := consumer.MyConsumer{}
		// 启动kafka消费组模式，消费的逻辑在上面的 ConsumeClaim 这个方法里
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			fmt.Println("consume failed; err : ", err)
			return
		}
	}
}

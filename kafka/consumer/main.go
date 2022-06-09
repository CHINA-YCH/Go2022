package main

import (
	"context"
	"fmt"
	logd "git.supremind.info/gobase/log-d"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	topic = "vehicle_record_topic"
	host  = "10.4.0.113:9095"
	group = "local06"
)

func init() {
	logd.SetLog()
}

// 接收数据
func main() {
	// 先初始化 kafka
	config := sarama.NewConfig()
	// Version 必须大于等于  V0_10_2_0
	config.Version = sarama.V0_10_2_1
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	log.Infof("start connect kafka . . . . . . ")
	// 开始连接kafka服务器
	group, err := sarama.NewConsumerGroup([]string{host}, group, config)
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
		topics := []string{topic}
		handler := MyConsumer{}
		// 启动kafka消费组模式，消费的逻辑在上面的 ConsumeClaim 这个方法里
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			fmt.Println("consume failed; err : ", err)
			return
		}
	}
}

// MyConsumer 实现  github.com/Shopify/sarama/consumer_group.go/ConsumerGroupHandler  这个接口
type MyConsumer struct {
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
		log.Infof("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		log.Infof("msg key: %v\n", string(msg.Key))
		log.Infof("msg value: %v\n", string(msg.Value))
		// 将消息标记为已使用
		sess.MarkMessage(msg, "")
	}
	return nil
}

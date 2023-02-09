package api

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
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
		log.Infof("\n - - - -consumer timestamp time =%v", time.Now())
		// 写文件
		//_ = readline.Do(string(msg.Value), consumer.File)
		// 将消息标记为已使用
		sess.MarkMessage(msg, "")
	}
	return nil
}

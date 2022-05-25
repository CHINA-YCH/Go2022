package main

import (
	"fmt"
	"git.supremind.info/gobase/io/read-line"
	"github.com/Shopify/sarama"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

var (
	host  = "127.0.0.1:9092"
	topic = "jtsj_2498_20"
	path  = "kafka/data/2498/2498_20.text"
)

func main() {
	line, err := read_line.ReadLine(path)
	if err != nil {
		log.Println("read error:", err)
		panic(nil)
	}
	for _, v := range line {
		Producer2(v)
	}
}
func Producer2(value string) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{host}, config)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(producer)
	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(10),
		Key:       sarama.StringEncoder("key"),
		Value:     sarama.ByteEncoder(value),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Send message Fail")
	}
	fmt.Printf(",time = %v,Partition = %d, offset=%d, value=%v \n", time.Now(), partition, offset, value)
}

func Producer(flag bool, fileName string) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{host}, config)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(producer)
	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		//Topic: "test",//包含了消息的主题
		Partition: int32(10),                   //
		Key:       sarama.StringEncoder("key"), //
	}
	value := getValue(flag, topic, fileName)
	var msgType = "string"
	fmt.Println("msgType = ", msgType, ",value = ", value)
	msg.Topic = topic
	//将字符串转换为字节数组
	msg.Value = sarama.ByteEncoder(value)
	//SendMessage：该方法是生产者生产给定的消息
	//生产成功的时候返回该消息的分区和所在的偏移量
	//生产失败的时候返回error
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Send message Fail")
	}
	fmt.Printf(",time = %v,Partition = %d, offset=%d\n", time.Now(), partition, offset)
}

var i = 0

func getValue(flag bool, topic string, fileName string) string {
	if flag {
		i = i + 1
		return strconv.FormatInt(int64(i), 10)
	} else {
		filename := fileName
		configData, _ := ioutil.ReadFile(filename)
		var value = string(configData)
		return value
	}
}

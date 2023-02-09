package main

import (
	"fmt"
	"git.supremind.info/gobase/io/read-line/read"
	"git.supremind.info/gobase/kafka/producer/api"
	"github.com/Shopify/sarama"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"time"
)

var (
	// ""100.100.142.15:32449 10.7.0.16:9092 100.100.142.232:9092 100.100.142.15:32449 100.100.152.232:9092 10.4.0.113:9094 jtsj_2498
	//host  = "100.100.142.132:9092"
	//topic = "BOX.EVENT_VEHICLE_MODEL"
	// path  = "/Users/hanchaoyue/Go2022/Go2022/kafka/data/zdcl/004.txt"
	host  = "100.100.142.177:9092"
	topic = "jtcs_2740"
	path  = "/Users/hanchaoyue/Go2022/Go2022/kafka/data/2740/2740.txt"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 100; i++ {
		Pro3()
	}
}

func Pro3() {
	msgSink, err := api.NewKafkaSink(api.SinkKafkaConfig{
		Brokers:           []string{host},
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		log.Println("log=", err.Error())
	}

	exec := api.NewExec(msgSink)
	line, err := read.ReadLine(path)
	if err != nil {
		log.Println("read error:", err)
		panic(nil)
	}
	for _, v := range line {
		time.Sleep(1 * time.Second)
		exec.MsgProcess(v)
	}

}
func Pro2() {
	line, err := read.ReadLine(path)
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

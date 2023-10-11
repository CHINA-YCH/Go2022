package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	channel := "hello"
	go subMessage(channel) // 消息订阅节点
	msgList := []string{"hello", "world", "redis pub", "redis sub"}
	// 此处发了两个消息
	for _, msg := range msgList {
		time.Sleep(2 * time.Second)
		pubMessage(channel, msg) // 消息发布节点
		fmt.Printf("已经发送msg: %s ,到 channel:%s\n", msg, channel)
	}
	time.Sleep(100000 * time.Second)

}

func redisConnect() (rdb *redis.Client) {
	var (
		redisServer string
		port        string
		password    string
	)
	redisServer = os.Getenv("RedisUrl")
	redisServer = "localhost"
	port = os.Getenv("RedisPort")
	password = os.Getenv("RedisPass")
	port = "6379"
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisServer + ":" + port,
		Password: password,
		DB:       0, // use default DB
	})

	return
}

func pubMessage(channel, msg string) {
	rdb := redisConnect()
	rdb.Publish(context.Background(), channel, msg)
}

func subMessage(channel string) {
	rdb := redisConnect()
	pubsub := rdb.Subscribe(context.Background(), channel)
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		panic(err)
	}

	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println("订阅 hell channel msg: ", msg.Channel, msg.Payload)
	}
}

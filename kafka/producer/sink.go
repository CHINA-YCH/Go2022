package producer

import (
	"context"
	"encoding/json"
	"fmt"
)

var SinkInit MsgSink
var sinkClient Sink

func NewKafkaSink(config SinkKafkaConfig) (MsgSink, error) {
	producer, err := NewKafkaClient(config)
	sinkClient = Sink{
		KafkaClient:  producer,
		DefaultTopic: config.Topic,
		Config:       config,
	}
	SinkInit = &sinkClient
	return &sinkClient, err
}

type Sink struct {
	KafkaClient  *KafkaClient
	DefaultTopic string
	Config       SinkKafkaConfig
}

func (s *Sink) Push(_ context.Context, msg interface{}) error {
	msgStr, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("KafkaSink json.Marshal error: %v", err)
	}
	err = s.KafkaClient.SendMessage(s.DefaultTopic, string(msgStr))
	return err
}

func (s *Sink) PushList(_ context.Context, msgs []string) error {
	_ = s.KafkaClient.SendMessages(s.DefaultTopic, msgs)
	return nil
}

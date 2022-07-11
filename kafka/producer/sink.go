package producer

import (
	"context"
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

func (s *Sink) Push(_ context.Context, msg string) error {
	err := s.KafkaClient.SendMessage(s.DefaultTopic, msg)
	return err
}

func (s *Sink) PushList(_ context.Context, msgs []string) error {
	_ = s.KafkaClient.SendMessages(s.DefaultTopic, msgs)
	return nil
}

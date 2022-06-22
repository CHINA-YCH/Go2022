package producer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"time"
)

type SinkKafkaConfig struct {
	Brokers           []string `json:"brokers" yaml:"brokers"`
	Topic             string   `json:"topic" yaml:"topic"`
	NumPartitions     int      `json:"num_partitions" yaml:"num_partitions"`
	ReplicationFactor int      `json:"replication_factor" yaml:"replication_factor"`
}
type KafkaClient struct {
	config   SinkKafkaConfig
	producer sarama.SyncProducer
}

func NewKafkaClient(clientConfig SinkKafkaConfig) (*KafkaClient, error) {
	if clientConfig.NumPartitions <= 0 {
		clientConfig.NumPartitions = 1
	}
	if clientConfig.ReplicationFactor <= 0 {
		clientConfig.ReplicationFactor = 1
	}
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Timeout = 5 * time.Second
	//init topics
	cAdmin, err := sarama.NewClusterAdmin(clientConfig.Brokers, kafkaConfig)
	if err != nil {
		log.Errorf("kafka producer new ClusterAdmin error: %s", err.Error())
		return nil, err
	}
	defer func(cAdmin sarama.ClusterAdmin) {
		err := cAdmin.Close()
		if err != nil {
			log.Errorf("close cluster admin error: %s", err.Error())
		}
	}(cAdmin)
	// broker
	topicMetas, err := cAdmin.DescribeTopics([]string{clientConfig.Topic})
	if err != nil {
		log.Errorf("kafka producer DescribeTopics error: %s", err.Error())
		return nil, err
	}
	for _, meta := range topicMetas {
		if meta.Err == sarama.ErrUnknownTopicOrPartition {
			log.Infof("kafka producer doesn't exited, Creating topic: %s", meta.Name)
			err = cAdmin.CreateTopic(meta.Name, &sarama.TopicDetail{
				NumPartitions:     int32(clientConfig.NumPartitions),
				ReplicationFactor: int16(clientConfig.ReplicationFactor),
			}, false)
			if err != nil {
				log.Errorf("CreateTopic error: %s.", err.Error())
				return nil, err
			}
		} else if len(meta.Partitions) < clientConfig.NumPartitions {
			err = cAdmin.CreatePartitions(meta.Name, int32(clientConfig.NumPartitions), nil, false)
			if err != nil {
				log.Warnf("CreatePartitions error: %s.", err.Error())
			}
		}
	}

	p, err := sarama.NewSyncProducer(clientConfig.Brokers, kafkaConfig)
	if err != nil {
		log.Errorf("kafka producer NewSyncProducer error: %s", err.Error())
		return nil, err
	}
	client := KafkaClient{
		producer: p,
		config:   clientConfig,
	}
	return &client, nil
}

func (c *KafkaClient) Close() error {
	return c.producer.Close()
}

func (c *KafkaClient) SendMessage(topic string, data string) error {
	msg := &sarama.ProducerMessage{
		Timestamp: time.Now(),
		Topic:     topic,
		Value:     sarama.ByteEncoder(data),
	}
	part, offset, err := c.producer.SendMessage(msg)
	if err != nil {
		marshal, _ := json.Marshal(data)
		log.Errorf("kafka producer send message error: %s, message:(%s)", err.Error(), string(marshal))
		return err
	}
	log.Infof("kafka producer send to kafka successfully，topic=%s, partition=%d, offset=%d, len(data)=%d", topic, part, offset, len(data))
	return nil
}

func (c *KafkaClient) SendMessageWithKey(topic, key, data string) error {
	msg := &sarama.ProducerMessage{
		Timestamp: time.Now(),
		Topic:     topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(data),
	}
	part, offset, err := c.producer.SendMessage(msg)
	if err != nil {
		marshal, _ := json.Marshal(data)
		log.Errorf("kafka producer send message error: %s, message:(%s)", err.Error(), string(marshal))
		return err
	}
	log.Infof("kafka producer send to kafka with key successfully，topic=%s, key=%s, partition=%d, offset=%d, len(data)=%d", topic, key, part, offset, len(data))
	return nil
}

func (c *KafkaClient) SendMessages(topic string, dataList []string) error {
	var msgs []*sarama.ProducerMessage
	for _, data := range dataList {
		msgs = append(msgs, &sarama.ProducerMessage{
			Timestamp: time.Now(),
			Topic:     topic,
			Value:     sarama.ByteEncoder(data),
		})
	}
	log.Infof("kafka producer has already send num: %d", len(msgs))
	err := c.producer.SendMessages(msgs)
	if err != nil {
		log.Errorf("kafka producer send message list error: %s", err.Error())
		return err
	}
	return nil
}

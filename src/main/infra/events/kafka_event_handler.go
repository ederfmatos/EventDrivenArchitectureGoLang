package events

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"
)

type KafkaEventHandler struct {
	configMap *kafka.ConfigMap
}

func NewKafkaEventHandler(server, groupId string) *KafkaEventHandler {
	return &KafkaEventHandler{
		configMap: &kafka.ConfigMap{
			"bootstrap.servers": server,
			"group.id":          groupId,
		},
	}
}

func (handler *KafkaEventHandler) Consume(topic string, messageHandler func(message []byte) error) {
	consumer, err := kafka.NewConsumer(handler.configMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Info().Str("topic", *msg.TopicPartition.Topic).Msg("Message received from kafka")
			err = messageHandler(msg.Value)
			if err != nil {
				log.Err(err).Str("topic", *msg.TopicPartition.Topic).Msg("Error on consume message from kafka")
			}
		}
	}
}

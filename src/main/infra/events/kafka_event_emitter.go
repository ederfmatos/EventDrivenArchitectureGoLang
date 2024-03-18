package events

import (
	"EventDrivenArchitectureGoLang/src/main/domain/event"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"
)

type KafkaEventEmitter struct {
	ConfigMap *kafka.ConfigMap
}

func NewKafkaEventEmitter(server string) *KafkaEventEmitter {
	return &KafkaEventEmitter{
		ConfigMap: &kafka.ConfigMap{
			"bootstrap.servers":   server,
			"delivery.timeout.ms": "0",
			"enable.idempotence":  "true",
		},
	}
}

func (eventEmitter *KafkaEventEmitter) Emit(event event.Event) error {
	producer, err := kafka.NewProducer(eventEmitter.ConfigMap)
	if err != nil {
		return err
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		return err
	}

	topic := event.GetName()
	key := event.GetId()
	log.Info().Str("topic", topic).Str("key", key).Msg("Sending message to kafka")
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          eventJson,
		Key:            []byte(key),
	}
	return producer.Produce(message, nil)
}

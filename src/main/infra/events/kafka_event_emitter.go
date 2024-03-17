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

func NewKafkaEventEmitter(server, groupId string) *KafkaEventEmitter {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          groupId,
	}
	return &KafkaEventEmitter{ConfigMap: configMap}
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
	log.Info().Str("topic", topic).Str("key", key).Any("payload", event.GetPayload()).Msg("Sending message o kafka")
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          eventJson,
		Key:            []byte(key),
	}
	return producer.Produce(message, nil)
}

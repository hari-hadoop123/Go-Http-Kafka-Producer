package util

import "time"

func getKafkaClient(kafkaBrokers []string) (sarama.AsyncProducer, error) {

	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 200 * time.Millisecond // Flush batches every 500ms
	config.Producer.MaxMessageBytes = 10000 * 1024
	prd, err := sarama.NewAsyncProducer(kafkaBrokers, config)

	return prd, err
}

func PublishToKafka(kafkaTopic string, msgData []byte) (err error) {
	producer, err := getKafkaClient(kafkaBrokers)

	if err != nil {
		log.Error("Error while creating kafka producer :", err)
		return
	}
	defer producer.Close()

	if kafkaTopic != "" {
		record := &sarama.ProducerMessage{Topic: kafkaTopic, Value: sarama.ByteEncoder(msgData)}

		producer.Input() <- record
	}

	return
}

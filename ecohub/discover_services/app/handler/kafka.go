package handler

import (
        "net/http"
	"fmt"
        "github.com/Shopify/sarama"
)

func kafka_connect(kafkaHost string, kafkaPort string, ProducerTopic string, ConsumerTopic string) (int, bool, bool){
        producerTopicFound := false
        consumerTopicFound := false
        config := sarama.NewConfig()
        url := kafkaHost + ":" + kafkaPort
        brokers := []string{url}
        cluster, err := sarama.NewConsumer(brokers, config)
	fmt.Println("cluster", cluster)
        if err != nil {
                logs.Errorf("Failed to connect kafka host")
                return http.StatusNotFound, producerTopicFound, consumerTopicFound
        }
        topics, _ := cluster.Topics()
        for index := range topics {
                if ProducerTopic == topics[index] {
                        producerTopicFound = true
                }
                if ConsumerTopic == topics[index] {
                        consumerTopicFound = true
                }
       	}
	return http.StatusOK, producerTopicFound, consumerTopicFound
}

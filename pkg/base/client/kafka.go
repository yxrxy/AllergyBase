/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	kafukago "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/yxrxy/AllergyBase/config"
	"github.com/yxrxy/AllergyBase/pkg/errno"
)

const (
	Timeout                       = 3 * time.Second
	DefaultReaderGroupID          = "default_reader_group"
	KafkaReadMaxBytes             = 10e6
	KafkaRetries                  = 3
	DefaultKafkaNumPartitions     = 1
	DefaultKafkaReplicationFactor = 1
	ConsumerOffsetsPartitions     = 50 // Kafka 默认值
)

// GetConn conn不能保证并发安全,仅可作为单线程的长连接使用。
func GetConn() (*kafukago.Conn, error) {
	dialer := getDialer()
	conn, err := dialer.Dial(config.Kafka.Network, config.Kafka.Address)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalKafkaErrorCode, fmt.Sprintf("failed dial kafka server,error: %v", err))
	}
	return conn, nil
}

// GetNewReader 创建一个reader示例，reader是并发安全的
func GetNewReader(topic string, groupID string) *kafukago.Reader {
	if groupID == "" {
		groupID = DefaultReaderGroupID
	}

	// 确保 __consumer_offsets 主题存在
	if err := createConsumerOffsetsTopic(); err != nil {
		log.Printf("警告: 创建 consumer offsets 主题失败: %v", err)
	}

	if err := createIfNotExist(topic); err != nil {
		log.Printf("Failed to create topic %s: %v", topic, err)
		return nil
	}

	cfg := kafukago.ReaderConfig{
		Brokers:     []string{config.Kafka.Address},
		Topic:       topic,
		GroupID:     groupID,
		MaxBytes:    KafkaReadMaxBytes,
		MaxAttempts: KafkaRetries,
		ErrorLogger: log.New(log.Writer(), "[Kafka Reader] ", log.LstdFlags),
		Dialer:      getDialer(),
		StartOffset: kafukago.FirstOffset, // 从最早的消息开始读取
	}
	return kafukago.NewReader(cfg)
}

func createConsumerOffsetsTopic() error {
	conn, err := GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.CreateTopics(kafukago.TopicConfig{
		Topic:             "__consumer_offsets",
		NumPartitions:     ConsumerOffsetsPartitions,
		ReplicationFactor: 1, // 单节点使用 1
	})
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
	}
	return nil
}

// GetNewWriter 创建一个writer示例，writer是并发安全的。
func GetNewWriter(topic string, async bool) (*kafukago.Writer, error) {
	if err := createIfNotExist(topic); err != nil {
		return nil, err
	}

	addr, err := net.ResolveTCPAddr(config.Kafka.Network, config.Kafka.Address)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalNetworkErrorCode, fmt.Sprintf("failed create kafka writer,error: %v", err))
	}

	return &kafukago.Writer{
		Addr:                   addr,
		Topic:                  topic,
		Balancer:               &kafukago.RoundRobin{},
		MaxAttempts:            KafkaRetries,
		RequiredAcks:           kafukago.RequireOne,
		Async:                  async,
		AllowAutoTopicCreation: true,
		ErrorLogger:            log.New(log.Writer(), "[Kafka Writer] ", log.LstdFlags),
		Transport:              getTransport(),
	}, nil
}

func createIfNotExist(topic string) error {
	conn, err := GetConn()
	if err != nil {
		return err
	}

	err = conn.CreateTopics(kafukago.TopicConfig{
		Topic:             topic,
		NumPartitions:     DefaultKafkaNumPartitions,
		ReplicationFactor: DefaultKafkaReplicationFactor,
	})
	if err != nil {
		return errno.NewErrNo(errno.InternalKafkaErrorCode, fmt.Sprintf("failed to create topic, err: %v", err))
	}
	return nil
}

func getDialer() *kafukago.Dialer {
	mechanism := plain.Mechanism{
		Username: config.Kafka.User,
		Password: config.Kafka.Password,
	}
	return &kafukago.Dialer{
		Timeout:       Timeout,
		DualStack:     true,
		SASLMechanism: mechanism,
	}
}

func getTransport() *kafukago.Transport {
	mechanism := plain.Mechanism{
		Username: config.Kafka.User,
		Password: config.Kafka.Password,
	}
	return &kafukago.Transport{
		SASL: mechanism,
	}
}

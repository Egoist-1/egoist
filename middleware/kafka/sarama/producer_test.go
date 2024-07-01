package sarama

import (
	"github.com/IBM/sarama"
	"testing"
)

var addrs = []string{"localhost:9094"}

func TestSyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
}

type name struct {
	Data any
}

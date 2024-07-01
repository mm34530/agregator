package publisher

import (
	"agregator/producer"
	"math/rand"
	"time"

	"github.com/msales/streams/v6"
)

type Publisher struct {
	producer     Producer
	messagesChan chan map[string]interface{}
}

type Producer interface {
	ProduceMessage(msgType string, source int) interface{}
}

func New(producer Producer) *Publisher {
	return &Publisher{
		producer:     producer,
		messagesChan: make(chan map[string]interface{}, 1000),
	}
}

func (p *Publisher) Publish(msgCount int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < msgCount; i++ {
		randomInt := rand.Intn(2)
		source := rand.Intn(1000)

		switch randomInt {
		case 0:
			msg := map[string]interface{}{
				producer.BussinesPartnerAType: p.producer.ProduceMessage(producer.BussinesPartnerAType, source),
			}

			p.messagesChan <- msg
		case 1:
			msg := map[string]interface{}{
				producer.BussinesPartnerBType: p.producer.ProduceMessage(producer.BussinesPartnerBType, source),
			}

			p.messagesChan <- msg
		}
	}
}

func (s *Publisher) Consume() (streams.Message, error) {
	msg := <-s.messagesChan

	var key string
	var value interface{}

	for k, v := range msg {
		key = k
		value = v
	}

	return streams.NewMessage(key, value), nil
}

func (s *Publisher) Commit(v interface{}) error {
	return nil
}

func (s *Publisher) Close() error {
	return nil
}

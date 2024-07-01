package publisher

type Publisher struct {
	producer Producer
}

type Producer interface {
	ProduceMessage(msgType string) interface{}
}

func New(producer Producer) *Publisher {
	return &Publisher{
		producer: producer,
	}
}

func (p *Publisher) Publish(msgCount int) {
}

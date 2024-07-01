package producer

type Producer struct {
}

func New() *Producer {
	return &Producer{}
}

func (p *Producer) ProduceMessage(msgType string) interface{} {
	return nil
}

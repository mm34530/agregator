package producer

import "log"

type Producer struct {
	generate *Generate
}

func New() *Producer {
	generate := NewGenerator()

	return &Producer{
		generate: generate,
	}
}

func (p *Producer) ProduceMessage(msgType string, source int) interface{} {
	msg, err := p.generate.GenerateMessage(msgType, 1)
	if err != nil {
		log.Println(err)

		return nil
	}

	return msg
}

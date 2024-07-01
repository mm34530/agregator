package receiver

import (
	"agregator/producer"
	"errors"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/msales/streams/v6"
)

type Receiver struct {
	m     sync.Mutex
	stats map[string]int
}

func New() *Receiver {
	return &Receiver{
		stats: make(map[string]int),
	}
}

func (r *Receiver) ProcessMessage(msgType string, msg interface{}) error {
	r.m.Lock()
	defer r.m.Unlock()

	switch msgType {
	case producer.BussinesPartnerAType:
		bp, ok := msg.(producer.BussinesPartnerA)
		if !ok {
			return errors.New("Error: invalid message type %T")
		}

		cratedAt := time.Unix(int64(bp.CreatedAt), 0)
		resolvedAt := time.Unix(int64(bp.ResolvedAt), 0)

		field := fmt.Sprintf("%s-%v-%d", bp.Country, bp.SourceID, roundUpDuration(resolvedAt.Sub(cratedAt)))

		if resolvedAt.Sub(cratedAt) > 23*time.Hour {
			log.Printf("Warning: Resolution time is higher than 23 hours for message %v", bp)
		}

		r.stats[field]++

	case producer.BussinesPartnerBType:
		bp, ok := msg.(producer.BussinesPartnerB)
		if !ok {
			return errors.New("Error: invalid message type %T")
		}

		duration := time.Duration(bp.Processing.Duration) * time.Second

		field := fmt.Sprintf("%s-%v-%d", bp.Origin.Geo, bp.Origin, roundUpDuration(duration))

		if duration > 23*time.Hour {
			log.Printf("Warning: Resolution time is higher than 23 hours for message %v", bp)
		}

		r.stats[field]++

	default:
		return errors.New("Error: invalid message type")
	}

	return nil
}

func roundUpDuration(d time.Duration) int {
	return int(math.Ceil(d.Hours()))
}

func (r *Receiver) GetStats() map[string]int {
	r.m.Lock()
	defer r.m.Unlock()
	return r.stats
}

func (r *Receiver) WithPipe(pipe streams.Pipe) {
}

func (r *Receiver) Process(msg streams.Message) error {
	return r.ProcessMessage(msg.Key.(string), msg.Value)
}

func (r *Receiver) Close() error {
	return nil
}

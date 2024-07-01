package receiver

import (
	"agregator/producer"
	"fmt"
	"log"
	"math"
	"sync"
	"time"
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

func (r *Receiver) ProcessMessage(msgType string, msg interface{}) {
	r.m.Lock()
	defer r.m.Unlock()

	switch msgType {
	case "BussinesPartnerA":
		bp, ok := msg.(producer.BussinesPartnerA)
		if !ok {
			log.Printf("Error: invalid message type %T", msg)

			return
		}

		cratedAt := time.Unix(int64(bp.CreatedAt), 0)
		resolvedAt := time.Unix(int64(bp.ResolvedAt), 0)

		field := fmt.Sprintf("%s-%v-%d", bp.Country, bp.SourceID, roundUpDuration(resolvedAt.Sub(cratedAt)))

		if resolvedAt.Sub(cratedAt) > 23*time.Hour {
			log.Printf("Warning: Resolution time is higher than 23 hours for message %v", bp)
		}

		r.stats[field]++

	case "BussinesPartnerB":
		bp, ok := msg.(producer.BussinesPartnerB)
		if !ok {
			log.Printf("Error: invalid message type %T", msg)

			return
		}

		duration := time.Duration(bp.Processing.Duration) * time.Second

		field := fmt.Sprintf("%s-%v-%d", bp.Origin.Geo, bp.Origin, roundUpDuration(duration))

		if duration > 23*time.Hour {
			log.Printf("Warning: Resolution time is higher than 23 hours for message %v", bp)
		}

		r.stats[field]++
	}
}

func roundUpDuration(d time.Duration) int {
	return int(math.Ceil(d.Hours()))
}

func (r *Receiver) GetStats() map[string]int {
	r.m.Lock()
	defer r.m.Unlock()
	return r.stats
}

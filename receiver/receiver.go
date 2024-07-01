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
		bp := msg.(producer.BussinesPartnerA)
		for _, sourceID := range bp.SourceID {
			field := fmt.Sprintf("%s-%s-%d", bp.Country, sourceID, roundUpDuration(bp.ResolvedAt.Sub(bp.CreatedAt)))

			if bp.ResolvedAt.Sub(bp.CreatedAt) > 23*time.Hour {
				log.Printf("Warning: Resolution time is higher than 23 hours for message %v", bp)
			}

			r.stats[field]++
		}

	case "BussinesPartnerB":
		bp := msg.(producer.BussinesPartnerB)
		for _, owner := range bp.Origin.Owner {
			field := fmt.Sprintf("%s-%s-%d", bp.Origin.Geo, owner, roundUpDuration(bp.Processing.Duration))

			if bp.Processing.Duration > 23*time.Hour {
				log.Printf("Warning: Resolution time is higher than 23 hours for message %v", bp)
			}

			r.stats[field]++
		}
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

package producer

import "time"

const (
	BussinesPartnerAType = "bussines_partner_a"
	BussinesPartnerBType = "bussines_partner_b"
)

type BussinesPartnerA struct {
	ID         string    `json:"id"`
	SourceID   []string  `json:"source_id"`
	Country    string    `json:"country"`
	CreatedAt  time.Time `json:"created_at"`
	ResolvedAt time.Time `json:"resolved_at"`
}

type BussinesPartnerB struct {
	TaskID     int        `json:"task_id"`
	Origin     Origin     `json:"origin"`
	Processing Processing `json:"processing"`
}

type Origin struct {
	Owner []string `json:"owner"`
	Geo   string   `json:"geo"`
}

type Processing struct {
	At       time.Time     `json:"at"`
	Duration time.Duration `json:"duration"`
}

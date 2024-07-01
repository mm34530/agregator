package producer

const (
	BussinesPartnerAType = "bussines_partner_a"
	BussinesPartnerBType = "bussines_partner_b"
)

type BussinesPartnerA struct {
	ID         string `json:"id"`
	SourceID   int    `json:"source_id"`
	Country    string `json:"country"`
	CreatedAt  int    `json:"created_at"`
	ResolvedAt int    `json:"resolved_at"`
}

type BussinesPartnerB struct {
	TaskID     int        `json:"task_id"`
	Origin     Origin     `json:"origin"`
	Processing Processing `json:"processing"`
}

type Origin struct {
	Owner int    `json:"owner"`
	Geo   string `json:"geo"`
}

type Processing struct {
	At       int `json:"at"`
	Duration int `json:"duration"`
}

package producer

import (
	"errors"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

type Generate struct {
	sequentialNumber atomic.Int32
}

func NewGenerator() *Generate {
	return &Generate{
		sequentialNumber: atomic.Int32{},
	}
}

func (g *Generate) GenerateMessage(msgType string, source int) (interface{}, error) {
	switch msgType {
	case BussinesPartnerAType:
		return g.GenerateBussinesAPartnerAData(source), nil
	case BussinesPartnerBType:
		return g.GenerateBussinesBPartnerBData(source), nil
	default:
		return nil, errors.New("invalid message type")
	}
}

func (g *Generate) GenerateBussinesAPartnerAData(source int) BussinesPartnerA {
	id := uuid.New().String()
	countryCode := g.GenerateCountry()
	createdAt, resolvedAt := g.GenerateRandomDates()

	return BussinesPartnerA{
		ID:         id,
		SourceID:   source,
		Country:    countryCode,
		CreatedAt:  createdAt,
		ResolvedAt: resolvedAt,
	}
}

func (g *Generate) GenerateBussinesBPartnerBData(source int) BussinesPartnerB {
	g.sequentialNumber.Add(1)

	createdAt, resolvedAt := g.GenerateRandomDates()

	createdAtTime := time.Unix(int64(createdAt), 0)
	resolvedAtTime := time.Unix(int64(resolvedAt), 0)

	duration := time.Duration(rand.Intn(int(resolvedAtTime.Sub(createdAtTime).Seconds())))

	return BussinesPartnerB{
		TaskID:     int(g.sequentialNumber.Load()),
		Origin:     Origin{Owner: source, Geo: g.GenerateCountry()},
		Processing: Processing{At: createdAt, Duration: int(duration.Seconds())},
	}
}

func (g *Generate) GenerateCountry() string {
	var countryCodes = []string{
		"AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG",
		"AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB",
		"BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BA", "BW", "BV",
		"BR", "IO", "BN", "BG", "BF", "BI", "CV", "KH", "CM", "CA",
		"KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG",
		"CD", "CK", "CR", "CI", "HR", "CU", "CY", "CZ", "DK", "DJ",
		"DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "ET", "FK",
		"FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE",
		"DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG",
		"GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS",
		"IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP",
		"JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA",
		"LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MK",
		"MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU",
		"YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ",
		"MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG",
		"NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG",
		"PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO",
		"RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS",
		"SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK",
		"SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR",
		"SJ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG",
		"TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA",
		"AE", "GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG",
		"VI", "WF", "EH", "YE", "ZM", "ZW",
	}

	rand.Seed(time.Now().UnixNano())

	return countryCodes[rand.Intn(len(countryCodes))]
}

func (g *Generate) GenerateRandomDates() (int, int) {
	now := time.Now()

	createdAt := now.Add(-time.Duration(rand.Intn(24*60*60)) * time.Second)

	resolvedAt := createdAt.Add(time.Duration(rand.Intn(int(now.Sub(createdAt).Seconds()))) * time.Second)
	return int(createdAt.Unix()), int(resolvedAt.Unix())
}

package receiver_test

import (
	"agregator/producer"
	"agregator/receiver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReceiveMessage(t *testing.T) {
	t.Run("Should correctly calculate stats", func(t *testing.T) {
		messages := []producer.BussinesPartnerA{
			{"926b3c49-1cd0-4a6e-b4b5-9f68fcb01716", 101, "US", 1645615000, 1645615500},
			{"0c5b76f3-bb3c-45f2-8f47-1f76ff7c04c5", 102, "DE", 1645681000, 1645682000},
			{"b8216b2c-d62f-4c9c-bf4d-36a632f0a93c", 101, "US", 1645648700, 1645649500},
			{"dbbf98aa-1e4d-4d5e-956b-5d2761a6d589", 103, "FR", 1645670200, 1645672900},
			{"a78801d0-2a4c-4d3c-8b5a-4e76b1c1c31d", 102, "DE", 1645663100, 1645664500},
			{"5fbd8f1e-231f-4d7f-9a6d-19bf52b1285f", 103, "FR", 1645666000, 1645668000},
			{"b96a2e77-55a6-41d6-a9b8-6e98d8937926", 101, "US", 1645629300, 1645632000},
			{"f51a54a8-7ed2-47e8-87a6-04dbcc77ea1f", 102, "DE", 1645677400, 1645682200},
			{"e2d5ceeb-3c62-49ab-a146-6c1e19bbd378", 103, "FR", 1645659300, 1645660700},
			{"b7e3546d-5565-489e-aa14-af6e4b4e872c", 101, "US", 1645618600, 1645620900},
		}
		receiver := receiver.New()

		for _, message := range messages {
			receiver.ProcessMessage(producer.BussinesPartnerAType, message)
		}

		exapected := map[string]int{
			"US-101-1": 4,
			"FR-103-1": 3,
			"DE-102-1": 2,
			"DE-102-2": 1,
		}

		assert.Equal(t, exapected, receiver.GetStats(), "stats should be equal")
	})
}

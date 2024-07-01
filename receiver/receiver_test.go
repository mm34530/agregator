package receiver_test

import "testing"

func Test_ReceiveMessage(t *testing.T) {
	t.Run("Should correctly calculate stats", func(t *testing.T) {
		data := `
		 	[{"id": "926b3c49-1cd0-4a6e-b4b5-9f68fcb01716", "source_id": 101, "country": "US", "created_at": 1645615000, "resolved_at": 1645615500}
			{"id": "0c5b76f3-bb3c-45f2-8f47-1f76ff7c04c5", "source_id": 102, "country": "DE", "created_at": 1645681000, "resolved_at": 1645682000},
			{"id": "b8216b2c-d62f-4c9c-bf4d-36a632f0a93c", "source_id": 101, "country": "US", "created_at": 1645648700, "resolved_at": 1645649500},
			{"id": "dbbf98aa-1e4d-4d5e-956b-5d2761a6d589", "source_id": 103, "country": "FR", "created_at": 1645670200, "resolved_at": 1645672900},
			{"id": "a78801d0-2a4c-4d3c-8b5a-4e76b1c1c31d", "source_id": 102, "country": "DE", "created_at": 1645663100, "resolved_at": 1645664500},
			{"id": "5fbd8f1e-231f-4d7f-9a6d-19bf52b1285f", "source_id": 103, "country": "FR", "created_at": 1645666000, "resolved_at": 1645668000},
			{"id": "b96a2e77-55a6-41d6-a9b8-6e98d8937926", "source_id": 101, "country": "US", "created_at": 1645629300, "resolved_at": 1645632000},
			{"id": "f51a54a8-7ed2-47e8-87a6-04dbcc77ea1f", "source_id": 102, "country": "DE", "created_at": 1645677400, "resolved_at": 1645682200},
			{"id": "e2d5ceeb-3c62-49ab-a146-6c1e19bbd378", "source_id": 103, "country": "FR", "created_at": 1645659300, "resolved_at": 1645660700},
			{"id": "b7e3546d-5565-489e-aa14-af6e4b4e872c", "source_id": 101, "country": "US", "created_at": 1645618600, "resolved_at": 1645620900}] 
		  `
		_ = data

	})
}

package producer_test

import (
	"agregator/producer"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateMessage(t *testing.T) {
	g := producer.NewGenerator()

	t.Run("Should return data for correct business partner type", func(t *testing.T) {
		msg, err := g.GenerateMessage(producer.BussinesPartnerAType, 100)
		assert.Nil(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("Should return error for incorrect business partner type", func(t *testing.T) {
		msg, err := g.GenerateMessage("some_type", 100)
		assert.NotNil(t, err)
		assert.Nil(t, msg)
	})
}

func Test_GenerateCountry(t *testing.T) {
	g := producer.NewGenerator()

	t.Run("Should return any country code", func(t *testing.T) {
		country := g.GenerateCountry()
		assert.NotEmpty(t, country)
	})
}

func Test_GenerateRandomDates(t *testing.T) {
	g := producer.NewGenerator()

	t.Run("Should return dates in correct range", func(t *testing.T) {
		createdAt, resolvedAt := g.GenerateRandomDates()

		createdAtTime := time.Unix(int64(createdAt), 0)
		resolvedAtTime := time.Unix(int64(resolvedAt), 0)

		assert.True(t, createdAtTime.Before(resolvedAtTime))
	})
}

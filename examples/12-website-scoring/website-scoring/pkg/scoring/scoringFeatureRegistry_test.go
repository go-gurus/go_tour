package scoring

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterScoringFeature(t *testing.T) {
	fakeFeature := func(_ *goquery.Document) float64 {
		return 1
	}

	fakeRegistration := FeatureRegistration{
		Feature: fakeFeature,
		Title:   "FAKE_REGISTRATION",
		Tags:    []string{"FAKE"},
	}

	fakeRegistration2 := FeatureRegistration{
		Feature: fakeFeature,
		Title:   "FAKE_REGISTRATION2",
		Tags:    []string{"FAKE2"},
	}

	t.Run("should return the previously registered feature when only one registered", func(t *testing.T) {
		registry := FeatureRegistry{}
		registry.Register(fakeRegistration)
		features := registry.GetFeatures()
		assert.Len(t, features, 1)
	})

	t.Run("should return all registered features when multiple registered and no filter provided", func(t *testing.T) {
		registry := FeatureRegistry{}
		registry.Register(fakeRegistration, fakeRegistration2)
		features := registry.GetFeatures()
		assert.Len(t, features, 2)
	})

	t.Run("should return features matching filter", func(t *testing.T) {
		registry := FeatureRegistry{}
		registry.Register(fakeRegistration, fakeRegistration2)
		features := registry.GetFeatures("FAKE2")
		assert.Len(t, features, 1)
	})
}

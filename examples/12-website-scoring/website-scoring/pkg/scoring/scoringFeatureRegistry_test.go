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

	RegisterScoringFeature(fakeRegistration)

	t.Run("should return the previously registered feature", func(t *testing.T) {
		features := GetFeatures()
		assert.Len(t, features, 1)
	})
}

func TestRegisterScoringFeatureFilter(t *testing.T) {
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

	RegisterScoringFeature(fakeRegistration)
	RegisterScoringFeature(fakeRegistration2)

	t.Run("should return features matching filter", func(t *testing.T) {
		features := GetFeatures("FAKE2")
		assert.Len(t, features, 1)
	})

	t.Run("should return all features when empty filter provided", func(t *testing.T) {
		features := GetFeatures()
		assert.Len(t, features, 2)
	})

}

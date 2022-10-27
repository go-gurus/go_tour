package scoring

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoring(t *testing.T) {
	fakeDocument := goquery.Document{
		Selection: nil,
		Url:       nil,
	}

	t.Run("should return score = 0 when invoked with empty document and empty featureset", func(t *testing.T) {
		assert.Equal(t, float64(0), computeScore(nil, &fakeDocument))
	})

	t.Run("should return score equal to fake set when invoked with empty document and single featureset", func(t *testing.T) {
		fakeFeatureset := FeatureSet{
			func(_ *goquery.Document) float64 {
				return 1
			},
		}
		assert.Equal(t, float64(1), computeScore(fakeFeatureset, &fakeDocument))
	})
}

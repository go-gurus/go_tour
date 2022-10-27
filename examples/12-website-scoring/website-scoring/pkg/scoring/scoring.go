package scoring

import (
	"codecentric.com/website-scoring/pkg/download"
	"github.com/PuerkitoBio/goquery"
)

type Feature func(document *goquery.Document) float64
type FeatureSet []Feature

func computeScore(features FeatureSet, document *goquery.Document) (score float64) {
	for _, feature := range features {
		score += feature(document)
	}
	return
}

func Score(url string, featureTags ...string) (score float64) {
	registry := NewDefaultRegistry()
	features := registry.GetFeatures(featureTags...)
	document, _ := download.DownloadWebsite(url)
	return computeScore(features, document)
}

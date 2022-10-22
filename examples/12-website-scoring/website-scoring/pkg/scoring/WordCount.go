package scoring

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

var sanitizeString = regexp.MustCompile("[^a-zA-ZüäößÜÄÖ]")

type WordFrequencyResult struct {
	TotalWords   int            `json:"TotalWords"`
	CountedWords int            `json:"CountedWords"`
	WordCounts   map[string]int `json:"WordCounts"`
}

func (w *WordFrequencyResult) score() float64 {
	return float64(w.TotalWords) / 10
}

func WordCount(doc *goquery.Document) float64 {
	result := WordFrequencyResult{
		WordCounts: make(map[string]int),
	}

	text := doc.Find("p").Contents().Text()
	words := strings.Split(text, " ")
	result.TotalWords = len(words)

	for i := 0; i < len(words); i++ {
		wordSanitized := strings.ToUpper(sanitizeString.ReplaceAllString(words[i], ""))
		if len(wordSanitized) > 4 {
			result.WordCounts[wordSanitized]++
		}
	}

	for _, v := range result.WordCounts {
		result.CountedWords += v
	}

	return result.score()
}

func init() {
	RegisterScoringFeature(FeatureRegistration{
		Feature: WordCount,
		Title:   "Word Count",
		Tags:    []string{"CONTENT"},
	})
}

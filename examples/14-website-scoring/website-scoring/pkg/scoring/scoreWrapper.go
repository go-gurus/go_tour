package scoring

import "github.com/PuerkitoBio/goquery"

// Scoreable constraints to types providing a method to be scored
type Scoreable interface {
	Score() float64
}

// ScoreWrapper returns a function converting its result to a score
func ScoreWrapper[T Scoreable](scoreFunction func(doc *goquery.Document) T) Feature {
	return func(document *goquery.Document) float64 {
		result := scoreFunction(document)
		return result.Score()
	}
}

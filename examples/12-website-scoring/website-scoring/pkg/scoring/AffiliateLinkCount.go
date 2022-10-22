package scoring

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

var affiliateLinksExpression, _ = regexp.Compile("(\\.amazon\\.)")
var maskedAffiliateLinksExpression, _ = regexp.Compile("(amzn\\.to)")
var shortenedUrlExpression, _ = regexp.Compile("(bit\\.ly|tinyurl\\.com)")

type LinkCountResult struct {
	TotalLinks           int `json:"TotalLinks"`
	LocalLinks           int `json:"LocalLinks"`
	AffiliateLinks       int `json:"AffiliateLinks"`
	MaskedAffiliateLinks int `json:"MaskedAffiliateLinks"`
	ShortendedUrls       int `json:"ShortendedUrls"`
}

func (l *LinkCountResult) score() float64 {
	return float64(l.TotalLinks + l.LocalLinks - l.AffiliateLinks*2 - l.MaskedAffiliateLinks*4 - l.ShortendedUrls)
}

func AffiliateLinkCount(doc *goquery.Document) float64 {
	result := LinkCountResult{}
	localdomain := doc.Url.Host

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		result.TotalLinks++
		link, exists := s.Attr("href")

		if exists {
			if strings.Contains(link, localdomain) {
				result.LocalLinks++
			}

			if affiliateLinksExpression.MatchString(link) {
				result.AffiliateLinks++
			}

			if maskedAffiliateLinksExpression.MatchString(link) {
				result.MaskedAffiliateLinks++
			}

			if shortenedUrlExpression.MatchString(link) {
				result.ShortendedUrls++
			}
		}

	})
	return result.score()
}

func init() {
	RegisterScoringFeature(FeatureRegistration{
		Feature: AffiliateLinkCount,
		Title:   "Affiliate Link Count",
		Tags:    []string{"MARKETING", "AFFILIATE", "UNTRUSTWORTHY"},
	})
}

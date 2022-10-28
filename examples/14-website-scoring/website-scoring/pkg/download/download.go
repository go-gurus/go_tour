package download

import (
	"github.com/PuerkitoBio/goquery"
)

func DownloadWebsite(url string) (*goquery.Document, error) {
	return goquery.NewDocument(url)
}

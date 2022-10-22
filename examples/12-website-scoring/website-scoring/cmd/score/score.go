package main

import (
	"codecentric.com/website-scoring/pkg/scoring"
	"flag"
	"fmt"
)

func main() {
	var targetUrl = flag.String("url", "", "URL of the site to be parsed")
	flag.Parse()
	score := scoring.Score(*targetUrl)
	fmt.Printf("Website=%s Score=%f", *targetUrl, score)
}

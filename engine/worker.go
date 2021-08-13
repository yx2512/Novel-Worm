package engine

import (
	"log"

	"github.com/yx2512/crawler/fetcher"
)

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher: error fetching url: %s with error: %v", r.Url, err)
		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(body)

	return parseResult, nil
}

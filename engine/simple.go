package engine

import (
	"crowler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request) {

	//队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		worker, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, worker.Requests...)

		for _, item := range worker.Items {
			log.Printf("Got Item %v", item)
		}

	}
}

func worker(r Request) (ParseResult,error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetch url:%s, %v",r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body),nil
}

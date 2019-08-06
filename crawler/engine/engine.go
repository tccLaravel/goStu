package engine

import (
	"goStu/crawler/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds{
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		bytes, e := fetcher.Fetch(r.Url)
		if e != nil{
			log.Printf("Fetcher: error fetching url %s: %v",r.Url,e)
			continue
		}
		ParseResult := r.ParserFunc(bytes)
		requests = append(requests,ParseResult.Requests...)
		for _,item := range ParseResult.Items{
			log.Printf("Got chapter %s \n",item)
		}
	}
}

package engine

import (
	"goStu/crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds{
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := woker(r)
		//noinspection GoNilness
		requests = append(requests,parseResult.Requests...)
		//noinspection GoNilness
		for _,item := range parseResult.Items{
			log.Printf("Got chapter %s \n",item)
		}
		if err != nil{
			log.Printf("END2")
			continue
		}
	}
}

/**
通过url获取内容,然后用不同的解析器解析解析内容
 */
func woker(r Request)  (ParseResult, error){
	bytes, e := fetcher.Fetch(r.Url)
	if e != nil{
		log.Printf("Fetcher: error fetching url %s: %v",r.Url,e)
		return ParseResult{},e
	}
	return r.ParserFunc(bytes),nil
}

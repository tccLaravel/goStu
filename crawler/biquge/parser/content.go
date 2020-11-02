package parser

import (
	"fmt"
	"goStu/crawler/engine"
	"goStu/crawler/model"
	"log"
	"regexp"
	"strconv"
)

var contentsRe = regexp.MustCompile(`<div id="content">(?s:(.*?))</div>`)
var titleRe = regexp.MustCompile(`<h1>([0-9]+)([^<]+)</h1>`)

func ContentsParser(contents []byte) engine.ParseResult{
	result  := engine.ParseResult{}
	article := model.Article{}

	titleMatchs := titleRe.FindSubmatch(contents)
	if titleMatchs != nil {
		article.Title = string(titleMatchs[2])
		chapterNum,err := strconv.Atoi(string(titleMatchs[1]))
		fmt.Println(chapterNum,err)
		if err != nil{
			article.ChapterNum = 0
		}else{
			article.ChapterNum = chapterNum
		}
	}
	article.Content = extractString(contents,contentsRe)
	result.Items = append(result.Items,article)
	log.Printf("Got Result: %v \n",result)
	return result
}

func extractString(contents []byte,re *regexp.Regexp)  string{
	match := re.FindSubmatch(contents)
	if match != nil && len(match) >= 2{
		return string(match[1])
	}
	return ""
}
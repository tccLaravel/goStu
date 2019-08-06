package parser

import (
	"regexp"
)
//([\s|\S]*)?
const ContentsParserRe = `<div id="content">(?s:(.*?))<\/div>`

//const ContentsParserRe = `<li><a href="/xclass/1/1.html">([^<]+)</a></li>`
func ContentsParser(contents []byte) []interface{}{
	re := regexp.MustCompile(ContentsParserRe)
	//log.Printf(" regexp  %v \n",re)
	//log.Printf("contents is %s \n",contents)
	matchs := re.FindAllSubmatch(contents,-1)
	//matchs := re.FindAll(contents, -1)
	var result  []interface{}
	//log.Printf(" matchs0----- %s \n",matchs[0])
	for _, m := range matchs{
		//fmt.Printf("match %s \n",m[1])
		result = append(result, m[1])
	}
	return result
}
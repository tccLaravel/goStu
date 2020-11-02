package parser

import (
	"fmt"
	"goStu/crawler/engine"
	"os"
	"regexp"
)

const fileName  = "./tcc.txt"
const parserChaptersRe  = `<dd><a href="(/77_77363/[0-9]+.html)"[^>]*>([^<]+)</a></dd>`
func ChaptersParser(contents []byte)  engine.ParseResult{
	re := regexp.MustCompile(parserChaptersRe)
	matchs := re.FindAllSubmatch(contents, -1)
	//matchs := re.FindAll(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs{
		//fmt.Printf("%s \n",m)
		//fmt.Printf("chapter: %s, URL: %s \n",m[2],m[1])
		//result.Items = append(result.Items,string(m[2]))
		err := appendToFile(fileName, string(m[2])+"\n")
		if err != nil {
			fmt.Println("write error : " + err.Error())
		}
		result.Requests = append(result.Requests,engine.Request{
			//Url:string(m[1]),
			Url:fmt.Sprintf("%s%s","https://www.xbiquge6.com",string(m[1])),
			ParserFunc:ContentsParser,
		})
	}
	return result
}

func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		//n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.Write([]byte(content))
	}
	defer f.Close()
	return err
}


package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"goStu/crawler/biquge/parser"
	engine2 "goStu/crawler/engine"
	"goStu/route"
	"net/http"
)

func main() {
	engine2.Run(engine2.Request{
		Url:"https://www.xbiquge6.com/77_77363/",
		ParserFunc:parser.ChaptersParser,
	})
	fmt.Println("使用外部包测试",configor.Config{})
	fmt.Println("使用内部包测试",route.Hello())
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	engine.Any("/", WebRoot)
	// 绑定端口，然后启动应用
	_ = engine.Run(":9205")
}
/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
 */
func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}

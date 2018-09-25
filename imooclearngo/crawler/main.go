package main

import (
	"lt.go/imooclearngo/crawler/engine"
	"lt.go/imooclearngo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

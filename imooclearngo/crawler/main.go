package main

import (
	"github.com/lightjameslyy/lt.go/imooclearngo/crawler/engine"
	"github.com/lightjameslyy/lt.go/imooclearngo/crawler/persist"
	"github.com/lightjameslyy/lt.go/imooclearngo/crawler/scheduler"
	"github.com/lightjameslyy/lt.go/imooclearngo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}

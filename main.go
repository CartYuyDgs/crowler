package main

import (
	"crowler/engine"
	"crowler/fang/parser"
	"crowler/schaduler"
)

const host = "https://www.fang.com/SoufunFamily.htm"

func main() {

	e := engine.ConcurrentEngine{
		Scheduler: &schaduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url: host,
		ParserFunc: parser.ParseCityList,
	})

}


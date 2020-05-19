package main

import (
	"crowler/engine"
	"crowler/fang/parser"
)

const host = "https://www.fang.com/SoufunFamily.htm"

func main() {

	engine.Run(engine.Request{
		Url: host,
		ParserFunc: parser.ParseCityList,
	})

}


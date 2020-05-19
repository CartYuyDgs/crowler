package parser

import (
	"crowler/engine"
	"regexp"
)

//<a href="http://ningxiang.fang.com/" target="_blank">宁乡</a>
const CityListRe = `<a href="(http://[a-z0-9A-Z]{1,20}.fang.com)/" target="_blank">(.+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	var houstList2Re = regexp.MustCompile(CityListRe)
	matchs := houstList2Re.FindAllSubmatch(contents,-1)

	//每一个url生成一个新的request
	result := engine.ParseResult{}
	for _,m := range matchs {
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
		//fmt.Printf("%d: City:%s,Url:%s\n",i,m[2],m[1])
	}
	return result
}
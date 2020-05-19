package parser

import (
	"crowler/engine"
	"regexp"
	"strings"
)

//<a href="http://ningxiang.fang.com/" target="_blank">宁乡</a>
const CityListRe = `<a href="(http://[a-z]{1,20}.fang.com)/" target="_blank">(.+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	var houstList2Re = regexp.MustCompile(CityListRe)
	matchs := houstList2Re.FindAllSubmatch(contents,-1)

	//每一个url生成一个新的request
	result := engine.ParseResult{}
	for _,m := range matchs {
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
			Url:        UrlTransNewhourse(string(m[1])),
			ParserFunc: ParseCity,
		})
		//fmt.Printf("%d: City:%s,Url:%s\n",i,m[2],m[1])
	}
	return result
}

func UrlTransNewhourse(url string) string {
	//https://sh.fang.com/
	//https://sh.newhouse.fang.com/house/s/
	matchs := strings.Split(url,".")
	return matchs[0]+".newhouse."+matchs[1]+"."+matchs[2]+"/house/s/"
}
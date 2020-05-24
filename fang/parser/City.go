package parser

import (
	"crowler/engine"
	"fmt"
	"regexp"
)

//const CityRe = "<a  target=\"_blank\" href=\"(//[a-z+].fang.com/)\">"
const CityRe = "<a target=\"_blank\" data-yd=\"\" onclick=yidiCityCookie\"(.*)\";  href=\"(//[a-z0-9+].fang.com/)\">"


//<a target="_blank" data-yd="" onclick=yidiCityCookie\(\".*\"\)
//<a target="_blank" data-yd="" onclick=yidiCityCookie\(\".*\"\);  href="//[a-z0-9]+.[a-z]+.com
//<a target="_blank" data-yd="" onclick=yidiCityCookie\(\".*\"\);  href="//[a-z0-9]+.[a-z]+.com/">.*
const houseRe = `<a target="_blank" data-yd="" onclick=yidiCityCookie\(\"(.*)\"\);  href="//([a-z0-9]+.[a-z]+.com)/">`


//https://yunbaofurongyuanlg.fang.com/

//解析页面，返回next Url和页面信息
func ParseCity(contents []byte) engine.ParseResult {
	var houstList2Re = regexp.MustCompile(houseRe)
	matchs := houstList2Re.FindAllSubmatch(contents,-1)

	//每一个url生成一个新的request
	result := engine.ParseResult{}
	for _,m := range matchs {
		fmt.Printf("*******city%s,Url:%s,len:%d\n",m[1],m[2],len(m))
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        "https://"+string(m[2]),
				ParserFunc: TransHouseProfile,
			})

	}
	return result
}

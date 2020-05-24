package parser

import (
	"crowler/engine"
	"crowler/model"
	"log"
	"regexp"
	"strings"
)

//https://yunbaofurongyuanlg.fang.com/house/3410167990/housedetail.htm
var profileRe = regexp.MustCompile(`<a href="(//[a-z0-9]+.fang.com/house/[0-9]+/housedetail.htm)" id=".*"  target="_self">`)
//<h1><a class="ts_linear" id="huxinxq_E02_15" href="//binjiangguojihw.fang.com/" title="恒威滨江国际" target="_blank">恒威滨江国际</a></h1>
var houstnameRe = regexp.MustCompile(`<h1><a class="ts_linear" id=".*" href=".*" title=".*" target="_blank">(.*)</a></h1>`)

var pricedRe = regexp.MustCompile(`<p><b>(.*)</b><em>(.*)</em></p>`)
var starRe = regexp.MustCompile(`<span style="margin-right: 5px;">(.*)</span>`)
var categoryRe = regexp.MustCompile(`<div class="list-right" title="(.*)">`)
var buildingRe = regexp.MustCompile(`<div class="list-right"><span class="bulid-type">([\s\S]*)</span></div>`)
var propertyRe = regexp.MustCompile(`<p style="width: 130px;float: left;">(.*)</p>`)
var renovateRe = regexp.MustCompile(`<div class="list-left">装修状况：</div>
                       <div class="list-right">([\s]+.*[\s]+)</div>
                    </li>`)

var trafficRe = regexp.MustCompile(`<span>交通</span>(.*)<br />`)
var footPrintRe = regexp.MustCompile(`<div class="list-left">建筑面积：</div>
                       <div class="list-right">(.*)</div>`)

var volumeRe = regexp.MustCompile(`<div class="list-left">绿<i style="margin-right: 6px;"></i>化<i style="margin-right: 6px;"></i>率：</div>
                       <div class="list-right">(.*)</div>`)

var parkingRe = regexp.MustCompile(`<div class="list-left">停<i style="margin-right: 6px;"></i>车<i style="margin-right: 6px;"></i>位：</div>
                       <div class="list-right" title=".*">(.*)</div>`)

var houseNumRe = regexp.MustCompile(`<div class="list-left">楼栋总数：</div>
                       <div class="list-right">(.*)</div>`)

//跳转详情页面
func TransHouseProfile(contents []byte) engine.ParseResult {

	request := engine.ParseResult{}
	match := profileRe.FindSubmatch(contents)
	if len(match) >=2 {
		log.Printf("TransHouseProfile Url:%s\n",match[1])
		request.Requests = append(request.Requests,engine.Request{
			Url: "https:"+string(match[1]),
			ParserFunc: ParseDetailInfo,
		})
	}
	return request
}

func ParseDetailInfo(contents []byte) engine.ParseResult {
	log.Println("commin in  ParseDetailInfo....................")
	houseProfile := model.HouseProfile{}

	houseName := extractString(contents,houstnameRe)
	houseProfile.HouseName = string(houseName)

	priced := getHousePriced(contents,pricedRe)
	houseProfile.Priced = append(houseProfile.Priced, priced...)

	star := extractString(contents,starRe)
	houseProfile.Star = star

	category := extractString(contents,categoryRe)
	houseProfile.Category = category

	building := extractString(contents, buildingRe)
	building = strings.Replace(building,"\t","",-1)
	building = strings.Replace(building,"\n","",-1)
	houseProfile.Building = building

	renovate := extractString(contents, renovateRe)
	renovate = strings.Replace(renovate,"\t","",-1)
	renovate = strings.Replace(renovate,"\n","",-1)
	houseProfile.Renovate = renovate

	property := getHouseProperty(contents, propertyRe)
	houseProfile.Property = property

	traffic := extractString(contents, trafficRe)
	houseProfile.Traffic = traffic

	footPrint := extractString(contents, footPrintRe)
	houseProfile.FootPrint = footPrint

	volume := extractString(contents, volumeRe)
	houseProfile.Volume = volume

	parking := extractString(contents, parkingRe)
	houseProfile.Parking = parking

	houseNum := extractString(contents, houseNumRe)
	houseProfile.HouseNum = houseNum

	log.Println("comman in  ParseDetailInfo...........end.........")

	result := engine.ParseResult{
		Items: []interface{} {houseProfile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}else {
		return ""
	}
}

func getHousePriced(contents []byte, re* regexp.Regexp) []string {

	var result []string
	match := re.FindAllSubmatch(contents, -1)
	for _,m := range match {
		//log.Printf("------------%s,%s\n",m[1],m[2])
		if len(m) >= 3 {
			result = append(result,string(m[1]) + ":" + string(m[2]))
		}
	}
	return result
}

func getHouseProperty(contents []byte, re* regexp.Regexp) []string {

	var result []string
	match := re.FindAllSubmatch(contents, -1)
	for _,m := range match {
		//log.Printf("------------%s,%s\n",m[1],m[2])
		if len(m) >= 2 {
			result = append(result,string(m[1]))
		}
	}
	return result
}

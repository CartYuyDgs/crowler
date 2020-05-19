package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"regexp"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

const host = "https://www.fang.com/SoufunFamily.htm"

//<a href="http://ningxiang.fang.com/" target="_blank">宁乡</a>
var houstList2Re = regexp.MustCompile(`<a href="(http://[a-z]{1,20}.fang.com)/" target="_blank">(.+)</a>`)

func main() {
	resp, err := http.Get(host)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: statuscode is ", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	content, err  := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	match := houstList2Re.FindAllSubmatch(content,-1)
	for i,m := range match {
		fmt.Printf("%d: %s,%s\n",i,m[2],m[1])
	}
}

func determineEncoding(r io.Reader) (e encoding.Encoding) {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, name, b := charset.DetermineEncoding(bytes,"")
	fmt.Println(e,name,b)
	return
}
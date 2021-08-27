package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type test struct {
	XMLName xml.Name `xml:"test"`
	Abc[]     abc    `xml:"abc"`
	Eee     string   `xml:"eee"`
}

type abc struct {
	Key   string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

func main() {
	// xml 파일 오픈
	fp, err := os.Open("/home/thkim/GolandProjects/project_210809/xml_test_2.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, err := ioutil.ReadAll(fp)

	// xml 디코딩
	var t test
	err = xml.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}



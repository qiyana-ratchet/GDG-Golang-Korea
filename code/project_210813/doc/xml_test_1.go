package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Member -
type Member struct {
	Name   string
	Age    int
	Active bool
}

//Members -
type Members struct {
	Member []Member
}

func main() {
	//////////////////////
	// xml 파일 오픈
	fp, err := os.Open("/home/thkim/GolandProjects/project_210809/xml_test_1.xml")
	//home/thkim/GolandProjects/project_210809/data_cut.xml
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, err := ioutil.ReadAll(fp)

	// xml 디코딩
	var members Members
	xmlerr := xml.Unmarshal(data, &members)
	if xmlerr != nil {
		panic(xmlerr)
	}
	//fmt.Println(data)

	fmt.Println(members)
	//////////////////////
}
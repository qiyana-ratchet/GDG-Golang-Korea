package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type measInfo struct {
	XMLName xml.Name `xml:"measInfo"`
	Job     job   `xml:"job"`
/*	GranPeriod string `xml:"granPeriod,attr,attr"`
	RepPeriod string `xml:"repPeriod,attr"`*/
	MeAsType[] measType `xml:"measType"`
	MeAsValue measValue `xml:"measValue"`
}

type job struct {
	Key   string `xml:"jobId,attr"`
}

type measType struct {
	Key   string `xml:"p,attr"`
	Value string `xml:",chardata"`
}

type measValue struct {
	XMLName xml.Name `xml:"measValue"`
	Key   string `xml:"measObjLdn,attr"`
	R[]   r     `xml:"r"`
}

type r struct {
	Key  string  `xml:"p,attr"`
	Value string `xml:",chardata"`
}

func main() {
	// xml 파일 오픈
	fp, err := os.Open("/home/thkim/GolandProjects/project_210809/data_cut.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, err := ioutil.ReadAll(fp)

	// xml 디코딩
	var pmInfo measInfo
	err = xml.Unmarshal(data, &pmInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println(pmInfo)
}


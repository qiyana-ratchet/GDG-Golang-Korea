package collector

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type measDataFile struct {
	XMLName    xml.Name   `xml:"measDataFile"`
	FileHeader fileHeader `xml:"fileHeader"`
	MeasData   measData   `xml:"measData"`
	FileFooter fileFooter `xml:"fileFooter"`
}
type fileHeader struct {
	XMLName           xml.Name   `xml:"fileHeader"`
	FileFormatVersion string     `xml:"fileFormatVersion,attr"`
	VendorName        string     `xml:"vendorName,attr"`
	DnPrefix          string     `xml:"dnPrefix,attr"`
	FileSender        fileSender `xml:"fileSender"`
	MeasData          measData   `xml:"measData"`
}
type fileSender struct {
	XMLName    xml.Name `xml:"fileSender"`
	SenderName string   `xml:"senderName,attr"`
	SenderType string   `xml:"senderType,attr"`
}
type fileFooter struct {
	XMLName  xml.Name `xml:"fileFooter"`
	MeasData measData `xml:"measData"`
}
type measData struct {
	XMLName    xml.Name   `xml:"measData"`
	MeasEntity measEntity `xml:"measEntity"`
	MeasInfo   []measInfo `xml:"measInfo"`
	BeginTime  string     `xml:"beginTime,attr"`
	EndTime    string     `xml:"endTime,attr"`
}
type measEntity struct {
	XMLName xml.Name `xml:"measEntity"`
	Key     string   `xml:"localDn,attr"`
	Key2    string   `xml:"swVersion,attr"`
}
type measInfo struct {
	XMLName    xml.Name   `xml:"measInfo"`
	MeasInfoID string     `xml:"measInfoId,attr"`
	Job        job        `xml:"job"`
	GranPeriod granPeriod `xml:"granPeriod"`
	RepPeriod  repPeriod  `xml:"repPeriod"`
	MeasType   []measType `xml:"measType"`
	MeasValue  measValue  `xml:"measValue"`
}
type job struct {
	XMLName xml.Name `xml:"job"`
	//XMLAttr xml.Attr `xml:"jobId,attr"`
	Key string `xml:"jobId,attr"`
}
type granPeriod struct {
	XMLName xml.Name `xml:"granPeriod"`
	Key     string   `xml:"duration,attr"`
	Key2    string   `xml:"endTime,attr"`
}
type repPeriod struct {
	XMLName xml.Name `xml:"repPeriod"`
	Key     string   `xml:"duration,attr"`
}
type measType struct {
	XMLName xml.Name `xml:"measType"`
	Key     string   `xml:"p,attr"`
	Value   string   `xml:",chardata"`
}
type measValue struct {
	XMLName xml.Name `xml:"measValue"`
	Key     string   `xml:"measObjLdn,attr"`
	R       []r      `xml:"r"`
}
type r struct {
	XMLName xml.Name `xml:"r"`
	Key     string   `xml:"p,attr"`
	Value   string   `xml:",chardata"`
}


func getPmData() (pmdata, error) {
	//var utsname unix.Utsname
	//if err := unix.Uname(&utsname); err != nil {
	//	return pmdata{}, err
	//}


	///////////////////
	// xml 파일 오픈
	fp, err := os.Open("/home/thkim/myhfrlab/test-node-exporter/collector/data_parse.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, err := ioutil.ReadAll(fp)

	// xml 디코딩
	var measDataFile measDataFile
	err = xml.Unmarshal(data, &measDataFile)
	if err != nil {
		panic(err)
	}
	//fmt.Println(ch)
	//printMeastype(measDataFile)
	///////////////////


	output := pmdata{
		SysName:    measDataFile.MeasData.MeasInfo[0].MeasInfoID,
		Release:    measDataFile.MeasData.MeasInfo[0].MeasInfoID,
		Version:    measDataFile.MeasData.MeasInfo[0].MeasInfoID,
		Machine:    measDataFile.MeasData.MeasInfo[0].MeasInfoID,
		NodeName:   measDataFile.MeasData.MeasInfo[0].MeasInfoID,
		DomainName: measDataFile.MeasData.MeasInfo[0].MeasInfoID,
	}

	return output, nil
}


//메트릭 프린트 함수
func printMeastype(measDataFile measDataFile) {
	measInfoList := measDataFile.MeasData.MeasInfo
	measInfoListLen := len(measInfoList)

	for i := 0; i < measInfoListLen; i++ {
		measTypeList := measDataFile.MeasData.MeasInfo[i].MeasType
		measInfoIdValue := measDataFile.MeasData.MeasInfo[i].MeasInfoID
		measTypeListLen := len(measTypeList)
		//fmt.Println(measInfoListLen, measTypeListLen)
		//메트릭 개수 디버깅용

		for j := 0; j < measTypeListLen; j++ {
			metricKey := strings.ToLower(strings.ReplaceAll(measTypeList[j].Value, ".", "_"))
			metricValue := measInfoList[i].MeasValue.R[j].Value
			printMetricInfo(metricKey)
			fmt.Println(metricKey+"{measInfoID=\""/*+measInfoIdName+*/+ measInfoIdValue+"\"}", metricValue)
		}
	}
}


//HELP 출력
func printMetricInfo(metricKey string) {
	fmt.Println("# HELP", metricKey, "This is sample info.")
	fmt.Println("# TYPE", metricKey, "typeName")
}

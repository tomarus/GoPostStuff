package main

import (
	"encoding/xml"
	"io/ioutil"
	//	"path"
	"regexp"
	"strings"
)

type Nzb struct {
	XMLName xml.Name  `xml:"nzb"`
	XMLns   string    `xml:"xmlns,attr"`
	Head    []Meta    `xml:"head>meta"`
	File    []NzbFile `xml:"file"`
}

type Meta struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",innerxml"`
}

type NzbFile struct {
	Poster   string       `xml:"poster,attr"`
	Date     int64        `xml:"date,attr"`
	Subject  string       `xml:"subject,attr"`
	Groups   []string     `xml:"groups>group"`
	Segments []NzbSegment `xml:"segments>segment"`
}

type NzbSegment struct {
	XMLName   xml.Name `xml:"segment"`
	Bytes     int64    `xml:"bytes,attr"`
	Number    int64    `xml:"number,attr"`
	MessageId string   `xml:",innerxml"`
}

func CreateNzb(filename string, nzb *Nzb) error {
	nzb.XMLns = "http://www.newzbin.com/DTD/2003/nzb"
	output, err := xml.MarshalIndent(nzb, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, output, 0x777)
	if err != nil {
		return err
	}
	return nil
}

func SafeFileName(str string) string {
	name := strings.ToLower(str)
	//name = path.Clean(path.Base(name))
	name = strings.Trim(name, " ")
	separators, err := regexp.Compile(`[ &_=+:]`)
	if err == nil {
		name = separators.ReplaceAllString(name, "-")
	}
	legal, err := regexp.Compile(`[^[:alnum:]-.]|-yenc-\d*.+`)
	if err == nil {
		name = legal.ReplaceAllString(name, "")
	}
	for strings.Contains(name, "--") {
		name = strings.Replace(name, "--", "-", -1)
	}
	return name
}

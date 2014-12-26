package main

import (
        "encoding/xml"
        "io/ioutil"
)

type Nzb struct {
    XMLName xml.Name    `xml:"nzb"`
    Head    []Meta      `xml:"head>meta"`
    File    []NzbFile   `xml:"file"`
}

type Meta struct {
    Type  string `xml:"type,attr"`
    Value string `xml:",innerxml"`
}

type NzbFile struct {
    Poster   string    `xml:"poster,attr"`
    Date     int64    `xml:"date,attr"`
    Subject  string    `xml:"subject,attr"`
    Groups   []string  `xml:"groups>group"`
    Segments []NzbSegment `xml:"segments>segment"`
}

type NzbSegment struct {
    XMLName     xml.Name    `xml:"segment"`
    Bytes       int64       `xml:"bytes,attr"`
    Number      int64       `xml:"number,attr"`
    MessageId   string      `xml:",innerxml"`
}

func CreateNzb(filename string, nzb *Nzb) (error) {
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


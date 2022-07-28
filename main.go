package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Feed struct {
	XMLName   xml.Name `xml:"feed"`
	Title     string   `xml:"title"`
	EntryList []Entry  `xml:"entry"`
}

type Entry struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Author      string `xml:"author"`
	PublishedAt string `xml:"published"`
	Summary     string `xml:"summary"`
}

const (
	RSSURL = "https://scu.ac.ir/%D8%B5%D9%81%D8%AD%D9%87-%D8%A7%D8%B5%D9%84%DB%8C/-/asset_publisher/HDov2wBiUjjC/rss?p_p_cacheability=cacheLevelPage"
)

func sendReqToSCU() *http.Response {

	resp, err := http.Get(RSSURL)

	simpleHandelError(err)
	return resp

}

func readByteDataFromSCU() []byte {

	resp := sendReqToSCU()
	byteDate, err := ioutil.ReadAll(resp.Body)

	simpleHandelError(err)

	return byteDate

}

func simpleHandelError(err error) {

	if err != nil {
		fmt.Println("err : ", err)
		os.Exit(1)
	}

}

func decodeXMLData() *Feed {

	byteData := readByteDataFromSCU()

	var r Feed

	err := xml.Unmarshal(byteData, &r)

	simpleHandelError(err)

	return &r

}

func main() {

	rssFeed := decodeXMLData()

	f, _ := os.Create("Entry.txt")

	defer f.Close()

	result := rssFeed.EntryList[0].Summary

	result = strings.Replace(result, "&nbsp;", "", -1)
	f.WriteString(result)

}

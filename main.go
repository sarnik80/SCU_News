package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Feed struct {
	XMLName   xml.Name
	Title     string
	EntryList []Entry
}

type Entry struct {
	Title       string
	Link        string
	Author      string
	PublishedAt string
	Summary     string
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

func main() {

}

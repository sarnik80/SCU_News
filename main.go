package main

import "encoding/xml"

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

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const blob = `[
	{"Title": "0redev", "URL": "http://oredev.org"},
	{"Title": "Strange Loop", "URL": "http://thestrangeloop.com"}
]`

type Item struct {
	Title string
	URL   string
}

func (item *Item) String() string {
	return fmt.Sprintf("Title: %v URL: %v", item.Title, item.URL)
}

func main() {
	var items []*Item
	err := json.NewDecoder(strings.NewReader(blob)).Decode(&items)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}

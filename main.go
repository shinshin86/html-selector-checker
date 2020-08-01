package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/logrusorgru/aurora"
)

func main() {
	htmlFile := flag.String("f", "index.html", "HTML file path")
	selector := flag.String("s", "body", "Selector")
	isDebug := flag.Bool("d", false, "Is debug mode")
	flag.Parse()

	if *isDebug == true {
		fmt.Println(BrightGreen(Bold("HTML file: " + *htmlFile)))
		fmt.Println(BrightGreen(Bold("Selector: " + *selector)))
	}

	fileInfos, err := ioutil.ReadFile(*htmlFile)

	if err != nil {
		panic(err)
	}

	stringReader := strings.NewReader(string(fileInfos))
	doc, err := goquery.NewDocumentFromReader(stringReader)

	if err != nil {
		panic(err)
	}

	selection := doc.Find(*selector)

	if *isDebug == true {
		fmt.Println(BrightGreen(Bold("==========")))
		fmt.Println(selection.Text())
		fmt.Println(BrightGreen(Bold("==========")))
	}

	selection.Each(func(index int, s *goquery.Selection) {
		fmt.Println(BrightYellow("=== " + strconv.Itoa(index) + " ==="))
		fmt.Println(s.Text())
	})
}

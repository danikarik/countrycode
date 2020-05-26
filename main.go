package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

var (
	output = flag.String("output", "", "Output file path")
)

func exit(code int, format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Println("")
	os.Exit(code)
}

func main() {
	flag.Parse()

	if *output == "" {
		exit(1, "output must be specified")
	}

	file, err := os.Create(*output)
	if err != nil {
		exit(1, "could not open %s: %v", *output, err)
	}
	defer file.Close()

	hmap := make(map[string]string, 0)

	c := colly.NewCollector()

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		name, code := "", ""

		e.ForEach("td", func(row int, ee *colly.HTMLElement) {

			if row == 0 {
				name = ee.Text
			} else if row == 1 {
				code = ee.Text
			}
		})

		if name != "" && code != "" {
			hmap[name] = code
		}
	})

	if err := c.Visit("https://countrycode.org/"); err != nil {
		exit(1, "could not parse page: %v", err)
	}

	for k, v := range hmap {
		file.WriteString(fmt.Sprintf("\"%s\": \"%s\",\n", k, v))
	}
}

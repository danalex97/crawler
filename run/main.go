package main

import (
  "github.com/danalex97/crawler"
)

func main() {
  c := crawler.NewCrawler("http:://tomblomfield.com")
  c.Run()
}

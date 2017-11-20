package main

import (
  "github.com/danalex97/crawler"
)

func main() {
  c := crawler.NewCrawler("http://tomblomfield.com")
  s := crawler.NewServer(c)

  go c.Run()
  s.ServerListenPathPort("/", 30000)
}

package main

import (
  "github.com/danalex97/crawler"
  "os"
)

func main() {
  args := os.Args

  c := crawler.NewCrawler(args[1])
  s := crawler.NewServer(c)

  go c.Run()
  s.ServerListenPathPort("/", 30000)
}

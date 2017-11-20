package crawler

import (
  "fmt"
)

type Crawler struct {
  domain   string
  fetchers []*fetcher
  sitemap  *sitemap
}

func NewCrawler(domain string) *Crawler {
  c := new(Crawler)

  c.domain   = domain
  c.fetchers = make([]*fetcher, 0)
  c.sitemap  = newSitemap(domain)
  c.fetchers = append(c.fetchers, newFetcher(domain))

  return c
}

func (c *Crawler) Run() error {
  for {
    fetchingChannel := make(chan fetchedData)
    for _, currFetcher := range c.fetchers {
      go func(currFetcher *fetcher) {
        fetchingChannel <- currFetcher.fetch()
      } (currFetcher)
    }

    buildingChannel := make(chan []string)
    lenFetchers     := len(c.fetchers)
    for i := 0; i < lenFetchers; i++ {
      data           := <- fetchingChannel
      url, urls, err := fromFetchedData(data)

      if err != nil {
        continue
      }

      go func() {
        builder := newBuilder(c.domain, c.sitemap)
        buildingChannel <- builder.buildPage(url, builder.filterPages(urls))
      } ()
    }

    uniqueUrls := map[string]struct{}{}
    for i := 0; i < lenFetchers; i++ {
      urls := <-buildingChannel
      for _, url := range urls {
        if _, ok := uniqueUrls[url]; !ok {
          uniqueUrls[url] = struct{}{}
        }
      }
    }

    c.fetchers = []*fetcher{}
    for url, _ := range(uniqueUrls) {
      c.fetchers = append(c.fetchers, newFetcher(url))
    }

    if len(c.fetchers) == 0 {
      fmt.Println("No more links found")
      return nil
    }
  }
  return nil
}

package crawler

import (
  "fmt"
  "net/http"
)

type Crawler struct {
  domain   string
  fetchers []fetcher
  sitemap  *sitemap
}

func NewCrawler(domain string) *Crawler {
  c := new(Crawler)

  c.domain   = domain
  c.fetchers = make([]fetcher, 0)
  c.sitemap  = newSitemap(domain)
  c.fetchers = append(c.fetchers, newHttpFetcher(domain, http.Client{}))

  return c
}

func (c *Crawler) fetchPages() chan fetchedData {
  // Calls fetch for each Fetcher and pipes the resulting data into a channel
  fetchingChannel := make(chan fetchedData)
  for _, currFetcher := range c.fetchers {
    go func(currFetcher fetcher) {
      fetchingChannel <- currFetcher.fetch()
    } (currFetcher)
  }
  return fetchingChannel
}

func (c *Crawler) buildPages(fetchingChannel chan fetchedData) chan []string {
  /* Gets the data from fetchers (see data.go), filteres it and builds
     all the correspoing pages. It resturns a channel with all the new urls.
     The new urls will be put in buildingChannel for further processing. */
  buildingChannel := make(chan []string)
  lenFetchers     := len(c.fetchers)

  for i := 0; i < lenFetchers; i++ {
    data                   := <- fetchingChannel
    url, urls, assets, err := fromFetchedData(data)

    if err != nil {
      continue
    }

    go func() {
      builder := newBuilder(c.domain, c.sitemap)
      buildingChannel <- builder.buildPage(url, builder.filterPages(urls), assets)
    } ()
  }
  return buildingChannel
}

func (c *Crawler) updateFetchers(buildingChannel chan []string) {
  // Creates a new list of fetchers from all unique newly available urls.
  uniqueUrls  := map[string]struct{}{}
  lenFetchers := len(c.fetchers)

  for i := 0; i < lenFetchers; i++ {
    urls := <-buildingChannel
    for _, url := range urls {
      if _, ok := uniqueUrls[url]; !ok {
        uniqueUrls[url] = struct{}{}
      }
    }
  }

  c.fetchers = []fetcher{}
  for url, _ := range(uniqueUrls) {
    c.fetchers = append(c.fetchers, newHttpFetcher(url, http.Client{}))
  }
}

func (c *Crawler) Run() error {
  for {
    fetchingChannel := c.fetchPages()
    buildingChannel := c.buildPages(fetchingChannel)
    c.updateFetchers(buildingChannel)

    if len(c.fetchers) == 0 {
      fmt.Println("No more links found")
      return nil
    }
  }
  return nil
}

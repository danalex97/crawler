package crawler

import (
  "strings"
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

func (c *Crawler) filterPages(pages []*page) (urls []string) {
  for _, page := range pages {
    url := page.getUrl()
    if strings.HasPrefix(url, c.domain) {
      urls = append(urls, url)
    }
  }
  return
}

func (c *Crawler) Run() error {
  for {
    fetchingChannel := make(chan fetchedData)
    for _, currFetcher := range c.fetchers {
      go func(currFetcher *fetcher) {
        fetchingChannel <- currFetcher.fetch()
      } (currFetcher)
    }

    lenFetchers := len(c.fetchers)
    for i := 0; i < lenFetchers; i++ {
      data           := <- fetchingChannel
      url, urls, err := fromFetchedData(data)

      if err != nil {
        continue
      }

      c.sitemap.Lock()
      page := c.sitemap.getPage(url)
      if page == nil {
        c.sitemap.Unlock()
        continue
      }
      for _, ref := range urls {
        target := c.sitemap.getPage(ref)
        if target != nil {
          page.addLink(target)
        }
      }
      page.setParsed()
      c.sitemap.Unlock()

      c.fetchers = make([]*fetcher, 0)

      c.sitemap.Lock()
      for _, url := range urls {
        c.fetchers = append(c.fetchers, newFetcher(url))
      }
      c.sitemap.Unlock()
    }
  }
  return nil
}

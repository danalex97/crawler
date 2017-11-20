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
  c.sitemap  = newSitemap()

  c.fetchers = append(c.fetchers, newFetcher(domain, c.sitemap))
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
    for _, fetcher := range c.fetchers {
      fetcher.fetch()
    }
    c.fetchers = make([]*fetcher, 0)

    c.sitemap.Lock()
    urls := c.filterPages(c.sitemap.getUnparsedPages())
    for _, url := range urls {
      c.fetchers = append(c.fetchers, newFetcher(url, c.sitemap))
    }
    c.sitemap.Unlock()

    if (len(urls) == 0) {
      return nil
    }
  }
  return nil
}

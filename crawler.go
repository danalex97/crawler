package crawler

type Crawler struct {
  domain   string
  fetchers map[string] *fetcher
  sitemap  *sitemap
}

func NewCrawler(domain string) *Crawler {
  c := new(Crawler)

  c.domain   = domain
  c.fetchers = make(map[string] *fetcher)
  c.sitemap  = newSitemap()

  c.fetchers[domain] = newFetcher(domain, c.sitemap)

  return c
}

func (c *Crawler) Run() error {
  for _, fetcher := range c.fetchers {
    fetcher.fetch()
  }
  return nil
}

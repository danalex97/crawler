package crawler

type fetcher struct {
  url     string
  sitemap *sitemap
}

func newFetcher(url string, sitemap *sitemap) *fetcher {
  c := new(fetcher)

  c.url     = url
  c.sitemap = sitemap

  return c
}

func (f *fetcher) fetch() *page {
  return nil
}

package crawler

import (
  "net/http"
)

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

func (f *fetcher) fetch() (*page, error) {
  resp, err := http.Get(f.url)
  if err != nil {
    return nil, err;
  }

  parser := newParser(f.url, resp.Body)
  return parser.parse()
}

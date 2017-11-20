package crawler

import (
  "fmt"
  "net/http"
)

type fetcher struct {
  url     string
  sitemap *sitemap
  client  *http.Client
}

func newFetcher(url string, sitemap *sitemap) *fetcher {
  c := new(fetcher)

  c.url     = url
  c.sitemap = sitemap
  c.client  = &http.Client{}

  return c
}

func (f *fetcher) fetch() (*page, error) {
  fmt.Println("Fetching page", f.url)

  resp, err := f.client.Get(f.url)
  if err != nil {
    return nil, err;
  }

  parser := newParser(f.url, resp.Body)
  urls   := parser.parse()

  f.sitemap.Lock()
  page := f.sitemap.getPage(f.url)
  for _, ref := range urls {
    target := f.sitemap.getPage(ref)
    if target != nil {
      page.addLink(target)
    }
  }
  page.setParsed()
  f.sitemap.Unlock()

  return page, nil
}

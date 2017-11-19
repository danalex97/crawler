package crawler

import (
  "fmt";
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

  page := f.sitemap.getPage(f.url)
  for _, ref := range urls {
    page.addLink(f.sitemap.getPage(ref))
  }
  page.setParsed()

  return page, nil
}

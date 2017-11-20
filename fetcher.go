package crawler

import (
  "fmt"
  "net/http"
)

type fetcher struct {
  url     string
  client  *http.Client
}

func newFetcher(url string) *fetcher {
  c := new(fetcher)
  c.url     = url
  c.client  = &http.Client{}
  return c
}

func (f *fetcher) fetch() fetchedData {
  fmt.Println("Fetching page", f.url)

  resp, err := f.client.Get(f.url)
  if err != nil {
    return toFetchedData(f.url, nil, err)
  }

  parser := newParser(f.url, resp.Body)
  urls   := parser.parse()

  return toFetchedData(f.url, urls[:], nil)
}

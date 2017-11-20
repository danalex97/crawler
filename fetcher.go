package crawler

import (
  "fmt"
  "net/http"
)

type fetcher interface {
  fetch() fetchedData
}

type httpFetcher struct {
  url     string
  client  http.Client
}

func newHttpFetcher(url string, client http.Client) fetcher {
  c := new(httpFetcher)
  c.url     = url
  c.client  = client
  return c
}

func (f *httpFetcher) fetch() fetchedData {
  fmt.Println("Fetching page", f.url)

  resp, err := f.client.Get(f.url)
  if err != nil {
    return toFetchedData(f.url, []string{}, err)
  }

  parser := newParser(f.url, resp.Body)
  urls   := parser.parse()

  if urls == nil {
    return toFetchedData(f.url, []string{}, nil)
  }

  return toFetchedData(f.url, urls[:], nil)
}

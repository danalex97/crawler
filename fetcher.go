package crawler

import (
  "fmt"
  "net/http"
)

type fetcher interface {
  fetch() fetchedData
}

/* An http fetcher is a fetecher that uses and HTTP client to fetch data from
  a specific url. The fetcher does a GET at a specific url and resturns all
  the links found inside the page if the request was sucessful. */
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
    return toFetchedData(f.url, []string{}, []string{}, err)
  }

  // Even though resp.Body returns a ReadCloser, we can cast it to Reader
  parser       := newParser(f.url, resp.Body)
  urls, assets := parser.parse()

  fmt.Printf("Found %v assets for page %v\n", len(assets), f.url)

  if urls == nil {
    urls = []string{}
  }
  if assets == nil {
    assets = []string{}
  }

  return toFetchedData(f.url, urls, assets, nil)
}

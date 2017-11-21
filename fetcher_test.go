package crawler

import (
  "testing"
  "net/http"
  "github.com/stretchr/testify/assert"
  "gopkg.in/h2non/gock.v1"
)

func TestHttpFetcherGivesAListOfLinksForWellFormatedInput(t *testing.T) {
  defer gock.Off()

  gock.New("http://foo.com").
    Get("/bar").
    Reply(200).
    BodyString("<htl><a href=\"test\"><p><a href=\"test2\"></html>")

  fetcher := newHttpFetcher("http://foo.com/bar", http.Client{})
  assert.Equal(t, fetcher.fetch(), toFetchedData("http://foo.com/bar", []string{"test", "test2"}, nil))
}

func TestHttpFetcherGivesNothingBackForBadFormatedInput(t *testing.T) {
  defer gock.Off()

  gock.New("http://foo.com").
    Get("/bar").
    Reply(200).
    BodyString("<htl><a hef=\"test\"><p><a hef=\"test2\"></html>")

  fetcher := newHttpFetcher("http://foo.com/bar", http.Client{})
  assert.Equal(t, fetcher.fetch(), toFetchedData("http://foo.com/bar", []string{}, nil))
}

func TestHttpFetcherGivesErrorFor404Response(t *testing.T) {
  defer gock.Off()

  gock.New("http://foo.com").
    Get("/bar").
    Reply(404)

  fetcher   := newHttpFetcher("http://foo.com/bar", http.Client{})
  fetchData := toFetchedData("http://foo.com/bar", []string{}, nil)
  assert.Equal(t, fetcher.fetch().url, fetchData.url)
  assert.Equal(t, fetcher.fetch().urls, fetchData.urls)
}

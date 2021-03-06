package crawler

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "errors"
)

type mockFectcher struct {
}

func (f *mockFectcher) fetch() fetchedData {
  return toFetchedData("test", []string {"test/url1", "tst/url2"}, []string{}, nil)
}

func testCrawler() *Crawler {
  crawler := NewCrawler("test")
  crawler.fetchers = []fetcher {&mockFectcher{}, &mockFectcher{}, &mockFectcher{}}
  return crawler
}

func TestFetchingPagesAddsAllTheFetcherDataToChannel(t *testing.T) {
  crawler        := testCrawler()
  fetchedChannel := crawler.fetchPages()
  for _, f := range crawler.fetchers {
    data1 := f.fetch()
    data2 := <-fetchedChannel
    assert.Equal(t, data1.url, data2.url)
    assert.Equal(t, data1.urls, data2.urls)
    assert.Equal(t, data1.err, data2.err)
  }
}

func TestBuildPagesReturnsArraysOfFilteredUrlsAndErrorsAreIgnored(t *testing.T) {
  fetchedChannel := make(chan fetchedData, 3)

  fetchedChannel <- toFetchedData("test", []string {"test/url1", "tst/url2"}, []string{}, nil)
  fetchedChannel <- toFetchedData("test", nil, []string{}, errors.New("test error"))
  fetchedChannel <- toFetchedData("test", []string {"test/url1", "test/url2"}, []string{}, nil)

  crawler := testCrawler()
  buildingChannel := crawler.buildPages(fetchedChannel)
  assert.Equal(t, []string {"test/url1", "test/url2"}, <-buildingChannel)
  assert.Equal(t, []string {"test/url1"}, <-buildingChannel)
}

func TestBuildPagesBuildsAPageAndFiltersItsLinks(t *testing.T) {
  fetchedChannel := make(chan fetchedData, 1)
  fetchedChannel <- toFetchedData("test", []string {"test/url1", "tst/url2"}, []string{}, nil)

  crawler := NewCrawler("test")
  crawler.fetchers = []fetcher {&mockFectcher{}}
  buildingChannel := crawler.buildPages(fetchedChannel)
  <-buildingChannel

  page1 := crawler.sitemap.getPage("test")
  assert.Equal(t, page1.getUrl(), "test")
  page2 := crawler.sitemap.getPage("test/url1")

  assert.Equal(t, page1.getLinks(), []*page{page2})
}

func TestUpdateFetchersCratesTheCorrectNumberOfFetchers(t *testing.T) {
  crawler := testCrawler()

  buildingChannel := make(chan []string, 3)
  buildingChannel <- []string {"url1", "url2"}
  buildingChannel <- []string {"url1", "url3", "url7"}
  buildingChannel <- []string {"url3", "url2", "url8"}

  crawler.updateFetchers(buildingChannel)
  assert.Equal(t, len(crawler.fetchers), 5)
}

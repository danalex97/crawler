package crawler

import (
	"testing"
  "github.com/stretchr/testify/assert"
)

func TestCanAddNewPage(t *testing.T) {
  sitemap := newSitemap("test")
  page1   := newPage("test")
  page2   := newPage("test2")

  sitemap.addPage(page1)
  sitemap.addPage(page2)
  assert.Equal(t, sitemap.getPage("test"), page1)
  assert.Equal(t, sitemap.getPage("test2"), page2)
}

func TestGetPageCreatesNewUnparsedPages(t *testing.T) {
  sitemap := newSitemap("test")
  assert.Equal(t, sitemap.getPage("test").getUrl(), "test")
  assert.Equal(t, sitemap.getPage("test2").getUrl(), "test2")
}

func TestGetUnparsedPages(t *testing.T) {
  sitemap := newSitemap("test")
  page1 := sitemap.getPage("test")
  page2 := sitemap.getPage("test2")

  assert.Equal(t, sitemap.getUnparsedPages(), []*page{page1, page2})
  page1.setParsed()
  assert.Equal(t, sitemap.getUnparsedPages(), []*page{page2})
  page2.setParsed()
  assert.Equal(t, len(sitemap.getUnparsedPages()), 0)
}

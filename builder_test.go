package crawler

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFilterPagesElimiatesTrailingHashtags(t *testing.T) {
  sitemap := newSitemap("test")
  builder := newBuilder("test", sitemap)

  pages := []string{
    "test/p1#test",
    "test/p2",
    "test/p4#what",
  }

  assert.Equal(t, builder.filterPages(pages), []string{
    "test/p1",
    "test/p2",
    "test/p4",
  })
}

func TestFilterPagesAddesPrefixForLocalLinks(t *testing.T) {
  sitemap := newSitemap("test")
  builder := newBuilder("test", sitemap)

  pages := []string{
    "/p1#test",
    "test/p2",
    "/p4#what",
  }

  assert.Equal(t, builder.filterPages(pages), []string{
    "test/p1",
    "test/p2",
    "test/p4",
  })
}

func TestFilterPagesFilteresPagesOutsideDomain(t *testing.T) {
  sitemap := newSitemap("test")
  builder := newBuilder("test", sitemap)

  pages := []string{
    "test/p1",
    "test/p2",
    "kappa/p3",
    "test/p4",
  }

  assert.Equal(t, builder.filterPages(pages), []string{
    "test/p1",
    "test/p2",
    "test/p4",
  })
}

func TestBuildPageAddsAllLinks(t *testing.T) {
  sitemap := newSitemap("test")
  builder := newBuilder("test", sitemap)

  pages := []string{
    "test/p1",
    "test/p2",
    "test/p3",
  }

  pageUrl := "test/page"
  builder.buildPage(pageUrl, pages, []string{})
  page := sitemap.getPage(pageUrl)
  assert.Equal(t, page.getUrl(), pageUrl)
  assert.Equal(t, len(page.getLinks()), 3)
  for _, p := range pages {
    assert.Contains(t, page.getLinks(), sitemap.getPage(p))
  }
}

func TestBuildPageFiltersOutParsedPages(t *testing.T) {
  sitemap := newSitemap("test")
  builder := newBuilder("test", sitemap)

  pages := []string{
    "test/p1",
    "test/p2",
    "test/p4",
    "test/p3",
  }

  sitemap.addPage(newPage("test/p4"))
  sitemap.getPage("test/p4").setParsed()

  pageUrl := "test/page"
  builder.buildPage(pageUrl, pages, []string{})

  assert.Equal(t, builder.buildPage(pageUrl, pages, []string{}), []string{
    "test/p1",
    "test/p2",
    "test/p3",
  })
}

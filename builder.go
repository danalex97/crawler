package crawler

import (
  "strings"
)

type Builder struct {
  domain  string
  sitemap *sitemap
}

func newBuilder(domain string, sitemap *sitemap) *Builder {
  b := new(Builder)

  b.domain  = domain
  b.sitemap = sitemap

  return b
}

func (b *Builder) filterPages(urls []string) (filtered []string) {
  for _, url := range urls {
    if strings.HasPrefix(url, b.domain) {
      filtered = append(filtered, url)
    }
  }
  return
}

func (b *Builder) buildPage(url string, urls []string) []string {
  b.sitemap.Lock()
  defer b.sitemap.Unlock()

  page     := b.sitemap.getPage(url)
  newPages := []string{}
  for _, ref := range urls {
    target := b.sitemap.getPage(ref)
    if !target.getParsed() {
      newPages = append(newPages, target.getUrl())
    }
    page.addLink(target)
  }
  page.setParsed()

  return newPages
}

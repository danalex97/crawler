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
  /* Normalizes and filters the page links.
     More normalization steps can be found here:
     https://en.wikipedia.org/wiki/URL_normalization

     If more normalization steps appear we can
     add filter structures: https://goo.gl/jvUxRa    */
  for _, url := range urls {
    // If the link is a local one: e.g. /archive; append the domain
    if strings.HasPrefix(url, "/") {
      url = b.domain + url
    }

    // Trimms the ending # in a link
    lastHashtag := strings.LastIndex(url, "#")
    if lastHashtag != -1 {
      url = url[:lastHashtag]
    }

    // Filteres the external links (i.e. not starting with the domain name)
    if strings.HasPrefix(url, b.domain) {
      filtered = append(filtered, url)
    }
  }
  return
}

func (b *Builder) buildPage(url string, urls []string) []string {
  b.sitemap.Lock()
  defer b.sitemap.Unlock()

  // Builds the links for the page url. The pages are accessed all
  // using the sitemap.
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

package crawler

import (
  "encoding/json"
  "sync"
  "strings"
)

type sitemap struct {
  domain string
  pages  map[string]* page
  sync.Mutex
}

func newSitemap(domain string) (* sitemap) {
  m := new(sitemap)

  m.pages  = make(map[string] *page)
  m.domain = domain

  return m
}

func (m *sitemap) addPage(page *page) {
  m.pages[page.url] = page
}

func (m *sitemap) getPage(url string) *page {
  if !strings.HasPrefix(url, m.domain) {
    return nil
  }
  page, exists := m.pages[url]
  if !exists {
    m.pages[url] = newPage(url)
    return m.pages[url]
  }
  return page
}

func (m *sitemap) getUnparsedPages() (pages []*page) {
  for _, page := range m.pages {
    if !page.getParsed() {
      pages = append(pages, page)
    }
  }
  return
}

func (m *sitemap) toJson() string {
  links := []map[string]string{}
  for _, x := range m.pages {
    for _, y := range x.getLinks() {
      link := map[string]string{
        "source": x.getUrl(),
        "target": y.getUrl(),
      }
      links = append(links, link)
    }
  }
  ans, _ := json.Marshal(links)
  return string(ans)
}

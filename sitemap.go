package crawler

import (
  "encoding/json";
  "sync"
)

type sitemap struct {
  pages map[string]* page
  sync.Mutex
}

func newSitemap() (* sitemap) {
  m := new(sitemap)

  m.pages = make(map[string] *page)

  return m
}

func (m *sitemap) addPage(page *page) {
  m.pages[page.url] = page
}

func (m *sitemap) getPage(url string) *page {
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
  links := []string{}
  for _, x := range m.pages {
    for _, y := range x.getLinks() {
      link := map[string]string{
        "source": x.getUrl(),
        "target": y.getUrl(),
      }
      serialized, _ := json.Marshal(link)
      links = append(links, string(serialized))
    }
  }
  ans, _ := json.Marshal(links)
  return string(ans)
}

package crawler

type sitemap struct {
  pages map[string]* page;
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

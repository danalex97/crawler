package crawler

type sitemap struct {
  pages map[string]* page;
}

func newSitemap() (* sitemap) {
  m := new(sitemap)

  m.pages = make(map[string] *page)

  return m
}

func (m *sitemap) newSite(page *page) {
  m.pages[page.url] = page
}

func (m *sitemap) getPage(url string) *page {
  page, exists := m.pages[url]
  if !exists {
    return newPage(url)
  }
  return page
}

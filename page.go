package crawler

type page struct {
  url    string
  parsed bool
  links  map[string]* page
}

func newPage(url string) *page {
  p := new(page)

  p.url   = url
  p.links = make(map[string] *page)
  p.parsed = false

  return p
}

func (p *page) addLink(link *page) {
  p.links[link.url] = link
}

func (p *page) getLinks() ([]*page) {
  pages := []*page {}
  for _, page := range p.links {
    pages = append(pages, page)
  }
  return pages
}

func (p *page) setParsed() {
  p.parsed = true
}

func (p *page) getParsed() bool {
  return p.parsed
}

func (p *page) getUrl() string {
  return p.url
}

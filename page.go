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

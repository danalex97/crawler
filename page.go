package crawler

type page struct {
  url string
  links map[string]* page
}

func newPage(url string) *page {
  p := new(page)

  p.url   = url
  p.links = make(map[string] *page)

  return p
}

func (p *page) addLink(link *page) {
  p.links[link.url] = link
}

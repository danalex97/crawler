package crawler

/* The page struct encapsulates the data needed for a page.
   It can be extending to remember the content or any other data. */
type page struct {
  url    string
  parsed bool
  links  map[string]* page
  assets []string
}

func newPage(url string) *page {
  p := new(page)

  p.url   = url
  p.links = make(map[string] *page)
  p.parsed = false
  p.assets = []string{}

  return p
}

func (p *page) addAssets(assets []string) {
  for _, asset := range assets{
    p.assets = append(p.assets, asset)
  }
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

/* We parse a page only once, so don't want to allow
  setting this to false, since the crawler renounces when
  it can't parse a page. */
func (p *page) setParsed() {
  p.parsed = true
}

func (p *page) getParsed() bool {
  return p.parsed
}

func (p *page) getUrl() string {
  return p.url
}

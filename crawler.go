package crawler

type Crawler struct {

}

func NewCrawler() *Crawler {
  return new(Crawler)
}

func (c *Crawler) Run() error {
  return nil
}

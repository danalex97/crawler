package crawler

import (
	"testing"
  "github.com/stretchr/testify/assert"
)

func TestFilterPagesFilteresPagesOutsideDomain(t *testing.T) {
  crawler := NewCrawler("test")

  pages := []*page {
    newPage("test/p1"),
    newPage("test/p2"),
    newPage("kappa/p3"),
    newPage("test/p4"),
  }

  assert.Equal(t, crawler.filterPages(pages), []string{
    "test/p1",
    "test/p2",
    "test/p4",
  })
}

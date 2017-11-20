package crawler

import (
	"testing"
  "github.com/stretchr/testify/assert"
)

func TestFilterPagesFilteresPagesOutsideDomain(t *testing.T) {
	sitemap := newSitemap("test")
	builder := newBuilder("test", sitemap)

  pages := []string{
    "test/p1",
    "test/p2",
		"kappa/p3",
    "test/p4",
  }

  assert.Equal(t, builder.filterPages(pages), []string{
    "test/p1",
    "test/p2",
    "test/p4",
  })
}

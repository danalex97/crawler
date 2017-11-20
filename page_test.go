package crawler

import (
	"testing";
  "github.com/stretchr/testify/assert"
)

func TestNewPageHasUrlAndIsNotParsed(t *testing.T) {
  page := newPage("test")
  assert.Equal(t, page.getUrl(), "test")
  assert.Equal(t, page.getParsed(), false)
}

func TestNewPageHasNoLinks(t *testing.T) {
  page := newPage("test")
  assert.Equal(t, len(page.getLinks()), 0)
}

func TestCanAddLinks(t *testing.T) {
  page1 := newPage("test")
  page2 := newPage("test2")
  page3 := newPage("test3")
  assert.Equal(t, len(page1.getLinks()), 0)

  page1.addLink(page2)
  assert.Equal(t, len(page1.getLinks()), 1)
  assert.Equal(t, page1.getLinks(), []*page{page2})

  page1.addLink(page3)
  assert.Equal(t, len(page1.getLinks()), 2)
  assert.Equal(t, page1.getLinks(), []*page{page2, page3})
}

func TestSetParsedModifiesParsed(t *testing.T) {
  page := newPage("test")
  assert.Equal(t, page.getParsed(), false)
  page.setParsed()
  assert.Equal(t, page.getParsed(), true)
  page.setParsed()
  assert.Equal(t, page.getParsed(), true)
}

package crawler

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "strings"
)

func TestCanParseFromWellFormatedResponse(t *testing.T) {
  reader := strings.NewReader(
    "<html><a href=\"test\"></a><p></p></html>")
  parser  := newParser("test", reader)
  urls, _ := parser.parse()
  assert.Equal(t, urls, []string{"test"})
}

func TestCanParseFromBadlyFormatedResponse(t *testing.T) {
  reader := strings.NewReader(
    "<htl><a href=\"test\"><p></html>")
  parser  := newParser("test", reader)
  urls, _ := parser.parse()
  assert.Equal(t, urls, []string{"test"})
}

func TestCanParseMultipleTags(t *testing.T) {
  reader := strings.NewReader(
    "<htl><a href=\"test\"><p><a href=\"test2\"></html>")
  parser  := newParser("test", reader)
  urls, _ := parser.parse()
  assert.Equal(t, urls, []string{"test", "test2"})
}

func TestParseNothingWhenNoHrefNoAsssetsPresent(t *testing.T) {
  reader := strings.NewReader(
    "<htl><a><p><a></html>")
  parser       := newParser("test", reader)
  urls, assets := parser.parse()
  assert.Equal(t, len(urls), 0)
  assert.Equal(t, len(assets), 0)
}

func TestParserGetsCorrectAssetsForWellFormattedInput(t *testing.T) {
  reader := strings.NewReader(
    "<htl><a><img src=\"danimocanu\"><p><a><link href=\"danimocanu\"></html>")
  parser    := newParser("test", reader)
  _, assets := parser.parse()
  assert.Equal(t, assets, []string{"danimocanu", "danimocanu"})
}

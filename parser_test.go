package crawler

import (
	"testing"
  "github.com/stretchr/testify/assert"
  "strings"
)

func TestCanParseFromWellFormatedResponse(t *testing.T) {
  reader := strings.NewReader(
    "<html><a href=\"test\"></a><p></p></html>")
  parser := newParser("test", reader)
  assert.Equal(t, parser.parse(), []string{"test"})
}

func TestCanParseFromBadlyFormatedResponse(t *testing.T) {
  reader := strings.NewReader(
    "<htl><a href=\"test\"><p></html>")
  parser := newParser("test", reader)
  assert.Equal(t, parser.parse(), []string{"test"})
}

func TestCanParseMultipleTags(t *testing.T) {
  reader := strings.NewReader(
    "<htl><a href=\"test\"><p><a href=\"test2\"></html>")
  parser := newParser("test", reader)
  assert.Equal(t, parser.parse(), []string{"test", "test2"})
}

func TestParseNothingWhenNoHrefPresent(t *testing.T) {
  reader := strings.NewReader(
    "<htl><a><p><a></html>")
  parser := newParser("test", reader)
  assert.Equal(t, len(parser.parse()), 0)
}

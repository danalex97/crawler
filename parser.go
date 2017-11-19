package crawler

import (
  "golang.org/x/net/html";
  "io";
  "fmt"
)

type parser struct {
  reader io.Reader
}

func (p *parser) parse() (*page, error) {
  tokenizer := html.NewTokenizer(p.reader)
  for {
    token := tokenizer.Next()
    switch {
    case token == html.ErrorToken:
      return nil, nil
    case token == html.StartTagToken:
      token := tokenizer.Token()
      if token.Data == "a" {
        fmt.Println(token.Data)
      }
    }
  }
  return nil, nil
}

func newParser(reader io.Reader) *parser {
  p := new(parser)
  p.reader = reader
  return p
}

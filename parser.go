package crawler

import (
  "golang.org/x/net/html";
  "io";
  "fmt"
)

type parser struct {
  reader io.Reader
}

func getElement(token html.Token, element string) (ok bool, href string) {
  for _, attribute := range token.Attr {
    if attribute.Key == element {
      href = attribute.Val
      ok   = true
      return
    }
  }
  return
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

      // We found start of <a> tag
      if token.Data == "a" {
        ok, url := getElement(token, "href")
        if ok {
          // Found new url
          fmt.Printf("New url found %v\n", url)
        }
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

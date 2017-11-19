package crawler

import (
  "golang.org/x/net/html";
  "io";
  "fmt"
)

type parser struct {
  url    string
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

func getUrls(tokenizer *html.Tokenizer) (urls []string) {
  for {
    token := tokenizer.Next()
    switch {
    case token == html.ErrorToken:
      return
    case token == html.StartTagToken:
      token := tokenizer.Token()

      // We found start of <a> tag
      if token.Data == "a" {
        ok, url := getElement(token, "href")
        if ok {
          urls = append(urls, url)
        }
      }
    }
  }
}

func (p *parser) parse() (*page, error) {
  tokenizer := html.NewTokenizer(p.reader)
  urls      := getUrls(tokenizer)

  for _, url := range urls {
    fmt.Printf("New url found %v\n", url)
  }

  return nil, nil
}

func newParser(url string, reader io.Reader) *parser {
  p := new(parser)

  p.url    = url
  p.reader = reader

  return p
}

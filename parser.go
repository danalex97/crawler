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

func (p *parser) parse() []string {
  tokenizer := html.NewTokenizer(p.reader)
  urls      := getUrls(tokenizer)

  for _, url := range urls {
    fmt.Printf("Url found: %v\n", url)
  }

  return urls
}

func newParser(url string, reader io.Reader) *parser {
  p := new(parser)

  p.url    = url
  p.reader = reader

  return p
}

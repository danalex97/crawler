package crawler

import (
  "golang.org/x/net/html"
  "io"
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

/* The Parser receives a Reader stream and tokenizes the input.
   When we arrive at a tag token, if the element is <a>, we retrieve the
   "href" attribute key.

   This can be extended to processing static resources or other tags.
 */
func (p *parser) parse() (urls []string) {
  tokenizer := html.NewTokenizer(p.reader)
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

func newParser(url string, reader io.Reader) *parser {
  p := new(parser)

  p.url    = url
  p.reader = reader

  return p
}

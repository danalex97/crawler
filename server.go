package crawler

import (
	"fmt"
	"log"
	"net/http"
  "strconv"
)

type Server struct {
  sitemap  *sitemap
}

func NewServer(c *Crawler) *Server {
  s := new(Server)
  s.sitemap = c.sitemap
  return s
}

func (s *Server) handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
  	w.WriteHeader(http.StatusOK)

    s.sitemap.Lock()
    data := s.sitemap.toJson()
    s.sitemap.Unlock()

    w.Header().Set("Content-Length", fmt.Sprint(len(data)))
    fmt.Fprint(w, string(data))
  }
}

func (s *Server) ServerListenPathPort(path string, port int) {
	http.HandleFunc(path, s.handler())
	log.Fatal(http.ListenAndServe(":" +  strconv.Itoa(port), nil))
}

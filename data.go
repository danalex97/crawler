package crawler

/* The fetched data is what arrives from a fetcher. */
type fetchedData struct {
  url  string
  urls []string
  err  error
}

func toFetchedData(url string, urls []string, err error) (f fetchedData) {
  f.url  = url
  f.urls = urls
  f.err  = err
  return
}

func fromFetchedData(f fetchedData) (url string, urls []string, err error) {
  url  = f.url
  urls = f.urls
  err  = f.err
  return
}

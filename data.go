package crawler

/* The fetched data is what arrives from a fetcher. */
type fetchedData struct {
  url    string
  urls   []string
  assets []string
  err    error
}

func toFetchedData(url string, urls []string, assets []string, err error) (f fetchedData) {
  f.url    = url
  f.urls   = urls
  f.assets = assets
  f.err    = err
  return
}

func fromFetchedData(f fetchedData) (url string, urls []string, assets []string, err error) {
  url    = f.url
  urls   = f.urls
  assets = f.assets
  err    = f.err
  return
}

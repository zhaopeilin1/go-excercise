package chans

type WebServiceCheck func(string) bool

type Result struct {
	url string
	ok  bool
}

func testUrls(c WebServiceCheck, list []string) map[string]bool {
	m := map[string]bool{}

	ch := make(chan Result)

	for _, url := range list {
		go func(u string) {
			ch <- Result{url: u, ok: c(u)}
		}(url)
	}
	for i := 0; i < len(list); i++ {
		r := <-ch
		m[r.url] = r.ok
	}
	return m
}

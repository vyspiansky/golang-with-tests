package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// To tell Go to start a new goroutine we turn a function call into a go statement
		// by putting the keyword go in front of it.
		// We often use anonymous functions when we want to start a goroutine.
		go func() {
			// results[url] = wc(url)

			// Send statement
			// resultChannel <- result{u, wc(u)}

			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

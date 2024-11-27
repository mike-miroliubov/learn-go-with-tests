package concurrency

type WebsiteChecker func(string) bool
type websiteCheckResult struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChan := make(chan websiteCheckResult)

	for _, url := range urls {
		go func() {
			resultsChan <- websiteCheckResult{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChan // await
		results[r.string] = r.bool
	}

	return results
}

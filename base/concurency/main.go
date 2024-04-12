package main

func main() {

}

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result, len(urls))

	for _, url := range urls {
		go func(u string) {
			resultChan <- result{
				u,
				wc(u),
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		results[r.string] = r.bool
	}

	return results
}

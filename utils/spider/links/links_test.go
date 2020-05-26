package links

import "testing"

func TestBreadthFirst(t *testing.T) {
	urls := []string{"https://www.jianshu.com/p/47691d870756"}

	BreadthFirst(ReadWithError, urls)
}

func TestConcurrentRead(t *testing.T) {
	url := "https://www.jianshu.com/p/47691d870756"

	ConcurrentRead(ReadWithError, url)
}

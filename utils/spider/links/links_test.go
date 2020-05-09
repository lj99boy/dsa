package links

import "testing"

func TestBreadthFirst(t *testing.T) {
	urls := []string{"https://www.jianshu.com/p/47691d870756"}

	BreadthFirst(ReadWithError,urls)
}

package channel

type Ints []int

func (ints Ints) Iterator() <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range ints {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

package writealternately

func WriteOdd() {
	tag := make(chan int, 1)
	tag <- 1
	holder := make(chan int, 100)

	go func() {
		for i := 1; i <= 100; {
			temp := <-tag
			if temp == 1 {
				holder <- i
				i = i + 2
				tag <- 2
			} else {
				tag <- temp
			}
		}
	}()

	go func() {
		for i := 2; i <= 100; {
			temp := <-tag
			if temp == 2 {
				holder <- i
				i = i + 2
				tag <- 1
			} else {
				tag <- temp
			}
		}
	}()

	for val := range holder {
		println(val)
	}
}

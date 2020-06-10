package alog

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int, positive bool) (slice []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if positive {
			slice = append(slice, rand.Intn(999))
		} else {
			slice = append(slice, rand.Intn(999)-rand.Intn(999))
		}
	}
	return
}

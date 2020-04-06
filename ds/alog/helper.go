package alog

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) (slice []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(999)-rand.Intn(999))
	}
	return
}

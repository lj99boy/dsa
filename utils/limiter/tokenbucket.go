package limiter

import (
	"fmt"
	"time"
)

const bucketCapacity = 10
const fillInterval = 1

type TokenBucket struct {
	bucket chan int
}

var tb *TokenBucket

func init() {
	tb = &TokenBucket{
		bucket: make(chan int, bucketCapacity),
	}
}

//fu

func TakeAvailable(block bool) bool {
	var takeRes bool
	if block {
		select {
		case <-tb.bucket:
			takeRes = true
		}
	} else {
		select {
		case <-tb.bucket:
			takeRes = true
		default:
			takeRes = false
		}
	}
	return takeRes
}

func Start() {
	ticker := time.NewTicker(time.Duration(time.Second * fillInterval))
	for {
		select {
		case <-ticker.C:
			select {
			case tb.bucket <- 1:
			default:
			}
			fmt.Println("current token cnt:", len(tb.bucket), time.Now())
		}
	}
}

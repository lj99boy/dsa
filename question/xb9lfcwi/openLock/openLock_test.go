package openLock

import (
	"fmt"
	"testing"
)

func Test_getNowCanReachLock(t *testing.T) {
	fmt.Println(getNowCanReachLock("0000"))
}

func Test_openLock(t *testing.T) {
	d := []string{"8887","8889","8878","8898","8788","8988","7888","9888"}
	openLock(d,"8888")
}
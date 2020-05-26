package limiter

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestStart(t *testing.T) {
	go Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

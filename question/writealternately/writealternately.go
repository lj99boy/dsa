package writealternately

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func WriteAlternately() {
	tag := make(chan bool)

	go func() {
		i := 1
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-tag:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				tag <- true
				break
			default:
				break
			}
		}
	}()
	go func() {
		str := "ABCDEHDQOWDHUQHDUWQID"
		i := 0
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-tag:
				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "")-1 {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				tag <- true
				break
			default:
				break
			}
		}
	}()
	tag <- true
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

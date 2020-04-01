package writealternately

import (
	"os"
	"os/signal"
)

func WriteAlternately() {
	closeSignal := make(chan os.Signal, 1)

	signal.Notify(closeSignal,os.Signal())
}

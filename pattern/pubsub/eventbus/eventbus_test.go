package eventbus

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var eb = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

func TestEventBus_Publish(t *testing.T) {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)

	eb.Subscribe("t1", ch1)
	eb.Subscribe("t2", ch2)
	eb.Subscribe("t3", ch3)

	go publishTo("t1", "send t1")
	go publishTo("t2", "send t2")

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch1", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}

func publishTo(topic string, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

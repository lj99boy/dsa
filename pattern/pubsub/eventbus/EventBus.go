package eventbus

import "sync"

// EventBus stores the information about subscribers interested for
// a particular topic
type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

//订阅，订阅者带着自己的碗（datachannel 管道）过来订阅某个topic，以后发布者消息发到碗里
func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	defer eb.rm.Unlock()

	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append(DataChannelSlice{}, ch)
	}
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	defer eb.rm.RUnlock()

	if chans, found := eb.subscribers[topic]; found {
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
}

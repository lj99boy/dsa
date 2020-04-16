package eventbus

//DataChannel is a channel which can accept an DataEvent
type DataChannel chan DataEvent

//DataChannelSlice is a slice of DataChannels
type DataChannelSlice [] DataChannel

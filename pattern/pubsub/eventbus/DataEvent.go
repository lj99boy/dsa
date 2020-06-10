package eventbus

//DataEvent is underlying data
type DataEvent struct {
	Data interface{}
	Topic string
}

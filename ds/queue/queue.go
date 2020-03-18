package queue

type Queue struct {
	queue []interface{}
	len int
	lock
}
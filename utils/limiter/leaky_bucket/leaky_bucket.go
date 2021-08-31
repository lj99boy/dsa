package leaky_bucket

type Bucket struct {
	Capacity int
	Drops int
	LastTime time
}

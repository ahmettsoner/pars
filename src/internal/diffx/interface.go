package diffx

type Comparable[T any] interface {
	Key() string
	IsEqual(other T) bool
}

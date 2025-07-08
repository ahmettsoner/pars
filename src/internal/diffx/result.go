package diffx

type DiffResult[T any] struct {
	Created []T
	Updated []struct{ Old, New T }
	Deleted []T
}

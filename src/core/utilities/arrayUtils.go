package utilities

func Reverse[T any](s []T) []T {
	n := len(s)
	reversed := make([]T, n)
	for i := 0; i < n; i++ {
		reversed[i] = s[n-1-i]
	}
	return reversed
}

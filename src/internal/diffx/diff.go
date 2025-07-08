package diffx

func DiffSlice[T Comparable[T]](oldList, newList []T) DiffResult[T] {
	oldMap := make(map[string]T)
	newMap := make(map[string]T)
	result := DiffResult[T]{}

	for _, o := range oldList {
		oldMap[o.Key()] = o
	}

	for _, n := range newList {
		newMap[n.Key()] = n
		old, exists := oldMap[n.Key()]
		if !exists {
			result.Created = append(result.Created, n)
		} else if !n.IsEqual(old) {
			result.Updated = append(result.Updated, struct{ Old, New T }{Old: old, New: n})
		}
	}

	for _, o := range oldList {
		if _, exists := newMap[o.Key()]; !exists {
			result.Deleted = append(result.Deleted, o)
		}
	}

	return result
}

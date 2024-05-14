package Slice

func RemoveFromSlice[T comparable](value T, slice []T) []T {
	if len(slice) > 1 {
		new := make([]T, len(slice))
		var n = 0
		for _, v := range slice {
			if v != value {
				new[n] = v
				n++
			}
		}
		return new[:n]
	}
	return []T{}
}

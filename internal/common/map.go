package common

func ToFieldMap[T any, K comparable](list []*T, keyFunc func(*T) K) map[K]*T {
	result := make(map[K]*T)
	for _, v := range list {
		result[keyFunc(v)] = v
	}
	return result
}

func MapList[T, K any](list []*T, mapFunc func(*T) K) []K {
	result := make([]K, len(list))
	for i, v := range list {
		result[i] = mapFunc(v)
	}
	return result
}

func MapListDistinct[T, K comparable](list []*T, mapFunc func(*T) K) []K {
	set := make(map[K]K)
	for _, v := range list {
		k := mapFunc(v)
		if _, ok := set[k];!ok {
			set[k] = k
		}
	}
	result := make([]K, len(set))
	for _, v := range set {
		result = append(result, v)
	}
	return result
}



package utils

func ConvertMap[T comparable, K any](m map[T]K) []K {
	values := make([]K, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}

	return values
}

func ConvertMapPointer[K comparable, V any, R any](m map[K]*V, mapper PointerMapperFunc[V, R]) []*R {
	values := make([]*R, 0, len(m))
	for _, value := range m {
		tmp := value
		values = append(values, mapper(tmp))
	}

	return values
}

type PointerMapperFunc[T, R any] func(t *T) *R

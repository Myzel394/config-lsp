package utils

func Group[T any, V comparable](values []T, f func(T) V) map[V][]T {
	result := make(map[V][]T)

	for _, value := range values {
		key := f(value)
		result[key] = append(result[key], value)
	}

	return result
}

func Map[T any, O any](values []T, f func(T) O) []O {
	result := make([]O, len(values))

	for index, value := range values {
		result[index] = f(value)
	}

	return result
}

func MapMap[T comparable, O any, P any](values map[T]O, f func(T, O) P) map[T]P {
	result := make(map[T]P)

	for key, value := range values {
		result[key] = f(key, value)
	}

	return result
}

func MapMapToSlice[T comparable, O any, P any](values map[T]O, f func(T, O) P) []P {
	result := make([]P, 0)

	for key, value := range values {
		result = append(result, f(key, value))
	}

	return result
}

func SliceToSet[T comparable](values []T) map[T]struct{} {
	set := make(map[T]struct{})

	for _, value := range values {
		set[value] = struct{}{}
	}

	return set
}

func SliceToMap[T comparable, O any](values []T, defaultValue O) map[T]O {
	set := make(map[T]O)

	for _, value := range values {
		set[value] = defaultValue
	}

	return set
}

func FilterWhere[T any](values []T, f func(T) bool) []T {
	result := make([]T, 0)

	for _, value := range values {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}

func FilterMap[T comparable, O any](
	values map[T]O,
	f func(T, O) bool,
) map[T]O {
	result := make(map[T]O)

	for key, value := range values {
		if f(key, value) {
			result[key] = value
		}
	}

	return result
}

func FilterMapWhere[T comparable, O any](values map[T]O, f func(T, O) bool) map[T]O {
	result := make(map[T]O)

	for key, value := range values {
		if f(key, value) {
			result[key] = value
		}
	}

	return result
}

func KeysAsSet[T comparable, O any](values map[T]O) map[T]struct{} {
	set := make(map[T]struct{})

	for key := range values {
		set[key] = struct{}{}
	}

	return set
}

func KeysOfMap[T comparable, O any](values map[T]O) []T {
	keys := make([]T, 0)

	for key := range values {
		keys = append(keys, key)
	}

	return keys
}

func ValuesOfMap[T comparable, O any](values map[T]O) []O {
	keys := make([]O, 0)

	for _, value := range values {
		keys = append(keys, value)
	}

	return keys
}

func MergeMaps[T comparable, O any](maps ...map[T]O) map[T]O {
	result := make(map[T]O)

	for _, m := range maps {
		for key, value := range m {
			result[key] = value
		}
	}

	return result
}

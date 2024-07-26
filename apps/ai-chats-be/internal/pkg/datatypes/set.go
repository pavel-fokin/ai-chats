package datatypes

// Set is a generic type that represents a set of elements.
// It is implemented as a map where the keys are the elements of type T and the values are empty structs.
// The Set type supports operations such as adding elements, removing elements, and checking if an element is present in the set.
type Set[T comparable] map[T]struct{}

// HashElemFunc is a function type that maps an element of type T to a key of type K.
// It is used in hash-based data structures to determine the key for each element.
// The key must be a comparable type.
type HashElemFunc[T any, K comparable] func(T) K

// NewSet creates a new set with the specified values and hash element function.
// The values parameter is a slice of elements of type T that will be added to the set.
// The fn parameter is a hash element function that takes an element of type T and returns a key of type K.
// The hash element function is used to determine the key for each element in the set.
// The returned set is a map with keys of type K and empty struct values.
func NewSet[T any, K comparable](values []T, fn HashElemFunc[T, K]) Set[K] {
	set := make(Set[K], len(values))
	for _, value := range values {
		set[fn(value)] = struct{}{}
	}
	return set
}

// Add adds a value to the set.
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Contains checks if the set contains the specified value.
// It returns true if the value is found in the set, otherwise it returns false.
func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

// Delete removes the specified value from the set.
func (s Set[T]) Delete(value T) {
	delete(s, value)
}

package set

// Set is a set of elements.
type Set[T comparable] struct {
	data map[T]struct{}
}

// New creates a new set.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

// NewWithMaxSize creates a new set with a maximum size.
func NewWithMaxSize[T comparable](maxSize int) *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}, maxSize),
	}
}

// NewWithValues creates a new set with the given values.
func NewWithValues[T comparable](values ...T) *Set[T] {
	s := NewWithMaxSize[T](len(values))
	for _, value := range values {
		s.Add(value)
	}
	return s
}

// Add adds a value to the set.
func (s *Set[T]) Add(value T) {
	s.data[value] = struct{}{}
}

// Contains returns true if the set contains the value.
func (s *Set[T]) Contains(value T) bool {
	_, ok := s.data[value]
	return ok
}

// Remove removes a value from the set.
func (s *Set[T]) Remove(value T) {
	delete(s.data, value)
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// Values returns a slice containing all the elements in the set.
func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.data))
	for value := range s.data {
		values = append(values, value)
	}
	return values
}

// Union returns a new set containing all the elements in either s or other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := New[T]()
	for value := range s.data {
		union.Add(value)
	}
	for value := range other.data {
		union.Add(value)
	}
	return union
}

// Intersection returns a new set containing all the elements in both s and other.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersection := New[T]()
	for value := range s.data {
		if other.Contains(value) {
			intersection.Add(value)
		}
	}
	return intersection
}

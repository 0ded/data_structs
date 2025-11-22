package set

// Set is a simple set implementation
type Set struct {
	elements map[interface{}]struct{}
}

// NewSet creates a new Set
func NewSet() *Set {
	return &Set{
		elements: make(map[interface{}]struct{}),
	}
}

// Add inserts an element into the set
func (s *Set) Add(value interface{}) {
	s.elements[value] = struct{}{}
}

// Remove deletes an element from the set
func (s *Set) Remove(value interface{}) {
	delete(s.elements, value)
}

// Contains checks if an element is in the set
func (s *Set) Contains(value interface{}) bool {
	_, exists := s.elements[value]
	return exists
}

// Size returns the number of elements in the set
func (s *Set) Size() int {
	return len(s.elements)
}

// Elements returns all elements in the set as a slice
func (s *Set) Elements() []interface{} {
	keys := make([]interface{}, 0, len(s.elements))
	for k := range s.elements {
		keys = append(keys, k)
	}
	return keys
}

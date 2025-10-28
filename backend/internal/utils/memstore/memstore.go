package memstore

import "sync"

type Store[V any] struct {
	mu   sync.RWMutex
	data V
	set  bool
}

// Creates a new memory store
func New[V any]() *Store[V] {
	return &Store[V]{}
}

// Sets the data for a memory store
func (s *Store[V]) Set(v V) {
	s.mu.Lock()
	s.data = v
	s.set = true
	s.mu.Unlock()
}

// Gets the data from a memory store
func (s *Store[V]) Get() (V, bool) {
	s.mu.RLock()
	v, ok := s.data, s.set
	s.mu.RUnlock()
	return v, ok
}

// Clears the data from a memory store
func (s *Store[V]) Clear() {
	s.mu.Lock()
	var zero V
	s.data = zero
	s.set = false
	s.mu.Unlock()
}

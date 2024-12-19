package structure

// Basic HashSet implementation
type HashSet[T comparable] struct {
	entries map[T]bool
}

func NewHashSet[T comparable]() *HashSet[T] {
	hs := HashSet[T]{
		entries: make(map[T]bool),
	}
	return &hs
}

func (h *HashSet[T]) Add(e T) {
	h.entries[e] = true
}

func (h *HashSet[T]) AddFromSlice(s []T) {
	for _, e := range s {
		h.Add(e)
	}
}

func (h *HashSet[T]) AddAll(es ...T) {
	for _, e := range es {
		h.Add(e)
	}
}

func (h *HashSet[T]) Size() int {
	return len(h.entries)
}

func (h *HashSet[T]) Remove(e T) {
	delete(h.entries, e)
}

func (h *HashSet[T]) Contains(e T) bool {
	return h.entries[e]
}

// shallow clone; elements are not cloned.
func (h *HashSet[T]) Clone() *HashSet[T] {
	hNew := HashSet[T]{
		entries: make(map[T]bool, h.Size()),
	}

	for k, v := range h.entries {
		hNew.entries[k] = v
	}

	return &hNew

}

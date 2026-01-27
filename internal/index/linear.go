package index

import (
	v "VectorDatabase/internal/vector"
)

type LinearIndex struct {
}

func NewLinearIndex() *LinearIndex {
	return &LinearIndex{}
}

func (*LinearIndex) Add(vec *v.Vector) error {
	return nil
}
func (*LinearIndex) Delete(id string) error {
	return nil
}
func (*LinearIndex) Get(id string) (*v.Vector, bool) {
	return &v.Vector{}, false
}
func (*LinearIndex) Search(query *v.Vector, k int) ([]SearchResult, error) {
	return []SearchResult{}, nil
}
func (*LinearIndex) Size() int {
	return 0
}

var _ VectorIndex = (*LinearIndex)(nil)

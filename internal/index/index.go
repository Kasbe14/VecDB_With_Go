package index

import (
	v "VectorDatabase/internal/vector"
)

type SearchResult struct {
	Vector *v.Vector
	Score  float64
}

type VectorIndex interface {
	Add(v *v.Vector) error
	Delete(id string) error
	Get(id string) (*v.Vector, bool)
	Search(query *v.Vector, k int) ([]SearchResult, error)
	Size() int
}

package index

import (
	v "VectorDatabase/internal/vector"
)

type VectorIndex interface {
	Add(v *v.Vector) error
	Delete(id string) error
	Get(id string) (*v.Vector, bool)
	Search(query *v.Vector, k int) ([]SearchResult, error)
	Size() int
}

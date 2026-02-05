package index

import (
	v "VectorDatabase/internal/vector"
)

//type DataType v.DataType
//type SimilarityMetric v.SimilarityMetric

type VectorIndex interface {
	Add(v *v.Vector) error
	Delete(id string) error
	Get(id string) (*v.Vector, bool)
	Search(query *v.Vector, k int) ([]SearchResult, error)
	Size() int
}

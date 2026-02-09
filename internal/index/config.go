package index

import (
	v "VectorDatabase/internal/vector"
)

type IndexConfig struct {
	IndexType IndexType
	DataType  v.DataType
	Metric    v.SimilarityMetric
	Dimension int
}

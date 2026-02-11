package index

import (
	"VectorDatabase/internal/types"
	"errors"
)

type IndexConfig struct {
	indexType types.IndexType
	modelType types.ModelType
	dataType  types.DataType
	metric    types.SimilarityMetric
	dimension int
}

// IndexConfig constructor with invariants checks
func NewIndexConfig(
	indexType types.IndexType,
	modelType types.ModelType,
	dataType types.DataType,
	metric types.SimilarityMetric,
	dimension int,
) (IndexConfig, error) {
	if dimension <= 0 {
		return IndexConfig{}, errors.New("invalid dimension")
	}
	switch indexType {
	case types.LinearIndex, types.HNSWIndex, types.IVFIndex, types.PQIndex:
		//ok valid input
	default:
		return IndexConfig{}, errors.New("invalid index type")
	}
	switch dataType {
	case types.Text, types.Audio, types.Video, types.Image:
		//ok valid input
	default:
		return IndexConfig{}, errors.New("invalid data type")
	}
	switch metric {
	case types.Cosine, types.Dot, types.Euclidean:
		//ok valid metrics
	default:
		return IndexConfig{}, errors.New("invalid metric type")
	}
	//==================TODO: Add models types============================
	switch modelType {
	case types.Testmodel:
		//ok
	default:
		return IndexConfig{}, errors.New("invalid model type")
	}
	return IndexConfig{
		indexType: indexType,
		modelType: modelType,
		dataType:  dataType,
		metric:    metric,
		dimension: dimension,
	}, nil

}

//getters for IndexConfig

func (c IndexConfig) IndexType() types.IndexType     { return c.indexType }
func (c IndexConfig) ModelType() types.ModelType     { return c.modelType }
func (c IndexConfig) DataType() types.DataType       { return c.dataType }
func (c IndexConfig) Metric() types.SimilarityMetric { return c.metric }
func (c IndexConfig) Dimension() int                 { return c.dimension }

// validate config
func (c IndexConfig) Validate() error {
	if c.indexType == 0 {
		return errors.New("index type is required")
	}
	if c.modelType == 0 {
		return errors.New("model type is required")
	}
	if c.dataType == 0 {
		return errors.New("data type is required")
	}
	if c.metric == 0 {
		return errors.New("similarity metric is required")
	}
	if c.dimension <= 0 {
		return errors.New("dimenison must be a positive integer")
	}
	return nil
}

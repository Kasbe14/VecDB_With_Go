package embedder

import (
	"VectorDatabase/internal/vector"
	"context"
)

// ID uniquely identifies a vector produced by an Embedder
// Ownership : Embedder
type VectorID string

// modality represents type of raw input data
type Modality string

const (
	ModalityText  Modality = "text"
	ModalityImage Modality = "image"
	// future
	// ModalityAudio Modality = "audio"
	// ModaliyVideo Modality = "video"
)

// SimilarityMetrics define how vector are compared
// Note declare early for clarity , the usage will come later
type SimilarityMetrics string

const (
	MetricCosine    SimilarityMetrics = "cosine"
	MetricDot       SimilarityMetrics = "dot"
	MetricEuclidean SimilarityMetrics = "euclidean"
)

// Embedder defines the contract  for converting  raw data into vectors
type Embedder interface {
	//converts raw data into vectors
	//guaratees -> Stable dimension and Unique ID for vector
	Embed(ctx context.Context, input any) (*vector.Vector, error)

	// Dimension returns fixed vector dimension produced by this Embedder
	Dimension() int

	//Modality returns supported input data type (modality : text, image, audio, video)
	Modality() Modality

	//Metric returns the similarity metric this metric was trained for..(use same metric for similarity score)
	Metric() SimilarityMetrics

	//Name returns a stable identifeir for thsi Embedder (eg - "openai-text-embedding-3-large")
	Name() string
}

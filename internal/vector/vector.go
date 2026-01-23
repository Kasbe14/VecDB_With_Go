package vector

import "errors"

type Vector struct {
	id         string
	values     []float32
	dimensions int
}

// consturctor for immutable vector
func NewVector(id string, vecValues []float32) (*Vector, error) {
	if len(vecValues) == 0 {
		return nil, errors.New("a vector must have atleast one dimensions")
	}
	//validate vector
	if err := validateValues(vecValues); err != nil {
		return nil, err
	}
	//normalize vector
	normalVec, err := Normalize(vecValues)
	if err != nil {
		return nil, err
	}
	//  copying for imutability
	// copied := make([]float32, len(vecValues))
	// copy(copied, vecValues)
	vec := &Vector{
		id:         id,
		values:     normalVec,
		dimensions: len(normalVec),
	}
	return vec, nil
}

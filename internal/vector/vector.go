package vector

import "errors"

type Vector struct {
	id         string
	values     []float32
	dimensions int
}

// consturctor for immutable vector
func New(id string, values []float32) (*Vector, error) {
	if len(values) == 0 {
		return nil, errors.New("VECTOR WITH NO DIMENSION : " +
			"At least one dimension needed")
	}

	//  copying for imutability
	copied := make([]float32, len(values))
	copy(copied, values)
	vec := &Vector{
		id:         id,
		values:     copied,
		dimensions: len(copied),
	}
	return vec, nil
}

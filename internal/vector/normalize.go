package vector

import "errors"

func Normalize(vec []float32) ([]float32, error) {
	if len(vec) == 0 {
		return nil, errors.New("vector with lenght zero")
	}
	magVec := Magnitude(vec)
	if magVec < epsilon {
		return nil, errors.New("zero magnitude vector")
	}
	normalized := make([]float32, len(vec))
	//copy(normalized, vec)
	for i, v := range vec {
		normalized[i] = v / float32(magVec)
	}
	return normalized, nil
}

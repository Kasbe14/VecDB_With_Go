package index

// SearchResult represent onematch result after vector serach, immutable, ordered by desceding similarity score
type SearchResult struct {
	vecId string
	score float64
}

func (r SearchResult) ID() string {
	return r.vecId
}
func (r SearchResult) Score() float64 {
	return r.score
}

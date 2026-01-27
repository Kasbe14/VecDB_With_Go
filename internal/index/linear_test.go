package index

import (
	"VectorDatabase/internal/vector"
	"testing"
)

func TestLinearIndex_AddAndGet(t *testing.T) {
	idx := NewLinearIndex()
	v, err := vector.NewVector(
		"mcV1bc",
		[]float32{1, 2, 3},
	)
	if err != nil {
		t.Fatalf("vector creation failed: %v", err)
	}
	if err := idx.Add(v); err != nil {
		t.Fatalf("failed to add vector %v", err)
	}
	got, ok := idx.Get("mcV1bc")
	if !ok {
		t.Fatalf("vector not got after add")
	}
	if got.ID() != v.ID() {
		t.Fatalf("expexted id %s, got %s", v.ID(), got.ID())
	}
	if got.Dimensions() != v.Dimensions() {
		t.Fatalf("dimensions mismatch")
	}

}

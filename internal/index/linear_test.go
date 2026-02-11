package index

import (
	"VectorDatabase/internal/types"
	v "VectorDatabase/internal/vector"
	"testing"
)

func TestNewLinearIndex_Constructor(t *testing.T) {
	// Sub-test 1: The Happy Path
	t.Run("ValidConfig", func(t *testing.T) {
		cfg, err := NewIndexConfig(types.LinearIndex, types.Testmodel, types.Text, types.Cosine, 128)
		if err != nil {
			t.Fatalf("Failed to create config: %v", err)
		}

		li, err := NewLinearIndex(cfg)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Check Invariants
		if li.vectors == nil {
			t.Fatal("Vector map was not initialized")
		}

		// Check Getters (Contracts)
		if li.Dimension() != cfg.Dimension() {
			t.Errorf("Dimension mismatch: got %d, want %d", li.Dimension(), cfg.Dimension())
		}

		// use Errorf instead of Fatalf here so the test continues
		// to check the other fields even if one fails.
		if li.config.Metric() != cfg.Metric() {
			t.Errorf("Metric mismatch: got %v, want %v", li.config.Metric(), cfg.Metric())
		}
	})

	// Sub-test 2: The Sad Path
	t.Run("InvalidConfig", func(t *testing.T) {
		invalidCfg := IndexConfig{} // Zero value
		lIdx, err := NewLinearIndex(invalidCfg)

		if err == nil {
			t.Error("Expected error for empty config, but got nil")
		}
		if lIdx != nil {
			t.Error("Expected nil index instance on failure")
		}
	})
}

// Helper to create a valid index for testing
func setupIndex(t *testing.T, dim int) *LinearIndex {
	cfg, _ := NewIndexConfig(types.LinearIndex, types.Testmodel, types.Text, types.Cosine, dim)
	idx, err := NewLinearIndex(cfg)
	if err != nil {
		t.Fatalf("failed to setup index: %v", err)
	}
	return idx
}

// Contract: Add must reject invalid IDs and dimension mismatches.
// Invariant: After Add, the vector must be retrievable via Get.
func TestLinearIndex_AddAndGet(t *testing.T) {
	idx := setupIndex(t, 3)
	vec, _ := v.NewVector([]float32{1.0, 0.0, 0.0}, 3)

	t.Run("Successful Add", func(t *testing.T) {
		exists, err := idx.Add("vec-1", vec)
		if err != nil || exists {
			t.Errorf("Expected success, got exists=%v, err=%v", exists, err)
		}

		// Verify via Get (RLock path)
		retrieved, ok := idx.Get("vec-1")
		if !ok || retrieved != vec {
			t.Error("Vector was not stored correctly")
		}
	})

	t.Run("Add Duplicate ID", func(t *testing.T) {
		exists, err := idx.Add("vec-1", vec)
		if err != nil || !exists {
			t.Error("Expected exists=true for duplicate ID")
		}
	})

	t.Run("Contract Violation: Dimension Mismatch", func(t *testing.T) {
		badVec, _ := v.NewVector([]float32{1.0, 0.0}, 2) // Dim 2 instead of 3
		_, err := idx.Add("bad-vec", badVec)
		if err == nil || err.Error() != "dimension mismatch" {
			t.Errorf("Expected dimension mismatch error, got %v", err)
		}
	})
}

// Contract: Delete must remove the item or return an error if missing.
// Post-condition: Get must return false after a successful Delete.
func TestLinearIndex_Delete(t *testing.T) {
	idx := setupIndex(t, 3)
	vec, _ := v.NewVector([]float32{1.0, 0.0, 0.0}, 3)
	idx.Add("vec-1", vec)

	t.Run("Successful Delete", func(t *testing.T) {
		err := idx.Delete("vec-1")
		if err != nil {
			t.Errorf("Delete failed: %v", err)
		}

		_, ok := idx.Get("vec-1")
		if ok {
			t.Error("Vector still exists after deletion")
		}
	})

	t.Run("Delete Non-existent", func(t *testing.T) {
		err := idx.Delete("ghost")
		if err == nil {
			t.Error("Expected error when deleting non-existent ID")
		}
	})
}

// Concurrency Test: Ensures no race conditions occur when multiple goroutines
// read and write at the same time.
func TestLinearIndex_Concurrency(t *testing.T) {
	idx := setupIndex(t, 3)
	vec, _ := v.NewVector([]float32{1.0, 0.0, 0.0}, 3)

	// Use a wait group to coordinate goroutines
	done := make(chan bool)

	// Start a writer
	go func() {
		for i := 0; i < 100; i++ {
			idx.Add("concur-vec", vec)
		}
		done <- true
	}()

	// Start a reader
	go func() {
		for i := 0; i < 100; i++ {
			idx.Get("concur-vec")
		}
		done <- true
	}()

	// Wait for both to finish
	for i := 0; i < 2; i++ {
		<-done
	}
	// If this test finishes without a panic, the Mutexes are working!
}

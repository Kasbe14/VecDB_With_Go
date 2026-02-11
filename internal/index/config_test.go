package index

import (
	"VectorDatabase/internal/types"
	"testing"
)

// Contract: NewIndexConfig must reject invalid inputs (Pre-conditions).
// Contract: Dimension must be strictly positive (> 0).
// Contract: IndexType, DataType, Metric, and ModelType must be valid enum values.
// Post-condition: If inputs are valid, a valid IndexConfig object is returned with no error.
func TestNewIndexConfig_Contracts(t *testing.T) {
	tests := []struct {
		name        string
		indexType   types.IndexType
		modelType   types.ModelType
		dataType    types.DataType
		metric      types.SimilarityMetric
		dimension   int
		expectError bool
		errorMsg    string // substring to check
	}{
		{
			name:        "Success: Valid Configuration",
			indexType:   types.LinearIndex,
			modelType:   types.Testmodel,
			dataType:    types.Text,
			metric:      types.Cosine,
			dimension:   128,
			expectError: false,
		},
		{
			name:        "Contract Violation: Invalid Dimension (Zero)",
			indexType:   types.LinearIndex,
			modelType:   types.Testmodel,
			dataType:    types.Text,
			metric:      types.Cosine,
			dimension:   0,
			expectError: true,
			errorMsg:    "invalid dimension",
		},
		{
			name:        "Contract Violation: Invalid Dimension (Negative)",
			indexType:   types.LinearIndex,
			modelType:   types.Testmodel,
			dataType:    types.Text,
			metric:      types.Cosine,
			dimension:   -5,
			expectError: true,
			errorMsg:    "invalid dimension",
		},
		{
			name:        "Contract Violation: Invalid IndexType",
			indexType:   types.IndexType(-1), // Casting a bad value
			modelType:   types.Testmodel,
			dataType:    types.Text,
			metric:      types.Cosine,
			dimension:   128,
			expectError: true,
			errorMsg:    "invalid index type",
		},
		{
			name:        "Contract Violation: Invalid DataType",
			indexType:   types.LinearIndex,
			modelType:   types.Testmodel,
			dataType:    types.DataType(999),
			metric:      types.Cosine,
			dimension:   128,
			expectError: true,
			errorMsg:    "invalid data type",
		},
		{
			name:        "Contract Violation: Invalid Metric",
			indexType:   types.LinearIndex,
			modelType:   types.Testmodel,
			dataType:    types.Text,
			metric:      types.SimilarityMetric(2020),
			dimension:   128,
			expectError: true,
			errorMsg:    "invalid metric type",
		},
		{
			name:        "Contract Violation: Invalid ModelType",
			indexType:   types.LinearIndex,
			modelType:   types.ModelType(-1),
			dataType:    types.Text,
			metric:      types.Cosine,
			dimension:   128,
			expectError: true,
			errorMsg:    "invalid model type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := NewIndexConfig(tt.indexType, tt.modelType, tt.dataType, tt.metric, tt.dimension)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error containing '%s', but got nil", tt.errorMsg)
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
				}
				// Ensure we returned an empty/zero value config on error
				if cfg.Dimension() != 0 {
					t.Error("Expected zero-value config on error, but got valid dimension")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

// Invariant: The IndexConfig object is Immutable.
// Invariant: Getters must return the exact values provided during construction.
// Note: Since fields are unexported (private), the only way to verify state is via getters.
func TestIndexConfig_InvariantsAndGetters(t *testing.T) {
	// Arrange
	expectedDim := 256
	expectedIdx := types.HNSWIndex
	expectedModel := types.Testmodel
	expectedData := types.Image
	expectedMetric := types.Euclidean

	// Act
	cfg, err := NewIndexConfig(expectedIdx, expectedModel, expectedData, expectedMetric, expectedDim)
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	// Assert: Verify Invariants via Getters
	if cfg.Dimension() != expectedDim {
		t.Errorf("Invariant broken: Expected dimension %d, got %d", expectedDim, cfg.Dimension())
	}
	if cfg.IndexType() != expectedIdx {
		t.Errorf("Invariant broken: Expected index type %v, got %v", expectedIdx, cfg.IndexType())
	}
	if cfg.ModelType() != expectedModel {
		t.Errorf("Invariant broken: Expected model type %v, got %v", expectedModel, cfg.ModelType())
	}
	if cfg.DataType() != expectedData {
		t.Errorf("Invariant broken: Expected data type %v, got %v", expectedData, cfg.DataType())
	}
	if cfg.Metric() != expectedMetric {
		t.Errorf("Invariant broken: Expected metric %v, got %v", expectedMetric, cfg.Metric())
	}
}

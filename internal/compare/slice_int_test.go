package compare

import (
	"testing"
)

func Test_validateSliceIntComparison(t *testing.T) {
	tests := []struct {
		name            string
		desiredSliceInt []interface{}
		actualSliceInt  []interface{}
		expectError     bool
		expectEqual     bool
	}{
		{
			name:            "ensure known good comparison passes",
			desiredSliceInt: []interface{}{1, 2, 3, 4},
			actualSliceInt:  []interface{}{4, 3, 2, 1},
			expectError:     false,
			expectEqual:     true,
		},
		{
			name:            "ensure known bad comparison fails with equal length",
			desiredSliceInt: []interface{}{1, 2, 3, 4},
			actualSliceInt:  []interface{}{5, 3, 2, 1},
			expectError:     false,
			expectEqual:     false,
		},
		{
			name:            "ensure known bad comparison fails with inequal length",
			desiredSliceInt: []interface{}{1, 2, 3, 4},
			actualSliceInt:  []interface{}{5, 4, 3, 2, 1},
			expectError:     false,
			expectEqual:     false,
		},
		{
			name:            "ensure out of range comparison fails",
			desiredSliceInt: []interface{}{1, 2, 3, 4},
			actualSliceInt:  []interface{}{},
			expectError:     false,
			expectEqual:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equal, err := EqualSliceInt(tt.desiredSliceInt, tt.actualSliceInt)
			hasError := err != nil
			if hasError != tt.expectError {
				t.Errorf("EqualSliceInt(%s, %s); hasError %s; expectError %v",
					tt.desiredSliceInt, tt.actualSliceInt, err, tt.expectError)
			}

			if equal != tt.expectEqual {
				t.Errorf("EqualSliceInt(%s, %s); equal %v; expectEqual %v",
					tt.desiredSliceInt, tt.actualSliceInt, equal, tt.expectEqual)
			}
		})
	}
}

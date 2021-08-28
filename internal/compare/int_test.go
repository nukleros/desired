package compare

import (
	"testing"
)

func Test_validateIntComparison(t *testing.T) {
	tests := []struct {
		name        string
		desiredInt  interface{}
		actualInt   interface{}
		expectError bool
		expectEqual bool
	}{
		{
			name:        "ensure known good comparison passes",
			desiredInt:  1,
			actualInt:   1,
			expectError: false,
			expectEqual: true,
		},
		{
			name:        "ensure known bad comparison fails",
			desiredInt:  1,
			actualInt:   2,
			expectError: false,
			expectEqual: false,
		},
		{
			name:        "ensure out of range comparison fails",
			desiredInt:  1,
			actualInt:   -1,
			expectError: false,
			expectEqual: false,
		},
		{
			name:        "ensure nil comparison fails",
			desiredInt:  1,
			actualInt:   nil,
			expectError: true,
			expectEqual: false,
		},
		{
			name:        "ensure desired type mismatch fails",
			desiredInt:  "string",
			actualInt:   1,
			expectError: true,
			expectEqual: false,
		},
		{
			name:        "ensure actual type mismatch fails",
			desiredInt:  1,
			actualInt:   "string",
			expectError: true,
			expectEqual: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equal, err := EqualInt(tt.desiredInt, tt.actualInt)
			hasError := err != nil
			if hasError != tt.expectError {
				t.Errorf("EqualInt(%s, %s); hasError %s; expectError %v",
					tt.desiredInt, tt.actualInt, err, tt.expectError)
			}

			if equal != tt.expectEqual {
				t.Errorf("EqualInt(%s, %s); equal %v; expectEqual %v",
					tt.desiredInt, tt.actualInt, equal, tt.expectEqual)
			}
		})
	}
}

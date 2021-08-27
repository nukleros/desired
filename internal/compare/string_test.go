package compare

import (
	"testing"
)

func Test_validateStringComparison(t *testing.T) {
	tests := []struct {
		name          string
		desiredString interface{}
		actualString  interface{}
		expectError   bool
		expectEqual   bool
	}{
		{
			name:          "ensure known good comparison passes",
			desiredString: "this is a test",
			actualString:  "this is a test",
			expectError:   false,
			expectEqual:   true,
		},
		{
			name:          "ensure known bad comparison fails",
			desiredString: "this is a test",
			actualString:  "this is another test",
			expectError:   false,
			expectEqual:   false,
		},
		{
			name:          "ensure out of range comparison fails",
			desiredString: "this is a test",
			actualString:  "⌥Z",
			expectError:   false,
			expectEqual:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equal, err := EqualString(tt.desiredString, tt.actualString)
			hasError := err != nil
			if hasError != tt.expectError {
				t.Errorf("EqualString(%s, %s); hasError %s; expectError %v",
					tt.desiredString, tt.actualString, err, tt.expectError)
			}

			if equal != tt.expectEqual {
				t.Errorf("EqualString(%s, %s); equal %v; expectEqual %v",
					tt.desiredString, tt.actualString, equal, tt.expectEqual)
			}
		})
	}
}

package compare

import (
	"testing"
)

func Test_validateSliceStringComparison(t *testing.T) {
	tests := []struct {
		name               string
		desiredSliceString []interface{}
		actualSliceString  []interface{}
		expectError        bool
		expectEqual        bool
	}{
		{
			name:               "ensure known good comparison passes",
			desiredSliceString: []interface{}{"one", "two", "three", "four"},
			actualSliceString:  []interface{}{"four", "three", "two", "one"},
			expectError:        false,
			expectEqual:        true,
		},
		{
			name:               "ensure known bad comparison fails with equal length",
			desiredSliceString: []interface{}{"one", "two", "three", "four"},
			actualSliceString:  []interface{}{"five", "three", "two", "one"},
			expectError:        false,
			expectEqual:        false,
		},
		{
			name:               "ensure known bad comparison fails with inequal length",
			desiredSliceString: []interface{}{"one", "two", "three", "four"},
			actualSliceString:  []interface{}{"five", "four", "three", "two", "one"},
			expectError:        false,
			expectEqual:        false,
		},
		{
			name:               "ensure out of range comparison fails",
			desiredSliceString: []interface{}{"one", "two", "three", "four"},
			actualSliceString:  []interface{}{},
			expectError:        false,
			expectEqual:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equal, err := EqualSliceString(tt.desiredSliceString, tt.actualSliceString)
			hasError := err != nil
			if hasError != tt.expectError {
				t.Errorf("EqualSliceString(%s, %s); hasError %s; expectError %v",
					tt.desiredSliceString, tt.actualSliceString, err, tt.expectError)
			}

			if equal != tt.expectEqual {
				t.Errorf("EqualSliceString(%s, %s); equal %v; expectEqual %v",
					tt.desiredSliceString, tt.actualSliceString, equal, tt.expectEqual)
			}
		})
	}
}

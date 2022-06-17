package compare

import (
	"testing"
)

func Test_validateTypeCompare(t *testing.T) {
	tests := []struct {
		name        string
		desiredType interface{}
		actualType  interface{}
		expectEqual bool
	}{
		{
			name:        "ensure known good comparison passes",
			desiredType: 1,
			actualType:  1,
			expectEqual: true,
		},
		{
			name:        "ensure known bad comparison fails",
			desiredType: 1,
			actualType:  "whoops",
			expectEqual: false,
		},
		{
			name:        "ensure out of range comparison fails",
			desiredType: 1,
			actualType:  nil,
			expectEqual: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equal := equalTypes(tt.desiredType, tt.actualType)
			if equal != tt.expectEqual {
				t.Errorf("EqualType(%s, %s); equal %v; expectEqual %v",
					tt.desiredType, tt.actualType, equal, tt.expectEqual)
			}
		})
	}
}

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

func TestSorter_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}

	tests := []struct {
		name   string
		sorter Sorter
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sorter.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Sorter.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

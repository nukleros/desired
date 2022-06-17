package desired

import (
	"testing"

	testdata "github.com/nukleros/desired/test"
)

func Test_equalDeepMapComparison(t *testing.T) {
	tests := []struct {
		name        string
		desiredData map[string]interface{}
		actualData  map[string]interface{}
		expectError bool
		expectEqual bool
	}{
		{
			name:        "ensure literal objects are equal",
			desiredData: testdata.EqualDeploymentDesired(),
			actualData:  testdata.EqualDeploymentDesired(),
			expectError: false,
			expectEqual: true,
		},
		{
			name:        "ensure actual can have fields unmanaged by desired",
			desiredData: testdata.EqualDeploymentDesired(),
			actualData:  testdata.EqualDeploymentActual(),
			expectError: false,
			expectEqual: true,
		},
		{
			name:        "ensure inequal objects return as not equal",
			desiredData: testdata.EqualDeploymentDesired(),
			actualData:  testdata.InequalDeployment(),
			expectError: false,
			expectEqual: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equal, err := Desired(tt.desiredData, tt.actualData)
			hasError := err != nil
			if hasError != tt.expectError {
				t.Errorf("Desired(%s, %s); hasError %s; expectError %v",
					tt.desiredData, tt.actualData, err, tt.expectError)
			}

			if equal != tt.expectEqual {
				t.Errorf("Desired(%s, %s); equal %v; expectEqual %v",
					tt.desiredData, tt.actualData, equal, tt.expectEqual)
			}
		})
	}
}

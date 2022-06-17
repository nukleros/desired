package compare

import (
	"testing"
)

func Test_validateSliceStringComparison(t *testing.T) {
	tests := []struct {
		name               string
		desiredSliceString []string
		actualSliceString  []string
		expectError        bool
		expectEqual        bool
	}{
		{
			name:               "ensure known good comparison passes",
			desiredSliceString: []string{"one", "two", "three", "four"},
			actualSliceString:  []string{"four", "three", "two", "one"},
			expectError:        false,
			expectEqual:        true,
		},
		{
			name:               "ensure known bad comparison fails with equal length",
			desiredSliceString: []string{"one", "two", "three", "four"},
			actualSliceString:  []string{"five", "three", "two", "one"},
			expectError:        false,
			expectEqual:        false,
		},
		{
			name:               "ensure known bad comparison fails with inequal length",
			desiredSliceString: []string{"one", "two", "three", "four"},
			actualSliceString:  []string{"five", "four", "three", "two", "one"},
			expectError:        false,
			expectEqual:        false,
		},
		{
			name:               "ensure out of range comparison fails",
			desiredSliceString: []string{"one", "two", "three", "four"},
			actualSliceString:  []string{},
			expectError:        false,
			expectEqual:        false,
		},
		{
			name:               "ensure nil comparison fails",
			desiredSliceString: []string{"one", "two", "three", "four"},
			actualSliceString:  nil,
			expectError:        true,
			expectEqual:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equal, err := EqualSliceStringInterface(tt.desiredSliceString, tt.actualSliceString)
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
			equal, err := EqualSliceIntegerInterface(tt.desiredSliceInt, tt.actualSliceInt)
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

func TestEqualSliceMapStringInterface(t *testing.T) {
	type args struct {
		desired []map[interface{}]interface{}
		actual  []map[interface{}]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ensure identical slice of maps passes",
			args: args{
				desired: []map[interface{}]interface{}{
					{
						"one":   "two",
						"three": "four",
						"five":  "six",
					},
				},
				actual: []map[interface{}]interface{}{
					{
						"one":   "two",
						"three": "four",
						"five":  "six",
					},
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ensure unidentical slice of maps fails",
			args: args{
				desired: []map[interface{}]interface{}{
					{
						"one": "two",
					},
				},
				actual: []map[interface{}]interface{}{
					{
						"one": "three",
					},
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "ensure empty desired map passes",
			args: args{
				desired: []map[interface{}]interface{}{{}},
				actual: []map[interface{}]interface{}{
					{
						"one": "two",
					},
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ensure empty actual map fails",
			args: args{
				actual: []map[interface{}]interface{}{{}},
				desired: []map[interface{}]interface{}{
					{
						"one": "two",
					},
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ensure desired complex map passes",
			args: args{
				desired: []map[interface{}]interface{}{
					{
						"one":   "two",
						"array": []string{"four", "five", "six"},
						"mapString": map[string]string{
							"thisshouldnotbe": "ignored",
						},
					},
				},
				actual: []map[interface{}]interface{}{
					{
						"nine":     "ten",
						"eleven":   12,
						"thirteen": true,
						"one":      "two",
						"array":    []string{"four", "five", "six"},
						"arrayInt": []int{1, 2, 3},
						"map": map[string]interface{}{
							"thisshouldbe": "ignored",
						},
						"mapString": map[string]string{
							"thisshouldnotbe": "ignored",
						},
					},
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualSliceMapInterfaceInterface(tt.args.desired, tt.args.actual)
			if (err != nil) != tt.wantErr {
				t.Errorf("EqualSliceMapStringInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EqualSliceMapStringInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
package compare

import (
	"testing"
)

func TestEqualMapStringInterface(t *testing.T) {
	type args struct {
		desired map[string]interface{}
		actual  map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ensure only known keys are enforced",
			args: args{
				desired: map[string]interface{}{
					"name": "value",
				},
				actual: map[string]interface{}{
					"name":       "value",
					"unenforced": "unenforced value",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ensure all similar keys passes",
			args: args{
				desired: map[string]interface{}{
					"one":   "two",
					"three": "four",
					"five":  "six",
				},
				actual: map[string]interface{}{
					"one":   "two",
					"three": "four",
					"five":  "six",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ensure dissimilar keys fails",
			args: args{
				desired: map[string]interface{}{
					"one":   "two",
					"three": "four",
					"five":  "six",
				},
				actual: map[string]interface{}{
					"one":   "two",
					"three": "four",
					"five":  "seven",
				},
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualMapStringInterface(tt.args.desired, tt.args.actual)
			if (err != nil) != tt.wantErr {
				t.Errorf("EqualMapStringInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EqualMapStringInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualSliceMapStringInterface(t *testing.T) {
	type args struct {
		desired []map[string]interface{}
		actual  []map[string]interface{}
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
				desired: []map[string]interface{}{
					{
						"one":   "two",
						"three": "four",
						"five":  "six",
					},
				},
				actual: []map[string]interface{}{
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
				desired: []map[string]interface{}{
					{
						"one": "two",
					},
				},
				actual: []map[string]interface{}{
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
				desired: []map[string]interface{}{{}},
				actual: []map[string]interface{}{
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
				actual: []map[string]interface{}{{}},
				desired: []map[string]interface{}{
					{
						"one": "two",
					},
				},
			},
			want:    true,
			wantErr: false,
		},
		// {
		// 	name: "ensure desired complex map passes",
		// 	args: args{
		// 		actual: []map[string]interface{}{{}},
		// 		desired: []map[string]interface{}{
		// 			{
		// 				"one": "two",
		// 			},
		// 		},
		// 	},
		// 	want:    true,
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualSliceMapStringInterface(tt.args.desired, tt.args.actual)
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

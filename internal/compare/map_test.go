package compare

import (
	"testing"
)

func TestEqualMap(t *testing.T) {
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
			got, err := EqualMap(tt.args.desired, tt.args.actual)
			if (err != nil) != tt.wantErr {
				t.Errorf("EqualMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EqualMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

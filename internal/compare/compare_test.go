package compare

import (
	"testing"
)

func Test_Compare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		desiredType interface{}
		actualType  interface{}
		want        bool
		wantErr     bool
	}{
		{
			name:        "ensure known good comparison passes",
			desiredType: 1,
			actualType:  1,
			want:        true,
			wantErr:     false,
		},
		{
			name:        "ensure known bad comparison fails",
			desiredType: 1,
			actualType:  "whoops",
			want:        false,
			wantErr:     true,
		},
		{
			name:        "ensure out of range comparison fails",
			desiredType: 1,
			actualType:  nil,
			want:        false,
			wantErr:     false,
		},
		{
			name:        "ensure desired with nil comparison fails",
			desiredType: nil,
			actualType:  "value",
			want:        false,
			wantErr:     false,
		},
		{
			name:        "ensure both desired and actual with nil passes",
			desiredType: nil,
			actualType:  nil,
			want:        true,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Compare(tt.desiredType, tt.actualType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNil(t *testing.T) {
	t.Parallel()

	type args struct {
		value interface{}
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ensure nil value returns true",
			args: args{
				value: nil,
			},
			want: true,
		},
		{
			name: "ensure string value returns false",
			args: args{
				value: "<nil>",
			},
			want: false,
		},
		{
			name: "ensure integer value returns false",
			args: args{
				value: -1,
			},
			want: false,
		},
		{
			name: "ensure bool value returns false",
			args: args{
				value: false,
			},
			want: false,
		},
		{
			name: "ensure structured returns false",
			args: args{
				value: []string{"one"},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isNil(tt.args.value); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

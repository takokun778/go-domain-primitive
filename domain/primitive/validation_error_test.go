package primitive_test

import (
	"errors"
	"go-domain-primitive/domain/primitive"
	"testing"
)

func TestAsValidationError(t *testing.T) {
	t.Parallel()

	type args struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "NewValidationError()",
			args: args{
				err: primitive.NewValidationError("test", nil),
			},
			want: true,
		},
		{
			name: "errors.New()",
			args: args{
				err: errors.New("test"), //nolint:goerr113
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := primitive.AsValidationError(tt.args.err); got != tt.want {
				t.Errorf("AsValidationError() = %v, want %v", got, tt.want)
			}
		})
	}
}

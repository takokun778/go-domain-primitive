package sample

import (
	"go-domain-primitive/domain/primitive"
	"reflect"
	"testing"
)

func Test_name_New(t *testing.T) {
	t.Parallel()

	type fields struct {
		value string
	}

	type args struct {
		v string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    primitive.Plimitive[string]
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				value: "test",
			},
			args: args{
				v: "test",
			},
			want: name{
				value: "test",
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				value: "",
			},
			args: args{
				v: "",
			},
			want: name{
				value: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := name{
				value: tt.fields.value,
			}
			got, err := n.New(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("name.New() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_name_Reconstruct(t *testing.T) {
	t.Parallel()

	type fields struct {
		value string
	}

	type args struct {
		v string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   primitive.Plimitive[string]
	}{
		{
			name: "success",
			fields: fields{
				value: "test",
			},
			args: args{
				v: "test",
			},
			want: name{value: "test"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := name{
				value: tt.fields.value,
			}
			if got := n.Reconstruct(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name.Reconstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_name_Validate(t *testing.T) {
	t.Parallel()

	type fields struct {
		value string
	}

	type args struct {
		v string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				value: "test",
			},
			args: args{
				v: "test",
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				value: "",
			},
			args: args{
				v: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := name{
				value: tt.fields.value,
			}
			if err := n.Validate(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("name.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_name_Value(t *testing.T) {
	t.Parallel()

	type fields struct {
		value string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				value: "test",
			},
			want: "test",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := name{
				value: tt.fields.value,
			}
			if got := n.Value(); got != tt.want {
				t.Errorf("name.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_name_String(t *testing.T) {
	t.Parallel()

	type fields struct {
		value string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				value: "test",
			},
			want: "*****",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := name{
				value: tt.fields.value,
			}
			if got := n.String(); got != tt.want {
				t.Errorf("name.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

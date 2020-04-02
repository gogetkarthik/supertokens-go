package supertokens

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		options []Options
	}
	tests := []struct {
		name string
		args args
		want SuperTokensClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHost(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		{
			name: "Set host test",
			args: args{host: "testHost"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHost(tt.args.host); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithScheme(t *testing.T) {
	type args struct {
		scheme string
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithScheme(tt.args.scheme); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_superTokens_Print(t *testing.T) {
	type fields struct {
		host   string
		scheme string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := superTokens{
				host:   tt.fields.host,
				scheme: tt.fields.scheme,
			}
		})
	}
}

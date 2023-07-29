package person2

import (
	"reflect"
	"testing"
)

func TestEldest(t *testing.T) {
	type args struct {
		a []any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Employees",
			args: args{a: []any{Employee{age: 25}, Employee{age: 50}}},
			want: Employee{age: 50},
		},
		{
			name: "Customers",
			args: args{a: []any{Customer{age: 25}, Customer{age: 50}}},
			want: Customer{age: 50},
		},
		{
			name: "Employees & Customers",
			args: args{a: []any{Employee{age: 25}, Customer{age: 50}}},
			want: Customer{age: 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eldest(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eldest() = %v, want %v", got, tt.want)
			}
		})
	}
}

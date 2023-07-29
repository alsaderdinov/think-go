package person

import "testing"

func TestEldest(t *testing.T) {
	type args struct {
		a []Ager
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Employees",
			args: args{a: []Ager{&Employee{age: 25}, &Employee{50}}},
			want: 50,
		},
		{
			name: "Customers",
			args: args{a: []Ager{&Customer{age: 25}, &Customer{50}}},
			want: 50,
		},
		{
			name: "Employees and Customers",
			args: args{a: []Ager{&Employee{age: 25}, &Customer{50}}},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eldest(tt.args.a...); got != tt.want {
				t.Errorf("Eldest() = %v, want %v", got, tt.want)
			}
		})
	}
}

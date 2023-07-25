package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestSortInts(t *testing.T) {
	got := []int{10, 4, 7, 9, 11, 2, 1, 3}
	sort.Ints(got)
	want := []int{1, 2, 3, 4, 7, 9, 10, 11}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v\n", got, want)
	}
}

func TestSortStrings(t *testing.T) {
	tests := []struct {
		name string
		got  []string
		want []string
	}{
		{
			name: "Test #1",
			got:  []string{"e", "d", "c", "b", "a"},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "Test #2",
			got:  []string{"Nirvana", "Black Sabbath", "Metallica", "Misfits", "Slayer"},
			want: []string{"Black Sabbath", "Metallica", "Misfits", "Nirvana", "Slayer"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.got)
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("got %v, want %v\n", tt.got, tt.want)
			}
		})
	}
}

func sampleInts() []int {
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	return data
}

func sampleFloats() []float64 {
	var data []float64
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Float64())
	}

	return data
}

func BenchmarkSortInts(b *testing.B) {
	data := sampleInts()
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

func BenchmarkSortFloats(b *testing.B) {
	data := sampleFloats()
	for i := 0; i < b.N; i++ {
		sort.Float64s(data)
	}
}

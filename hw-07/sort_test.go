package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func isEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func TestSortInts(t *testing.T) {
	got := []int{10, 4, 7, 9, 11, 2, 1, 3}
	sort.Ints(got)
	want := []int{1, 2, 3, 4, 7, 9, 10, 11}
	if !isEqual(got, want) {
		t.Errorf("получили %d, ожидалось %d\n", got, want)
	}
}

func TestSortStrings(t *testing.T) {
	tests := []struct {
		name string
		got  []string
		want []string
	}{
		{
			name: "Тест№1",
			got:  []string{"e", "d", "c", "b", "a"},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "Тест№2",
			got:  []string{"Nirvana", "Black Sabbath", "Metallica", "Misfits", "Slayer"},
			want: []string{"Black Sabbath", "Metallica", "Misfits", "Nirvana", "Slayer"},
		},
	}
	for _, tt := range tests {
		sort.Strings(tt.got)
		if !isEqual(tt.got, tt.want) {
			t.Errorf("получили %s, ожидалось %s\n", tt.got, tt.want)
		}
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

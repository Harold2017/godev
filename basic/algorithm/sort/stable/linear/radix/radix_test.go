package radix

import (
	sort2 "godev/basic/algorithm/sort"
	"godev/utils"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	a, b := make(sort2.IntSlice, 100), make(sort2.IntSlice, 100)
	v := 0
	for i := range a {
		v = utils.GenerateRandomIntInRange(0, 100)
		a[i], b[i] = v, v
	}
	sort.Stable(a)
	Sort(b)

	if !sort2.Equal(a, b) {
		t.Fatal(a, b)
	}
}

// BenchmarkSort-8   	 3000000	       481 ns/op
func BenchmarkSort(b *testing.B) {
	data := make(sort2.IntSlice, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = utils.GenerateRandomIntInRange(0, b.N)
	}

	b.ResetTimer()

	Sort(data)
}

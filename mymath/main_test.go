package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{1, 2, 3, 4, 5}, 3},
		test{[]int{0, 1, 2, 3, 4}, 2},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

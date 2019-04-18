package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	type test struct {
		data   int
		answer int
	}

	tests := []test{
		test{1, 7},
		test{10, 70},
		test{2, 14},
		test{3, 21},
	}

	for _, v := range tests {
		x := Years(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, ", got", x)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(10)
	if x != 70 {
		t.Error("Expected 70, got", x)
	}
}

func ExampleYears() {
	x := Years(10)
	fmt.Println(x)
	// Output:
	// 70
}

func ExampleYearsTwo() {
	x := YearsTwo(10)
	fmt.Println(x)
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

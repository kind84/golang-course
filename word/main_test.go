package word

import (
	"fmt"
	"sort"
	"testing"

	"github.com/kind84/golang-course/quote"
)

func TestCount(t *testing.T) {
	s := "The quick brown fox jumps over the lazy dog"
	x := Count(s)
	if x != 9 {
		t.Error("Expected 9, got", x)
	}
}

func TestUseCount(t *testing.T) {
	s := "hello world hello world hello"
	m := UseCount(s)
	for k, v := range m {
		switch k {
		case "hello":
			if v != 3 {
				t.Error("Word", k, ", expected 3, got", v)
			}
		case "world":
			if v != 2 {
				t.Error("Word", k, ", expected 2, got", v)
			}
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	s := "The quick brown fox jumps over the lazy dog"
	x := Count(s)
	fmt.Println(x)
	// Output:
	// 9
}

func ExampleUseCount() {
	s := "hello world hello world hello"
	m := UseCount(s)
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, m[v])
	}
	// Output:
	// hello 3
	// world 2
}

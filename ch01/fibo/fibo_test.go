package fibo

import (
	"testing"
)

var pairs = []struct {
	n     int
	value int
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{10, 55},
	{20, 6765},
	{30, 832040},
	{40, 102334155},
}

func TestFib(t *testing.T) {
	for _, pair := range pairs {
		if got := Fib(pair.n); got != pair.value {
			t.Errorf("Fib(%d) expect %d,but got%d", pair.n, pair.value, got)
		}
	}
}
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(40)
	}

}

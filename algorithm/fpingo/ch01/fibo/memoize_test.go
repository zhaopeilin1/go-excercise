package fibo

import (
	"testing"
)

func TestFibMemo(t *testing.T) {
	for _, pair := range pairs {
		if got := FibMemo(pair.n); got != pair.value {
			t.Errorf("Fib(%d) expect %d,but got%d", pair.n, pair.value, got)
		}
	}
}
func BenchmarkFibMemo(b *testing.B) {
	//	fn := FibMemoized
	for i := 0; i < b.N; i++ {
		FibMemo(40)
	}
}

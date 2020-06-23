package sum

import (
	"testing"
)

func benchmarkSumLoop(s []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumLoop(s)
	}
	return
}

func BenchmarkSumLoop40(b *testing.B) {
	benchmarkSumLoop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, b)

}

func benchmarkSumRecursive(s []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumRecursive(s)
	}
}

func BenchmarkSumRecursive40(b *testing.B) {
	benchmarkSumRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, b)
}

func benchmarkSumTailcall(s []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumTailCall(s)
	}
}

func BenchmarkSumTailcall40(b *testing.B) {
	benchmarkSumTailcall([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, b)
}

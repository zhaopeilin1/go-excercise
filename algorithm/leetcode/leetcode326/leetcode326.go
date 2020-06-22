package leetcode326

func isPowerOfThree(n int) bool {
	m:=1
	for m<n{
		m=m*3
	}
	return (m==n)
}

package qsort

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	//取得中间值
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			//如果左边的某个值大于mid。就把这个值移到
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}

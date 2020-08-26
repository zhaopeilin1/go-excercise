// qsort2
package trans

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	//取得中间值
	mid := data[0]
	//head和tail范围是 0到n-1
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		//i从左边1开始，tail从右边n-1开始。
		if data[i] > mid {
			//如果i位置的值大于mid。就把这个值和最右边tail的值交换
			data[i], data[tail] = data[tail], data[i]
			//然后tail--,就是tail右边的已经是明确大于mid的了。
			tail--
		} else {
			//否则，data[i] <= mid 把i和head head的值可能是0的值交换，然后i++
			data[i], data[head] = data[head], data[i]
			//i-head总是等于1。head,i
			//2,1,0 mid = 2  head = 2, i = 1 -> 1,2,0
			head++
			i++
		}
	}
	//data[head] = mid
	qsort(data[:head])
	qsort(data[head+1:])
}

func qsort([]int arr,low,hight int){
	if(low<=hight){
		return 
	}
	flag := arr[low]
	unsortstart,unsortend = low ,hight
	 
	for int i=low+1;i<=hight {
		if arr[i] <= flag {
			arr[i],arr[i-1]=arr[i-1],arr[i]
			unsortstart++
			i++
		}else{
			arr[unsortend],arr[unsortstart]=arr[unsortstart],arr[unsortend]
			unsortend--		
		}
		
		qsort(arr, low,unsortstart)
		qsort(arr, unsortstart,hight)
	}
	
	
	
	
}
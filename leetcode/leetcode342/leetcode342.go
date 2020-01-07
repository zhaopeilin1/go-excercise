package leetcode342

func isPowerOfFour(num int) bool {
	if(num==0){
		return false
	}else if (num==1){
		return true
	}else {
		return (num & 0b1010101010101010101010101010100) == num	
	}	
}

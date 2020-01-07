package leetcode342

func isPowerOfFour(num int) bool {
	if(num==0){
		return false
	}else if (num==1){
		return true
	}else {
		//4的幂
		return (num & 0b1010101010101010101010101010101) == num	
	}	
}

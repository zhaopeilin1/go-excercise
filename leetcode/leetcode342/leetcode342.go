package leetcode342

func isPowerOfFour(num int) bool {
	if(num==0){
		return false
	}else {
		for i:=1;i<=0b1000000000000000000000000000000;i=i<<2{
			if (num==i){
				return true
			}
			if(i>num){
				return false
			}
		}
		return false
		//4的幂
		//return (num & 0b1010101010101010101010101010101) == num	
	}	
}

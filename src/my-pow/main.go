package main
/*
	{
		from:"https://leetcode-cn.com/problems/my-pow",
		reference:[],
		description:"Pow(x,n)",
		solved:true,
	}
*/

/*
	1.直接乘n次 out of time
	2.递归pow（x，n/2）*pow（x，（n+1）/2））
	3.分奇偶讨论，减少一次递归
*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		n = -n
		x = 1 / x
	}
	if n%2 == 1 {
		var xx = pow(x, n/2)
		return xx * xx * x
	} else {
		var xx = pow(x, n/2)
		return xx * xx
	}
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n == 2 {
		return x * x
	}
	if n%2 == 1 {
		var xx = pow(x, n/2)
		return xx * xx * x
	} else {
		var xx = pow(x, n/2)
		return xx * xx
	}
}

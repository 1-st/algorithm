package _go

/*
	{
		from:"https://leetcode-cn.com/problems/multiply-strings",
		reference:[],
		description:"字符串相乘",
		solved:true,
	}
*/

/*
	1.string => uint8 array =>string
*/

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return "0"
	}
	if len(num1) == 1 && num1[0] == '0' {
		return "0"
	}
	if len(num2) == 1 && num2[0] == '0' {
		return "0"
	}

	var n1 = []uint8(num1)
	var n2 = []uint8(num2)

	var ans = make([]uint8, len(n1)+len(n2))

	for i := 0; i < len(num1); i++ {
		n1[i] -= '0'
	}
	for i := 0; i < len(num2); i++ {
		n2[i] -= '0'
	}

	for i1 := len(n1) - 1; i1 >= 0; i1-- {
		for i2 := len(n2) - 1; i2 >= 0; i2-- {
			level := len(n1) - 1 - i1 + len(n2) - 1 - i2
			add := n1[i1] * n2[i2]
			addList(&ans, uint8(len(ans)-1-level), add%10)
			addList(&ans, uint8(len(ans)-1-level-1), add/10)
		}
	}

	var i = 0
	for ; i < len(ans); i++ {
		if ans[i] != 0 {
			break
		}
	}
	for j := i; j < len(ans); j++ {
		ans[j] += '0'
	}
	return string(ans[i:])
}

func addList(ans *[]uint8, i, n uint8) {
	(*ans)[i] += n
	for (*ans)[i] >= 10 {
		n = (*ans)[i] / 10
		(*ans)[i] %= 10
		i--
		(*ans)[i]+=n
	}
}

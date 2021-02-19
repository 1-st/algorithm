package main

/*
{
	from:"https://leetcode-cn.com/problems/valid-number/",
	reference:[],
	description:"",
	solved:true,
	language:"Go",
}
*/
//     +/- 21 . 313 e +/- 1313
// 0	1  6  2  3  4  7    5

func isNumber(s string) bool {
	graph := []map[string]int{
		{"blank": 0, "sign": 1, "digit": 6},
		{"digit": 6, "dot": 2},
		{"digit": 3},
		{"digit": 3, "e": 4},
		{"digit": 5, "sign": 7},
		{"digit": 5},
		{"digit": 6, "dot": 3, "e": 4},
		{"digit": 5}}
	p := 0
	var getType = func(b byte)string{
		if b>='0'&&b<=9{
			return "digit"
		}else if b==' '{
			return "blank"
		}else if b=='.'{
			return "dot"
		}else if b=='e'{
			return "e"
		}else if b=='+'||b=='-'{
			return "sign"
		}else{
			return ""
		}
	}
	for ; p < len(s); {
		
	}
}

func main() {
}

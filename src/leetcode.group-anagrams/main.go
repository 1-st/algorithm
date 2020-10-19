package main

import (
	"fmt"
	"sort"
)

/*
{
	from:"https://leetcode-cn.com/problems/group-anagrams",
	reference:[],
	description:"字母异位分组",
	solved:true,
}
*/

/*
 */

type byteSlice []byte

func (p byteSlice) Len() int           { return len(p) }
func (p byteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p byteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func groupAnagrams(strs []string) [][]string {
	var dict = make(map[string]int)
	var ans = make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		var list = byteSlice(strs[i])
		sort.Sort(list)
		if v, ok := dict[string(list)]; !ok {
			dict[string(list)] = len(ans)
			ans = append(ans, append([]string{},strs[i]))
		} else {
			ans[v] = append(ans[v], strs[i])
		}
	}
	return ans
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "ate", "tae", "ant", "nat", "tan", "but"}))
}

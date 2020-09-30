package main

import "sort"

type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func getPermutation(n int, k int) string {
	k--
	var ans = make(ByteSlice, n)
	for i := 0; i < len(ans); i++ {
		ans[i] = byte(i + 1)
	}
	var getN func(i int) int
	getN = func(i int) int {
		if i == 0 {
			return 1
		}
		if i == 1 {
			return 1
		}
		return i * getN(i-1)
	}
	var N int
	var i = 0
	for i+1 < n {
		N = getN(n - (i + 1))
		j := k / N
		jj := k % N
		k = jj
		if j != 0 {
			ans[i], ans[i+j] = ans[i+j], ans[i]
		}
		sort.Sort(ans[i+1:])
		i++
	}
	for i := 0; i < len(ans); i++ {
		ans[i] += '0'
	}
	return string(ans)
}

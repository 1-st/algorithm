package main

import "bytes"

/*
{
	from:"https://leetcode-cn.com/problems/zigzag-conversion",
	reference:[],
	description:"Z字形变换",
	solved:false,
}
*/


func convert(s string, numRows int) string {
	strLen:= len(s)
	if numRows==1{
		return s
	}
	var buf bytes.Buffer
	for i:=0;i<strLen; {
		_ = buf.WriteByte(s[i])
		i+=2*numRows-2
	}
	for i:=1;i<numRows-1;i++{
		for j :=i;j<strLen;{
			_ = buf.WriteByte(s[j])
			j+=2*numRows-2*i-2
			if j<strLen{
				_ = buf.WriteByte(s[j])
				j+=2*i
			}
		}
	}
	for i:=numRows-1;i<strLen;{
		_ = buf.WriteByte(s[i])
		i+=2*numRows-2
	}
	return buf.String()
}
package main




func max(a int,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}


func min(a int,b int)int{
	if a>b{
		return b
	}else{
		return a
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	var a1,a2 *[]int
	if len1>len2{
		len1,len2 = len2,len1
		a1,a2 = &nums2,&nums1
	}else{
		a1,a2 = &nums1,&nums2
	}
	i1 := 0
	i2 := 0
	for i,j:=0,len1;i<=j;{
		i1 = (i+j)/2
		i2 = (len1+len2+1)/2-i1
		if i1<len1 && (*a1)[i1]<(*a2)[i2-1]{
			i = i1+1
		}else if i1>0 && (*a1)[i1-1]>(*a2)[i2]{
			j = i1-1
		}else{
			var maxLeft int
			var minRight int
			if i1==0{
				maxLeft = (*a2)[i2-1]
			}else if i2==0 {
				maxLeft = (*a1)[i1-1]
			}else{
				maxLeft = max((*a1)[i1-1],(*a2)[i2-1])
			}
			if (len1+len2)%2==1{
				return float64(maxLeft)
			}
			if i1==len1{
				minRight = (*a2)[i2]
			}else if i2==len2 {
				minRight = (*a1)[i1]
			}else{
				minRight = min((*a1)[i1],(*a2)[i2])
			}
			return float64(maxLeft+minRight)/2.0
		}
	}
	return 0.0
}


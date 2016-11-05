###4.1-1
  数组所有元素都是负数时候，FIND-MAXIMUM-SUBARRAY返回数组最小值。二分法分到组后一个数组各个元素。两个负数的和肯定小于原来任何一个负数。


###4.1-2
```
  FIND-MAXIMUM-SUBARRAY(A[]) {
	max = -∞
	left = 0
	right = 0
	for i = 1 to A.lenght()  
		sum = 0
		for j = i to A.length 
			sum = sum + A[j]
			if sum > max 
				left = i
				right = j
				max = sum

	return (left, right, max)			
}
```

###4.1-4
   修改现有的算法，允许返回空最大子数组。主要应对整个子数组都是负的环境。

```
FIND-MAXIMUM-SUBARRAY(A, low, high) 
	if high == low
		if A[low] < 0
			return (0, 0, 0)
		else 
			return (low, high, a[low])
	else 
		mid = (low + high) / 2
		(left-low, left-high, left-sum) = FIND-MAXIMUM-SUBARRAY(A, low, mid)
		(righ-low, righ-high, righ-sum) = FIND-MAXIMUM-SUBARRAY(A, mid+1, high)	
		(cross-low, cross-high, cross-sum) = FIND-MAX-CROSSING-SUBARRAY(A, low, mid, high)
   
	if left-sum >= right-sum and left-sum >= cross-sum 
		return (left-low, left-high, left-sum)
	else if right-sum >= left-sum and right-sum >= cross-sum
		return (right-low, right-high, right-sum)
    else 
		return (cross-low, cross-high, cross-sum)	
```

###4.1-5
   线性查找最大子数组的代码(chapter-4-code/Find_max_subarray_linear/main.go)

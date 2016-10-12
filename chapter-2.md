###2.1-1

###2.1-2

```
// The start index array is 0
INSERT_SORT(A[])  {
 for i=1; i<A.length(); i++ {
	key = A[i];  
    j=i-1;
    while j >= 0 and A[j] < key {
        A[j + 1] = A[j] 
		j = j -1
    }
   A[j+1] = key 
 }
}
```

### 2.1-3

伪代码
```
// Find the value in array，start index 1
FIND_VALUE_INDEX(A[], Value) {
  v = nil
  for i=1 to i<A.length() {
	if A[i] == Value  {
     v = i
     break
  }
  
  return v 
}
```

- 循环不变式: 在整个for循环中(循环变量i), v=i并且A[i] == value， 或者 v = nil

-证明过程: 
  --1) 初始化：在第一次循环前A[1, 0]为空数组，v=nil; 不需要证明.
  --2) 保持: 证明每次迭代都不变： 在整for循环中，如果找到A[i]等于Value， 记录v=i。
  --3) 终止：当i>A.length; 能够找到v！= nil代表找到值，v=nil能够找到值，循环式不变。


### 2.1-4

伪代码
```
Binary_Value_Inc(B1[], B2[]) {
	// 数组起始索引为1
	let R[1, B1.length+1] be new array， init as zero
    for i = 1 to B1.length {
		t = B1[i] + B2[i] + R[i]
        R[i+1] = t / 2
        R[i] = t % 2    // 求余数		
    }
}
```
- 循环不变式: 在整个for循环中(循环变量i) R[1,i+1] 代表 B1[1, i] 与 B2[1, i] 两个二进制的和
- 证明过程: 
  --1) 初始化：在第一次循环前(当i=1时)，B1[1, 0] 与 B2[i, 0] 为空数组，两个空数组和为0, R[1,1] 为0, 循环不变式成立
  --2) 保持: 证明每次迭代都不变： 在整for循环中，分别重B1， B2两个二进制低位到高位计算， t 是求几个位十进制值， 通过和二求余计算当前位二进制值
             R[i] = t % 2， 与2相除求进位的值R[i+1] = t / 2, 整个过程中R[1,i+1] 为B1[1, i] 与 B2[i, i] 数组二进制值数的和 
  --3) 终止：当i>A.length; R[1,i+1] 代表 B1[1, i] 与 B2[i, i] 两个二进制的和不变。


### 2.2-1

    f(n) = (n^3) / 1000 - 100 X (n ^2) - 100 X n + 3 

    O(n^3)


### 2.2-2

选择排序法的为代码:
```
SELECT_SORT(A[]) {
// 数组起始索引为1

  for i = 1 to A.length -1 {
	for j = i + 1  to A.length {
		if a[j] < a[i] {
			t = a[i]
            a[i] = a[j]
			a[j] = t
        }
     }
  }
} 
```

- 循环不变式: 在整个for循环中(循环变量i), 子数组A[1, i-1] 为有序数组`A[1]<A[2]...<A[i-1]`, A[1, i-1]为A数组前i-1个最小值 
- 对A前n-1个元素配排序:每次都选择A[1,i-1]的最小值， A[n]必定为数组最大值,放在末尾.
- 选择排序最好和最坏的情况O(n^2)


### 2.2-3

   - 假定搜索的元素为数组任意一个元素，平均搜索元素个数。假定每个元素查找概率是ri，查找次数是n

```
   平均查找次数 = (n*r1*1 + n*r2*2 + ......n*rn*n) = n(r1x1 + r2x2 + r3x3 ..... rn * n)

   如果出现概率是一样: 1 + 2 + 3... +n = (1+n)/2
```
   - 最坏的情况是查找n次

   - 平均和最坏量级是分别 O(n), O(n)

   
### 2.2-4

    修改算法，使之有更低增长量级。


### 2.3-1
   数组A=<3, 41, 52, 26, 38, 57, 9, 49>

- 合并1: 3, 41 | 26, 52 | 38, 57 | 9, 49
- 合并2: 3, 26, 41, 52  | 9, 38, 49, 57 
- 合并3: 3, 9 , 26, 38, 41, 49, 52, 57


### 2.3-2
 
不使用哨兵，一个数组合并完成，另外数组拷贝到A
```
MERGE(A, p, q, r) {
	n1 = q - p + 1
	n2 = r - q

    let L[1, n1] and R[1, n2] be new arrays

	//复制数据
	for i = 1 to n1 {
		L[i] = A[p + i -1]
	}

    for i = 1 to n2 {
		R[i] = A[q + i]
	}

	il = 1
    ir = 1
	for k = p to r {
		// the left array copy finishe
		if il > L.Length {
			A[k] = R[ir]
			ir = ir + 1
			continue
		} else if ir > R.Length {
			// the right array copy finishe
			A[k] = L[ir] 
			il = il + 1
			continue
		}

		if L[il] < R [ir] {
			A[k] = L[il]
			il = il + 1
		} else {
			A[k] = R[ir]
			ir = ir + 1
		}
	}
}
```

### 2.3-3 

证明:

	当n刚好是2的幂的时候递归式f(n) = nlgn. 当为n=2^i

- 当i=1 (2的一次幂)时候， T(n) = 2  = 2*lg2 成立
- 当i=i n=2^i的时候, i = lgn
    f(n) =  2^i + 2*2^(i-1) + 2^2*2^(i-2) .... 2^i * 2 = i * (2^i) = lgn * n 
- 当n=2^(i+1)  i+1 = lgn
	f(n) =  2^(i+1) + 2 * 2^(i+1-1) ....2^(i+1)*2 = nlgn


### 2.3-4 

插入排序法的递归式:
```
	     1           (n = 1) 
T(n) = { 
         T(n-1) + n  (n > 1)
```

### 2.3-5

二分查找法

```
BINARY_SEARCH(A[], value) {

l = 1;
r = A.length
index = nil 
while l < r {
	m = (r - 1) / 2
	if A[m] == value {
		index =  m
		break
	} if (value > m) {
		l = m + 1
	} else {
		r = m - 1
	}
}

return index
}
```

二分查找法的递归式:
```
         1           (n = 1)
T(n) = {
         T(n/2) + 1  (n > 1)

```

### 2.3-6

  插入排序使用二分法可以更快的查找到插入位置，但是最坏情况下，移动元素复杂度不变。还是O(n^2)

### 2.3-7
   可以.使用归并排序的时候，合并元素的时候如果有两个元素相等，而且为x时候就可以了。

思考题目：

### 2-1

- 在最坏情况下

  k(k+1)/2 * (n/k) = n(k+1)/2 < nk

- 因为每次递归需要合并元素个都为n， 递归次数为lg(n/k)次

- k的选择

  在最坏的情况下插入排序法是时间是k(k+1)/2 * (n/k) = n(k+1)/2
  最坏情况下归并排序: nlg(n/k)  

  
  平均情况是
  （n(k+1)/2 + 1 ） / 2 + nlg(n/k) < nlgn

  (k+1)/4 + 1/2 +  lgn - lgk < lgn
  `k+3/2 < 4lgk

### 2-2 

```
1 for i=1; to A.length-1 
2    for j=A.length downto i+1
3        if A[j] < A[j-1] 
4			exchange A[j] with A[j-1]
```

- 2~4行代码循环不变式，并且证明:
  -- 2~4 代码循环不式:  在整个for循环中(循环变量j) A[j] 为 A[j..A.length]数组的最小的元素. 
  --1) 初始化：在第一次循环前(当j=A.length时), A[j] 为 A[A.length..A.length]数组最小元素，因为数组只有一个元素，所以成立.
  --2) 保持: 证明每次迭代都不变： 在整for循环中, 如果A[j] < A[j-1]就交换，交换后A[j-1]为最小元素，因为执行以后j=j-1 A[j]为A[j..A.length]最小元素成立。
  --3) 终止：当结束以后, j = i ， A[i] = A[j] 为A[j..A.length]数组的最小的元素。

- 1~4 代码循环不变式： 在整个for循环中(循环变量i) A[1..i-1]是从小到大排序。
  --1) 初始化：在第一次循环前(当i=1时), A[1..0] 为空数组，所以是已经排序.
  --2) 保持: 第一行for循环结束后产生A[i]都是A[i..A.length]的最小的元素，所以有A[1] < A[2]... A[i], 循环不变式不变.
  --3) 终止: i = length-1, A[i] = A[length-1, length], 执行以后A[length]为数组最大元素，循环不变式不变。




### 2-4

查找逆序对

- 列出A数组<2,3,8,6,1>逆序对: [2,1] [3,1] [8,6] [8,1] [6,1]	
- 由集合{1,2,3,4,5,...n}组成的什么集合有最多逆序对<n,n-1,n-2.....3,2,1>最多逆序对。有(1+(1-n))n/2=(n^2)/2 对逆序对
- 逆序对越多，插入排序法就越慢。例如当插入数组第n个元素时候，n与前面n-1个元素都是逆序对，插入必须一一移动前面各个元素。
- [在最坏情况nlgn可以搜索数组逆序对](chapter-2-code/Find-Inversion/main.go)

  

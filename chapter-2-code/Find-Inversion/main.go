// Find-Inversion project main.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func ComputeInversionCountTwoZone(a []int, p, q, r int) int {
	result := 0
	n1 := q - p + 1
	n2 := r - q

	al := make([]int, n1)
	ar := make([]int, n2)

	fmt.Printf("p=%d q=%d, r=%d\n", p, q, r)
	// init the left array
	for i := 0; i < n1; i++ {
		al[i] = a[p+i]
	}
	fmt.Println(al)

	// init the right array
	for i := 0; i < n2; i++ {
		ar[i] = a[q+1+i]
	}
	fmt.Println(ar)

	il := 0
	ir := 0
	// merge two array
	for i := p; i <= r; i++ {

		if il >= len(al) {
			a[i] = ar[ir]
			ir++
			continue
		} else if ir >= len(ar) {
			a[i] = al[il]
			il++
			continue
		}

		// is inversion
		if al[il] <= ar[ir] {
			a[i] = al[il]
			il++
		} else {
			a[i] = ar[ir]
			ir++
			result += len(al) - il
		}
	}

	return result
}

func Binary_Find_Inversion(a []int, l, r int, count *int) {
	if l < r {
		m := (r + l) / 2
		fmt.Printf("l=%d, r=%d\n", l, r)
		Binary_Find_Inversion(a, l, m, count)
		Binary_Find_Inversion(a, m+1, r, count)
		result := ComputeInversionCountTwoZone(a, l, m, r)
		fmt.Println(result)
		*count += result
	}
}

func main() {
	array_len := len(os.Args)
	a := make([]int, array_len-1)
	var err error
	var tmp int
	var count int

	for i := range a {
		tmp, err = strconv.Atoi(os.Args[i+1])
		if err == nil {
			a[i] = tmp
		}
	}

	count = 0
	Binary_Find_Inversion(a, 0, len(a)-1, &count)

	fmt.Printf("The count of inversion: %d\n", count)
	fmt.Println(a)

}

// Find_max_subarray_linear project main.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func Find_max_subarray(a []int) (left, right, max int) {

	if len(a) <= 0 {
		return -1, -1, -1
	}

	if len(a) == 1 {
		return 0, 0, a[0]
	}

	max = a[0]
	left = 0
	right = 0
	lost_seat_sum := 0

	// lost_seat_sum [max array ... lost_seat_sum ... current index]
	for i := 1; i < len(a); i++ {
		if (a[i] + lost_seat_sum + max) > max {
			max = (a[i] + lost_seat_sum + max)
			lost_seat_sum = 0
			if i != right+1 {
				sum := 0
				new_left := left
				for j := i; j >= left; j-- {
					sum = sum + a[j]
					if sum > max {
						new_left = j
						max = sum
					}
				}
				left = new_left
			}
			right = i
		} else if a[i] > max {
			left = i
			right = i
			max = a[i]
			lost_seat_sum = 0
		} else {
			lost_seat_sum += a[i]
		}

		fmt.Printf("left:%d right:%d max:%d, lost_seat_sum:%d \n", left, right, max, lost_seat_sum)
	}

	return left, right, max
}

func main() {

	array_len := len(os.Args)
	a := make([]int, array_len-1)
	var err error
	var tmp int

	for i := range a {
		tmp, err = strconv.Atoi(os.Args[i+1])
		if err == nil {
			a[i] = tmp
		}
	}

	left, right, max := Find_max_subarray(a)

	fmt.Printf("The max of sub array is [%d, %d] value is %d \n", left, right, max)
}

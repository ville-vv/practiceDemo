package main

import "fmt"

// 二分查找方法
func BinarySearch(a []int, b int) int {
	var (
		low     int = 0
		hight   int = len(a) - 1
		counter     = 0
	)

	for i := 0; i < len(a); i++ {
		middle := (low + hight) / 2
		counter++
		if a[middle] > b {
			hight = middle - 1
		} else if a[middle] < b {
			low = middle + 1
		} else {
			fmt.Println("次数：", counter)
			return middle
		}
		if hight < low {
			fmt.Println("次数：", counter)
			return -1
		}
	}
	fmt.Println("次数：", counter)
	return -1
}

func FastSort(ds []int, left, right int) {
	var (
		index = left
	)

	if len(ds) <= 1 {
		return
	}
	tempVal := ds[index]
	for i := left + 1; i < right; i++ {
		fmt.Printf(" temp=%d, i=%d, index=%d \t", tempVal, i, index)
		if tempVal > ds[i] {
			index++
			ds[index], ds[i] = ds[i], ds[index]
		}
		fmt.Println(ds)
	}
	ds[index], ds[left] = ds[left], ds[index]
	if left > index-1 {
		return
	}
	FastSort(ds, left, index)
	FastSort(ds, index+1, right)
}

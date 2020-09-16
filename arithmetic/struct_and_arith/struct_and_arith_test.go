package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"math/rand"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	set := mapset.NewSet()

	for i := 0; i < 24; i++ {
		set.Add(rand.Intn(10000000))
	}
	setarr := set.ToSlice()
	bucket := make([]int, len(setarr))
	for i := 0; i < len(setarr); i++ {
		bucket[i] = setarr[i].(int)
	}
	n := bucket[1]
	// bucket = []int{128162, 954425, 1902081, 2186258, 3024728, 4941318, 4965466, 5511528, 6122540, 6203300}
	sort.Ints(bucket)
	fmt.Println("数组：", bucket, n)
	fmt.Println("查找到的位置：", BinarySearch(bucket, 128162))

}

func TestFastSort(t *testing.T) {
	ds := []int{14, 13, 5, 7, 1, 2, 8, 20, 59}
	FastSort(ds, 0, len(ds))
	fmt.Println(ds)
}

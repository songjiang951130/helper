package base

import (
	"fmt"
	"sort"
)

func DisplayArraySort() {
	var arr []int = []int{34, 21, 6, 17, 100, 99, 1, 5, 8}
	fmt.Println("sort before:", arr)
	Sort(arr)
	fmt.Println("sort after :", arr)
}

func Sort(arr []int) {
	//默认升序
	sort.Ints(arr)
	fmt.Println("sort res:", arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println("sort desc res:", arr)
}

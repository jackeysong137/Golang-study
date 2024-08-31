package main

import (
	"fmt"
	"sort"
)

func main() {

	//排序
	var numSlice = []int{9, 8, 7, 6, 4, 5}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] > numSlice[j] {
				//交换 swp
				swap2(numSlice, i, j)
			}
		}
	}
	fmt.Println(numSlice)
	//排序可以直接有sort包
	var numSlice_1 = []int{9, 8, 7, 6, 4, 5}
	sort.Ints(numSlice_1)
	fmt.Println(numSlice_1)
	//降序
	sort.Sort(sort.Reverse(sort.IntSlice(numSlice_1)))
	fmt.Println(numSlice_1)

}

func swap(numSlice []int, i int, j int) {
	numSlice[i] = numSlice[i] ^ numSlice[j]
	numSlice[j] = numSlice[i] ^ numSlice[j]
	numSlice[i] = numSlice[i] ^ numSlice[j]
}

/*
使用元素赋值语法来实现交换
*/
func swap2(numSlice []int, i int, j int) {
	numSlice[i], numSlice[j] = numSlice[j], numSlice[i]
}

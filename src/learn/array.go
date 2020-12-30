package main

import "fmt"

//初始化数组
func InitArray() {
	var arr0 [5]int = [5]int{1, 2, 3}
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4, 5, 6}
	var str = [5]string{3: "hello world", 4: "tom"}

	fmt.Println(arr0, arr1, arr2, str)
}

package reverseInt

import (
	"strconv"
	"strings"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
func ReverseInt(x int) int {
	moreZero := x > 0
	var abs int
	if moreZero {
		abs = x
	} else {
		abs = x * -1
	}

	str := strconv.Itoa(abs)
	// fmt.Printf("%d\n", abs)
	// fmt.Println(str)
	arr := strings.Split(str, "")
	// fmt.Printf("%+v\n", arr)
	ListRverse(&arr)

	str = strings.Join(arr, "")
	result, _ := strconv.Atoi(str)

	if !moreZero {
		result *= -1
	}

	return result
}

//ListRverse 字符串数组反转
func ListRverse(list *[]string) {
	var temp string
	var length int = len(*list)
	for i := 0; i < length/2; i++ {
		temp = (*list)[i]
		(*list)[i] = (*list)[length-1-i]
		(*list)[length-1-i] = temp
	}
}

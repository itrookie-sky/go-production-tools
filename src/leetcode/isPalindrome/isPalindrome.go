package isPalindrome

import (
	"math"
)

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？

整型各个位提取数值公式
n/10^(site-1)%10
*/
func IsPalindrome(x int) bool {
	// fmt.Println("x ", x)

	temp := x
	length := 0
	for temp > 0 {
		temp /= 10
		length++
	}

	if length == 1 || x == 0 {
		return true
	}

	one := x % 10

	if one == 0 || x < 0 {
		return false
	}
	left, right := 0, 0
	leftTemp, rightTemp := x, x

	half := length / 2
	// fmt.Println("half ", half)
	for i := 0; i < half; i++ {
		curDigits := int(math.Pow(10, float64(i)))

		left += (leftTemp / int(math.Pow(10, float64(length-1-i)))) % 10 * curDigits
		// fmt.Printf("length-i %d\nleft %d\n", leftTemp/int(math.Pow(10, float64(length-1-i))), (leftTemp/int(math.Pow(10, float64(length-1-i))))%10)

		cur := rightTemp % 10
		// fmt.Printf("rightTemp %d\ncur %d\n", rightTemp, cur)

		rightTemp /= 10
		right += cur * curDigits
		// fmt.Printf("10^i %d\n", curDigits)
		// fmt.Printf("left %d \nright %d\n\n", left, right)
	}

	return left == right
}

//官方解答
func LeetcodeIsPalindrome(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

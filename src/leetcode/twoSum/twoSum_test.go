package twoSum

import "testing"

//单元测试
func TestTwoSum(t *testing.T) {
	if r := TwoSum([]int{1, 2, 3, 4, 5}, 6); !(r[0] == 1 && r[1] == 3) {
		t.Errorf("nums [1 2 3 4 5] target 6 result [1 3],but %+v go", r)
	}
}

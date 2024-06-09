package noknow

import "testing"

func TestSilce(t *testing.T) {
	var num1 []interface{}
	num2 := []int{1, 2, 3}
	num1 = append(num1, num2)
	t.Log(num1)
}

package slice

import (
	"fmt"
	"testing"
)

func Test_SliceV1(t *testing.T) {
	arr := []int{1, 2}
	fmt.Printf("%p\n", &arr)
	doSlice(&arr)
	fmt.Println(arr)

}
func doSlice(arr *[]int) {
	fmt.Println(arr)
	fmt.Printf("%p\n", *arr)
	*arr = append(*arr, 1, 3, 4, 5)
}

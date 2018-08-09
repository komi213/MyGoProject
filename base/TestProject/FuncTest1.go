package main

import "fmt"

//Go语言函数

func main() {
	var a int = 21
	var b int = 20
	// 调用函数并返回最大值
	maxValue := max(a,b)
	fmt.Printf("最大值为:%d\n",maxValue)

	x,y := swap("张","三")
	fmt.Printf(x,y)

	//返回函数
	var sequenceFunc = getSequence()
	fmt.Printf("%d\n",sequenceFunc())

	//杨辉三角
	PascalTriangle()

	nums1 := [] int{1,3,5,7,9}
	nums2 := append(nums1,3)
	fmt.Printf("%v",nums1)//直接输出数组nums1
	fmt.Printf("%v",nums2)//直接输出数组nums2


}

//1.基础
// 函数返回两个数的最大值
func max(a, b int) int {
	if a > b{
		return a
	}else{
		return b
	}
}
//2.Go语言函数可以返回多个值。
//交换位置
func swap(x, y string) (string, string) {
	return y,x;
}

//B、闭包
//Go语言支持匿名函数，可作为闭包。
//匿名函数是一个"内联"语句或表达式。
//匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
//创建函数getSequence()，函数体内返回另外一个函数。
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//行数
const LINES int = 10
// 杨辉三角
func PascalTriangle() {
	nums := []int{}
	for i := 0; i < LINES; i++ {
		//补空白
		for j := 0; j < (LINES - i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, " ")
		}
		fmt.Println("")
	}
}


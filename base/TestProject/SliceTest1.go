package main

import "fmt"

//Go语言切片

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main() {
	/*var slice1 [] int
	var slice2 [] int = make([]int, 6)
	slice1 = []int{1,2,3,4}
	slice2 = []int{1,2,3}
	printSlice(slice1)
	printSlice(slice2)

	s :=[] int {1,2,3 }
	fmt.Println(s)*/

	//4. nil切片
	/*var slice1 [] int
	printSlice(slice1)
	if slice1 == nil{
		fmt.Println("slice1 is nil")
	}*/

	//5.切片截取
	//可以通过设置下限及上限来设置截取切片[lower-bound:upper-bound]。
	//创建切片
	/*numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	//打印原始切片
	fmt.Println("numbers==",numbers);

	//打印子切片从索引1(包含) 到索引4(不包含)
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	//默认下限为0
	fmt.Println("numbers[:3] ==", numbers[:3])

	//默认上限为 len(s)
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)//len=0 cap=5 slice=[]

	//打印子切片从索引0(包含)到索引2(不包含)
	number2 := numbers[:2]
	printSlice(number2)//len=2 cap=9 slice=[0 1]

	//打印子切片从索引3(包含)到索引7(不包含)
	number3 := numbers[3:7]
	printSlice(number3)//len=4 cap=6 slice=[3 4 5 6]

	//6、切片的本质
	//切片在本质上是数组的视图。切片是一个数组片段的描述，
	//包含了指向数组的指针、片段的长度和容量（片段的最大长度）。切片与底层数组的关系如下：
	numbers[5] = 88
	printSlice(numbers)
	printSlice(number3)*/

	//sliceExtend()

	//7、切片追加和拷贝
	sliceOperate()
}


func sliceExtend(){
	//定义数组
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := arr[2:]
	s4 := arr[:6]
	s5 := arr[:]
	printSlice(s1)//len=4 cap=8 slice=[2 3 4 5]
	printSlice(s2)//len=2 cap=5 slice=[5 6]
	printSlice(s3)//len=8 cap=8 slice=[2 3 4 5 6 7]
	printSlice(s4)//len=6 cap=10 slice=[0 1 2 3 4 5]
	printSlice(s5)//len=10 cap=10 slice=[0 1 2 3 4 5 6 7]

	//对切片中数据进行修改，底层数组数据也被修改
	s1[0] = 100
	printSlice(s1)//len=4 cap=8 slice=[100 3 4 5]
	printSlice(s5)//len=10 cap=10 slice=[0 1 100 3 4 5 6 7]
	fmt.Println(arr)//[0 1 100 3 4 5 6 7]
	s1 = s1[2:6]
	printSlice(s1)//len=4 cap=6 slice=[4 5 6 7]
	//切片不能引用底层数组范围外的数据
	//printSlice(s1[2:9])//报错
}


func sliceOperate() {

	var numbers = [3]int{0,1}//数组
	fmt.Println(numbers)//[0 1 0]
	s1 := numbers[1:]//切片
	s1[0] = 101//修改切片的元素的值
	printSlice(s1)//len=2 cap=2 slice=[101 0]
	fmt.Println(numbers)//[0 101 0]

	//向切片追加一个元素(底层数组的容量会翻倍增长cap从2变为4)
	s1 = append(s1, 2)
	printSlice(s1)//len=3 cap=4 slice=[101 0 2]

	fmt.Println(numbers)//[0 101 0]

	//向切片添加一个元素，已经超越cap，系统会为切片分配更大的底层数组
	s1 = append(s1, 3)
	printSlice(s1)//len=4 cap=4 slice=[101 0 2 3]
	fmt.Println(numbers)//[0 101 0]

	//同时添加多个元素，系统为切片分配新的底层数组
	s1 = append(s1, 4,5,6)
	printSlice(s1)//len=7 cap=8 slice=[101 0 2 3 4 5 6]
	fmt.Println(numbers)//[0 101 0]

	//创建切片numbers1,len与numbers切片相同，cap为2倍
	s2 := make([]int, len(s1), (cap(s1))*2)

	//拷贝numbers的内容到numbers1
	copy(s2,s1)
	printSlice(s2)//len=7 cap=16 slice=[101 0 2 3 4 5 6]

	//将s2中的0删除
	s2 = append(s2[0:1], s2[2:]...)
	printSlice(s2)//len=6 cap=16 slice=[101 2 3 4 5 6]

	//删除第一个元素
	front := s2[0]
	fmt.Println(front)//101
	s3 := s2[1:]
	printSlice(s3)//len=5 cap=15 slice=[2 3 4 5 6]

	//删除最后一个元素
	tail := s2[len(s2)-1]
	s4 := s2[:len(s2) - 1]
	fmt.Println(tail)//6
	printSlice(s4)//len=5 cap=16 slice=[101 2 3 4 5]


	fmt.Println(numbers)//[0 101 0]
	fmt.Println(len(numbers),cap(numbers))//3 3

}

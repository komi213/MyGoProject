package main

import "fmt"
//range

func main() {
	rangeDemo()
	a := map[int]string{1:"hello",2:"world"}
	for k,v := range a{
		fmt.Print(k,v)
	}
}

func rangeDemo(){
	//使用range遍历切片
	nums := []int{2, 3, 4}
	sum := 0
	//返回的切片的索引可以省略
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//在切片上使用range,返回index和值
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	fmt.Println(sum)

	//使用range迭代map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//使用range遍历Unicode字符串，返回字符的索引与字符（Unicode的值）本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
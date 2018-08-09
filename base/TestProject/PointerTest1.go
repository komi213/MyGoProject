package main

import "fmt"

//指针

const MAX int = 3
func main() {
	/*
	//1.简单实例
	var a int= 20   // 声明实际变量
	var ip *int        //声明指针变量
	ip = &a  //指针变量的存储地址
	fmt.Printf("a 变量的地址是: %x\n", &a  )
	//指针变量的存储地址
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	//使用指针访问值
	fmt.Printf("*ip 变量的值: %d\n", *ip )

	//2.nil指针也称空指针。
	//nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
	//一个指针变量通常缩写为ptr。
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr  )*/


	//3.指针数组
	/*a := []int{10,100,200}
	var i int
	var ptr [MAX]*int;

	for  i = 0; i < MAX; i++ {
		ptr[i] = &a[i] //整数地址赋值给指针数组
	}

	for  i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
	}*/

	//4.二级指针
	//如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
	//当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：
	/*var a int
	var ptr *int
	var pptr **int
	var ppptr ***int
	a = 3000
	// 指针 ptr 地址
	ptr = &a
	// 指向指针 ptr 地址
	pptr = &ptr
	ppptr = &pptr

	// 获取 pptr 的值
	fmt.Printf("变量 a = %d, a的地址 a = %x\n", a, &a)
	fmt.Printf("指针变量 *ptr = %d, *ptr的地址 *ptr = %x\n", *ptr,ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d, **pptr的地址 **pptr = %x\n", **pptr, pptr)
	fmt.Printf("指向指针的指针变量 ***ppptr = %d, **ppptr的地址 **pptr = %x\n", ***ppptr, ppptr)
	*/

	//5.Go语言指针参数
	//Go语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。

	//定义局部变量
	var a int = 100
	var b int = 200
	fmt.Printf("交换前a的值:%d\n",a)
	fmt.Printf("交换前b的值:%d\n",b)

	//调用函数用于交换值 &a 指向 a 变量的地址  &b 指向 b 变量的地址
	swap2(&a,&b);
	fmt.Printf("交换后 a 的值 : %d\n", a )
	fmt.Printf("交换后 b 的值 : %d\n", b )

}

func swap2(x, y *int)  {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
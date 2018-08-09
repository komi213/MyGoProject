package main

import "fmt"

func main() {
	/*var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	fmt.Printf("a 的值为: %d\n", a)

	for a := 0; a < 10; a++ {
		fmt.Printf("a1 的值为: %d\n", a)
	}

	fmt.Printf("a<10 的值为: %d\n", a)

	for a < b {
		a++
		fmt.Printf("a2 的值为: %d\n", a)
	}

	fmt.Printf("a<b 的值为: %d\n", a)

	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}*/

	/* 定义局部变量 */
	/*var i, j int

	for i=2; i < 15; i++ {
		for j=2; j <= (i/j); j++ {
			if(i%j==0) {
				break; // 如果发现因子，则不是素数
			}
		}
		if(j > (i/j)) {
			fmt.Printf("%d  是素数\n", i);
		}
	}
	fmt.Printf("%d\n",i);
	fmt.Printf("%d",j);*/

	/* 定义局部变量 */
	/*var a int = 10

	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			//使用 break 语句跳出循环
			break;
		}
	}
	fmt.Printf("循环后 a 的值为 : %d\n", a)//a=16
	*/

	/* 定义局部变量 */
	/*var a int = 10
	for a < 20 {
		if a == 15 {
			// 跳过此次循环
			a++
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
	fmt.Printf("循环后 a 的值为 : %d\n", a)//a=20*/

	/* 定义局部变量 */
	/*var a int = 10

LOOP:
	for a < 20 {
		if a == 15 {
			//跳过迭代
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}*/

	for true{
		fmt.Printf("这是无限循环。\n");
	}

}

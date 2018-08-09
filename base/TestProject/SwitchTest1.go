package main

import (
	"fmt"
)

func grade(score int)string{
	var grade string;
	switch {
	case score<0 && score>100:
		fmt.Println("score is error")
		grade = "Error"
	case score<60:
		fmt.Println("F")
		grade = "F"
	case score < 80:
		fmt.Println("C")
		grade = "C"
	case score<90:
		fmt.Println("B")
		grade = "B"
	case score<=100:
		fmt.Println("A")
		grade = "A"
	}
	return grade
}

func operator(a,b int,op string)int{
	var result int
	switch op {
	case "+":
		result = a+b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func typeSwitch()  {
	var x interface{}
	switch i := x.(type){
	case nil:
		fmt.Printf("x的类型:%T",i);
	case int:
		fmt.Printf("x是int型")
	case float64:
		fmt.Printf("x是float64")
	case func(int)float64:
		fmt.Printf("x是funct(int)型")
	case bool,string:
		fmt.Printf("x是bool或string型")
	default:
		fmt.Printf("未知型")
	}
}


func main()  {
	//fmt.Println(grade(100))
	//fmt.Println(operator(3,6,"*"))
	typeSwitch();
}
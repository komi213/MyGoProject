package main

import "fmt"

//Map操作
func main()  {
	mapDemo()
	fmt.Println("-----分隔线----")
	deleteMap()
}

func mapDemo() {
	//map集合创建
	//var countryCapitalMap map[string]string
	countryCapitalMap := make(map[string]string)

	//map插入key-value对
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"
	//map值的修改
	countryCapitalMap["France"] = "巴黎"

	//输出国家首都信息
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	//查看元素在集合中是否存在
	capital, ok := countryCapitalMap["美国"]
	//如果确定是真实的,则存在,否则不存在
	if ok {
		fmt.Println("美国的首都是", capital)
	} else {
		fmt.Println("美国的首都不存在")
	}
}

func deleteMap(){
	//创建map
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	//打印地图
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}

	//删除元素
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	//打印地图
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}
}


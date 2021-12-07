package main

import (
	"fmt"
	"time"
)

/**
https://www.jianshu.com/p/79c029c0bd58
*/
func main() {

	// defer执行顺序
	// deferOrderTest()

	// 匿名返回值 命名返回值
	// deferReturnTest()
	// explain1OfDeferReturnTest()
	// explain2OfDeferReturnTest()

	// defer计算顺序
	// deferComputedOrderTest()

}

/*
https://www.cnblogs.com/zhangboyu/p/7753120.html
defer声明时会先计算确定参数的值，defer推迟执行的仅是其函数体
*/
func deferComputedOrderTest() {
	defer printTime(time.Now())
	time.Sleep(time.Second)
	fmt.Println("deferComputedOrderTest time:", time.Now()) // deferComputedOrderTest time: 2021-12-08 22:22:04.389826 +0800 CST m=+1.001281501
}

func printTime(t time.Time) {
	fmt.Println("defer input time:", t)          // defer input time: 2021-12-08 22:22:03.388731 +0800 CST m=+0.000171251
	fmt.Println("defer print time:", time.Now()) // defer print time: 2021-12-08 22:22:04.39119 +0800 CST m=+1.002645210
}

/*
https://www.cnblogs.com/zhangboyu/p/7753120.html
return 其实应该包含前后两个步骤：第一步是给返回值赋值（若为有名返回值则直接赋值，若为匿名返回值则先声明再赋值）
*/
func explain2OfDeferReturnTest() {
	intPtr := returnIntPtr()
	fmt.Println("returnIntPtr return:", *intPtr, intPtr) // returnIntPtr return: 2 0xc082008340
}

func returnIntPtr() *int {
	var i int

	defer func() {
		i++
		fmt.Println("returnIntPtr defer2:", i, &i) // returnIntPtr defer2: 2 0xc082008340
	}()
	defer func() {
		i++
		fmt.Println("returnIntPtr defer1:", i, &i) // returnIntPtr defer1: 1 0xc082008340
	}()
	return &i
}

/*
https://blog.csdn.net/zhghost/article/details/100738098
i的值在defer第一次走到的位置就被确认下来了
*/
func explain1OfDeferReturnTest() {
	i := 1
	// defer func() {
	// 	fmt.Println("defer i value:", i)	// defer i value: 2
	// }()

	defer fmt.Println("defer i value:", i) // defer i value: 1
	i++

	fmt.Println("main i value:", i) // main i value: 2
}

func deferReturnTest() {
	fmt.Printf("returnValues: %v\r", returnValues())
	// fmt.Printf("namedReturnValues: %v\r", namedReturnValues())
}

func returnValues() int {
	var result int
	defer func() {
		result++
	}()
	return result // 0
}

func namedReturnValues() (result int) {
	defer func() {
		result++
	}()
	return result // 1
}

func deferOrderTest() {
	fmt.Println("print 1")
	defer fmt.Println("defer 1")

	// return

	fmt.Println("print 2")
	defer fmt.Println("defer 2")
}

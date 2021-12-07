package main

import "golang-demo/scence"

func main() {

	// defer执行顺序
	// scence.DeferOrderTest()

	// 匿名返回值 命名返回值
	// scence.DeferReturnTest()
	// scence.DeferReturnTestExplain1()
	// scence.DeferReturnTestExplain2()

	// defer计算顺序
	scence.DeferComputedOrderTest()

}

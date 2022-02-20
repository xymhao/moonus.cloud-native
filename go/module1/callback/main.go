package main

import "fmt"

func main() {
	//匿名函数
	func() {
		fmt.Println("hello world")
	}()

	//某些特定场景执行，没有必要为全局定义这个函数，可以用闭包来做
	defer func() {
		//释放锁
		//回滚
		fmt.Println("recover in func")
	}()

	Do(1, add)
	Do(1, increase)

	a := 1
	changeParameter(a)
	fmt.Println(a) //1

	changeParameterPoint(&a)
	fmt.Println(a) //4
}

func add(a, b int) {
	fmt.Println(a + b)
}

func increase(a, b int) {
	fmt.Println(a - b)

}
func Do(value int, f func(a, b int)) {
	f(value, 1)
}

func changeParameter(a int) {
	a = 4
}

func changeParameterPoint(a *int) {
	//修改当前 *a 内存地址的值为4，类似于C#中的引用
	*a = 4
}

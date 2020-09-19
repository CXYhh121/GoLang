package main

import "fmt"

//赋值一个常量是，之后没有赋值的常量都会应用上一行的赋值表达式
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
	d = 5    // 5
	e
)

// 赋值两个常量，iota只会增长一次，而不会因为使用了两次就增长两次
const (
	Apple, Banana     = iota + 1, iota + 3 //Apple=1 Banana=3
	Cherimoya, Durian                      //Cherimoya=2 Durian=4
	Elderberry, Fig                        //Elderberry=3 Fig=5
)

//使用iota 结合位运算符 表示资源状态的使用案例
const (
	Open    = 1 << iota //0001 --1
	Close               //0010 --2
	Pending             //0100 --4
)

func main() {
	fmt.Println(a, b, c, d, e)

	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)

	fmt.Println(Open, Close, Pending)
}

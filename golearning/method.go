package main

import "fmt"

//type TwoInts struct {
//	a int
//	b int
//}
//
//func main() {
//	two1 := new(TwoInts)
//	two1.a = 12
//	two1.b = 10
//
//	fmt.Printf("The sum is: %d\n", two1.AddThem())
//	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))
//
//	two2 := TwoInts{3, 4}
//	fmt.Printf("The sum is: %d\n", two2.AddThem())
//}
//
//func (tn *TwoInts) AddThem() int {
//	return tn.a + tn.b
//}
//
//func (tn *TwoInts) AddToParam(param int) int {
//	return tn.a + tn.b + param
//}
//type Point struct {
//	x, y float64
//}
//
//func (p *Point) Abs() float64 {
//	return math.Sqrt(p.x*p.x + p.y*p.y)
//}
//
//type NamedPoint struct {
//	Point
//	name string
//}
//
//func main() {
//	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
//	fmt.Println(n.Abs()) // 打印5
//}

//内嵌类型的方法和继承
//type Engine interface {
//	Start()
//	Stop()
//}
//
//type Car struct {
//	Engine
//}
//
//func (c *Car) GoToWorkIn() {
//	//get in car
//	c.Start()
//	// drive to work
//	c.Stop()
//	//get out of car
//}
//
//func main() {
//	var c *Car
//	c.GoToWorkIn()
//}

// Log类型包含日志功能，通过聚合(组合)方式实现
type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	cus := new(Customer)
	cus.Name = "Lisa"
	cus.log = new(Log)
	cus.log.msg = "Lisa is my girlfriend!"

	// shorter
	cus = &Customer{"zhangweixin", &Log{"1 - this is my best friend!"}}
	//
	cus.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(cus.Log())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

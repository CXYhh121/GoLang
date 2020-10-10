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
//type Log struct {
//	msg string
//}
//
////type Customer struct {
////	Name string
////	log  *Log
////}
//type Customer struct {
//	Name string
//	Log
//}
//
//func main() {
//	//通过聚合方式实现
//	//cus := new(Customer)
//	//cus.Name = "Lisa"
//	//cus.log = new(Log)
//	//cus.log.msg = "Lisa is my girlfriend!"
//	//
//	//// shorter
//	//cus = &Customer{"zhangweixin", &Log{"1 - this is my best friend!"}}
//	////
//	//cus.Log().Add("2 - After me the world will be a better place!")
//	//fmt.Println(cus.Log())
//
//	//通过内嵌方式实现
//	c := &Customer{"lisa", Log{"1 - this is my best friend!"}}
//	c.Add("2 - After me the world will be a batter place!")
//	fmt.Println(c)
//}
//
//func (l *Log) Add(s string) {
//	l.msg += "\n" + s
//}
//
//func (l *Log) String() string {
//	return l.msg
//}
//
//func (c *Customer) String() string {
//	return c.Name + "\nLog: " + fmt.Sprintln(c.Log)
//}

//多重继承，Go可通过在类型中嵌入所有必要的父类型就可以简单的实现多重继承
//type Phone struct{}
//
//func (p *Phone) call()  string{
//	return "Ring Ring Ring"
//}
//
//type Camera struct{}
//
//func (c *Camera) takePhotos() string {
//	return "Click"
//}
//
//type CameraPhone struct {
//	Camera
//	Phone
//}
//
//func main() {
//	cp := new(CameraPhone)
//	fmt.Println("Our new CameraPhone exhibits multiple behaviors.....")
//	fmt.Println("It exhibits behavior of a Camera: ", cp.takePhotos())
//	fmt.Println("It exhibits behaviors is a Phone: ", cp.call())
//}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (sef Base) MoreMagic() {
	sef.Magic()
	sef.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}

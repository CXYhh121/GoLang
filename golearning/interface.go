package main

import "fmt"

//接口和反射
//类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
//实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
//一个类型可以实现多个接口。
//接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。

//type Shaper interface {
//	Area() float32
//}
//
//type Square struct {
//	side float32
//}
//
//func (sq *Square) Area() float32 {
//	return sq.side * sq.side
//}
//
//type Rectangle struct {
//	length, width float32
//}
//
//func (re Rectangle) Area() float32 {
//	return re.length * re.width
//}
//
//func main() {
//	//sq1 := new(Square)
//	//sq1.side = 5
//
//	//var areaIntF Shaper
//	//areaIntF = sq1
//	//fmt.Printf("The square has area: %f\n", areaIntF.Area())
//
//	r := Rectangle{5,3}  // Area() of Rectangle needs a value
//	q := &Square{5}      // Area() of Square needs a pointer
//	// shapes := []Shaper{Shaper(r), Shaper(q)}
//	// or shorter
//
//	shapes := []Shaper{r, q}
//	fmt.Println("Looping through shapes for area ....")
//	for n, _ := range shapes {
//		fmt.Println("shape details: ", shapes[n])
//		fmt.Println("Are of this shape is: ", shapes[n].Area())
//	}
//}

//valuable.go

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

//method to determine the value of a stock position
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

// method to determine the value of the car
func (c car) getValue() float32 {
	return c.price
}

//contract that defines different things that have value
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is: %f\n", asset.getValue())
}

func main() {
	var o valuable = stockPosition{"GOOG", 577.200, 4}
	showValue(o)
	o = car{"BWM", "M3", 657600}
	showValue(o)
}

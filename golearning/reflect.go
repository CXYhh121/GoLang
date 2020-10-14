package main

import "fmt"

//反射包，reflect
//func main() {
//	var x float64 = 3.43
//	fmt.Println("type:", reflect.TypeOf(x))
//	v := reflect.ValueOf(x)
//	fmt.Println("Value:", v)
//	fmt.Println("type", v.Type())
//	fmt.Println("kind", v.Kind())  //kind 返回最底层的类型
//	fmt.Println("value:", v.Float()) // 返回这个 float64 类型的实际值
//	fmt.Println(v.Interface())
//	fmt.Printf("value is %5.2e\n", v.Interface())
//	y := v.Interface().(float64)
//	fmt.Println(y)
//}

//通过反射修改(设置)值, 有些内容需要通过地址才能改变值或状态
//func main() {
//	var x float64 = 3.43
//	fmt.Println("type:", reflect.TypeOf(x))
//	v := reflect.ValueOf(x)
//	// setting a value:
//	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
//	fmt.Println("settability of v:", v.CanSet())
//	v = reflect.ValueOf(&x) // Note: take the address of x.
//	fmt.Println("type of v:", v.Type())
//	fmt.Println("settability of v:", v.CanSet())
//	v = v.Elem()
//	fmt.Println("The Elem of v is: ", v)
//	fmt.Println("settability of v:", v.CanSet())
//	v.SetFloat(3.1415) // this works!
//	fmt.Println(v.Interface())
//	fmt.Println(v)
//}

//reflect_struct.go

//type NotKnownType struct {
//	s1, s2, s3 string
//}
//
//func (n NotKnownType) String() string {
//	return n.s1 + " - " + n.s2 + " - " + n.s3
//}
//
//// variable to investigate:
//var secret interface{} = NotKnownType{"Ada", "Go", "Oberon"}
//
//func main() {
//	value := reflect.ValueOf(secret) // <main.NotKnownType Value>
//	typ := reflect.TypeOf(secret)    // main.NotKnownType
//	// alternative:
//	//typ := value.Type()  // main.NotKnownType
//	fmt.Println(typ)
//	knd := value.Kind() // struct
//	fmt.Println(knd)
//
//	// iterate through the fields of the struct:
//	for i := 0; i < value.NumField(); i++ {
//		fmt.Printf("Field %d: %v\n", i, value.Field(i))
//		// error: panic: reflect.Value.SetString using value obtained using unexported field
//		//value.Field(i).SetString("C#")
//	}
//
//	// call the first method, which is String():
//	results := value.Method(0).Call(nil)
//	fmt.Println(results) // [Ada - Go - Oberon]
//}

//print 函数中所用到的反射
//type Stringer interface {
//	String() string
//}
//
//type Celsius float64
//
//func (c Celsius) String() string {
//	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
//}
//
//type Day int
//var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
//
//func (day Day) String() string {
//	return dayName[day]
//}
//
//func print(args ...interface{}) {
//	for i, arg := range args {
//		if i > 0 {os.Stdout.WriteString(" ")}
//		switch a := arg.(type) {
//		case Stringer:
//			os.Stdout.WriteString(a.String())
//          fallthrough
//		case int:
//			os.Stdout.WriteString(strconv.Itoa(a))
//		case string:
//			os.Stdout.WriteString(a)
//		default:
//			os.Stdout.WriteString("????")
//		}
//	}
//}
//
//func main()  {
//	print(Day(1), "was", Celsius(18.39))
//}

//接口与动态类型
//type IDuck interface {
//	Quack()
//	Walk()
//}
//
//func DuckDance(duck IDuck) {
//	for i := 1; i <= 3; i++ {
//		duck.Quack()
//		duck.Walk()
//	}
//}
//
//type Bird struct {
//	// ...
//}
//
//func (b *Bird) Quack() {
//	fmt.Println("I am quacking!")
//}
//
//func (b *Bird) Walk()  {
//	fmt.Println("I am walking!")
//}
//func main() {
//	b := new(Bird)
//	DuckDance(b)
//}

type Any interface{}
type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
	// ...
}
type Cars []*Car

func main() {
	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	// query:
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)
	//
	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}

// Process all cars with the given function f:
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// Find all cars matching a given criteria.
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)

	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

// Process cars and create new data.
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	// Prepare maps of sorted cars.
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	// Prepare appender function:
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}

package main

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

//type stockPosition struct {
//	ticker     string
//	sharePrice float32
//	count      float32
//}
//
////method to determine the value of a stock position
//func (s stockPosition) getValue() float32 {
//	return s.sharePrice * s.count
//}
//
//type car struct {
//	make  string
//	model string
//	price float32
//}
//
//// method to determine the value of the car
//func (c car) getValue() float32 {
//	return c.price
//}
//
////contract that defines different things that have value
//type valuable interface {
//	getValue() float32
//}
//
//func showValue(asset valuable) {
//	fmt.Printf("Value of the asset is: %f\n", asset.getValue())
//}
//
//func main() {
//	var o valuable = stockPosition{"GOOG", 577.200, 4}
//	showValue(o)
//	o = car{"BWM", "M3", 657600}
//	showValue(o)
//}

//接口嵌套接口
//比如接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 Close() 方法。
//type ReadWrite interface {
//	Read(b Buffer) bool
//	Write(b Buffer) bool
//}
//
//type Lock interface {
//	Lock()
//	Unlock()
//}
//
//type File interface {
//	ReadWrite
//	Lock
//	Close()
//}

//类型断言：如何检测和转换接口变量的类型
//type_interface.go

//type Square struct {
//	side float32
//}
//
//type Circle struct {
//	radius float32
//}
//
//type Shaper interface {
//	Area() float32
//}
//
//func main() {
//	var areaIntF Shaper
//	sq1 := new(Square)
//	sq1.side = 21
//
//	areaIntF = sq1
//	// is Square the type of areaIntF
//	if t, ok := areaIntF.(*Square); ok {
//		fmt.Printf("the type of areaIntF is: %T\n", t)
//	}
//	if u, ok := areaIntF.(*Circle); ok {
//		fmt.Printf("the tyoe of areaIntfis: %T\n", u)
//	} else {
//		fmt.Println("areaIntF does not contain a variable of type Circle")
//	}
//
//	switch t := areaIntF.(type) {
//	case *Square:
//		fmt.Printf("Type Square %T with value %v\n", t, t)
//	case *Circle:
//		fmt.Printf("Type Circle %T with value %v\n", t, t)
//	case nil:
//		fmt.Printf("nil value: nothing to check?\n")
//	default:
//		fmt.Printf("Unexpected type %T\n", t)
//	}
//}
//
//func (sq *Square) Area() float32 {
//	return sq.side * sq.side
//}
//
//func (ci * Circle) Area() float32 {
//	return ci.radius * ci.radius * math.Pi
//}

//实现冒泡排序
//type Sorter interface {
//	Len() int
//	Less(i, j int) bool
//	Swap(i, j int)
//}
//
//func Sort(data Sorter) {
//	for pass := 1; pass < data.Len(); pass++ {
//		for i := 0; i < data.Len() - pass; i++ {
//			if data.Less(i + 1, i) {
//				data.Swap(i, i + 1)
//			}
//		}
//	}
//}
//
//func IsSorted(data Sorter) bool {
//	n := data.Len()
//	for i := n - 1; i > 0; i-- {
//		if data.Less(i, i-1) {
//			return false
//		}
//	}
//	return true
//}
//
//// Convenience type for common cases
//type IntArray []int
//
//func (p IntArray) Len() int           { return len(p) }
//func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
//func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
//
//
//type StringArray []string
//
//func (p StringArray) Len() int           { return len(p) }
//func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }
//func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
//
//// Convenience wrappers for common cases
//func SortIntS(a []int)       { Sort(IntArray(a)) }
//func SortStrings(a []string) { Sort(StringArray(a)) }
//
//func IntSAreSorted(a []int) bool       { return IsSorted(IntArray(a)) }
//func StringsAreSorted(a []string) bool { return IsSorted(StringArray(a)) }
//
////sortedMain.go
//func intS() {
//	data := []int{74, 59, 243, 542, -13, 9, 76, 234, 89, -234}
//	a := IntArray(data)
//	Sort(a)
//	if !IsSorted(a) {
//		panic("fails")
//	}
//	fmt.Printf("The sorted array is: %v\n", a)
//}
//
//func stringS() {
//	data := []string{"monday", "friday", "wednesday", "sunday", "thursday", "", "saturday"}
//	a := StringArray(data)
//	Sort(a)
//	if !IsSorted(a) {
//		panic("fails")
//	}
//	fmt.Printf("The sorted array is: %v\n", a)
//}
//
//type day struct {
//	num int
//	shortName string
//	longName string
//}
//
//type dayArray struct {
//	data []*day
//}
//
//func (d *dayArray) Len() int { return len(d.data) }
//func (d *dayArray) Less(i, j int) bool { return d.data[i].num < d.data[j].num}
//func (d *dayArray) Swap(i, j int) { d.data[i], d.data[j] = d.data[j], d.data[i]}
//
//func dayS() {
//	Sunday    := day{0, "SUN", "Sunday"}
//	Monday    := day{1, "MON", "Monday"}
//	Tuesday   := day{2, "TUE", "Tuesday"}
//	Wednesday := day{3, "WED", "Wednesday"}
//	Thursday  := day{4, "THU", "Thursday"}
//	Friday    := day{5, "FRI", "Friday"}
//	Saturday  := day{6, "SAT", "Saturday"}
//
//	data := []*day{&Tuesday, &Sunday, &Saturday, &Wednesday, &Thursday, &Friday, &Monday}
//	a := dayArray{data}
//	Sort(&a)
//	if !IsSorted(&a) {
//		panic("fails")
//	}
//
//	fmt.Print("The sorted array is: ")
//	for _, d := range data {
//		fmt.Printf("%s ", d.longName)
//	}
//	fmt.Println()
//}
//
//func main() {
//	intS()
//	stringS()
//	dayS()
//}

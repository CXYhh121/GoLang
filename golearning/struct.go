package main

import "fmt"

//type Person struct {
//	firstName string
//	lastName string
//}
//
//type TagType struct { // tags
//	field1 bool   "An important answer"
//	field2 string "The name of the thing"
//	field3 int    "How much there are"
//}
//
//func refTag(tt TagType, ix int) {
//	ttType := reflect.TypeOf(tt)
//	ixField := ttType.Field(ix)
//	fmt.Printf("%v\n", ixField.Tag)
//}
//
//func upPerson(p *Person) {
//	p.firstName = strings.ToUpper(p.firstName)
//	p.lastName = strings.ToUpper(p.lastName)
//}
//
//func main() {
//	tt := TagType{true, "Barak Obama", 1}
//	for i := 0; i < 3; i++ {
//		refTag(tt, i)
//	}
//
//	//1-struct as a value type
//	var per1 Person
//	per1.firstName = "Chris"
//	per1.lastName = "Woodward"
//	upPerson(&per1)
//
//	fmt.Printf("The name of the person is %s %s\n", per1.firstName, per1.lastName)
//
//	//2-struct as a pointer
//	per2 := new(Person)
//	per2.firstName = "Chris"
//	per2.lastName = "Woodward"
//	upPerson(per2)
//
//	fmt.Printf("The name of the person is %s %s\n", per2.firstName, per2.lastName)
//
//	//3-struct as a literal
//	per3 := &Person{"Chris", "Woodward"}
//	upPerson(per3)
//	fmt.Printf("The name of the person is %s %s\n", per3.firstName, per3.lastName)
//}

//匿名字段和内嵌结构体
type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	a int
	b float64
	int
	innerS
}

func main() {
	outer := new(outerS)
	outer.a = 87
	outer.b = 8.097
	outer.in1 = 98
	outer.in2 = 63

	fmt.Printf("outer.b is: %f\n", outer.b)
	fmt.Printf("outer.c is: %d\n", outer.a)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}

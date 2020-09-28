package basic

import (
	"fmt"
	"strings"
)

//strings and strconv 包

func main() {
	// HasPrefix  HasSuffix
	var str1 string = "This id an example of a string"
	fmt.Printf("T/F?Does this string \"%s\" have prefix %s? ", str1, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str1, "Th"))

	//Contains
	var ret = strings.Contains("string", "ing")
	fmt.Println(ret)

	//index_in_string.go
	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	//count_substring.go
	var str2 string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggggg"

	fmt.Printf("Number of H's in %s is: ", str2)
	fmt.Printf("%d\n", strings.Count(str2, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	//ToLower ToUpper
	var orig string = "Hey, how are you George"
	var lower string
	var upper string

	fmt.Printf("The original string: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lower string: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The upper string: %s\n", upper)
}

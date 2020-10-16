package main

import (
	"flag"
	"os"
)

// os_args.go

//func main() {
//	who := "Alice "
//	if len(os.Args) > 1 {
//		who += strings.Join(os.Args[1:], " ")
//	}
//	fmt.Println("Good Morning", who)
//}

//Flag 包

type Value struct {
}

type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}

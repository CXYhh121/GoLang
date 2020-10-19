package main

// os_args.go

//func main() {
//	who := "Alice "
//	if len(os.Args) > 1 {
//		who += strings.Join(os.Args[1:], " ")
//	}
//	fmt.Println("Good Morning", who)
//}

//Flag 包

//type Value struct {
//}
//
//type Flag struct {
//	Name     string // name as it appears on command line
//	Usage    string // help message
//	Value    Value  // value as set
//	DefValue string // default value (as text); for usage message
//}
//
//var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool
//
//const (
//	Space   = " "
//	Newline = "\n"
//)
//
//func main() {
//	flag.PrintDefaults()
//	flag.Parse() // Scans the arg list and sets up flags
//	var s string = ""
//	for i := 0; i < flag.NArg(); i++ {
//		if i > 0 {
//			s += " "
//			if *NewLine { // -n is parsed, flag becomes true
//				s += Newline
//			}
//		}
//		s += flag.Arg(i)
//	}
//	os.Stdout.WriteString(s)
//}

//用buffer读取文件
//func cat(r *bufio.Reader) {
//	for {
//		buf, err := r.ReadBytes('\n')
//		if err == io.EOF {
//			break
//		}
//		fmt.Fprint(os.Stdout, "%s", buf)
//	}
//	return
//}
//
//func main() {
//	flag.Parse()
//	if flag.NArg() == 0 {
//		cat(bufio.NewReader(os.Stdin))
//	}
//	for i := 0; i < flag.NArg(); i++ {
//		f, err := os.Open(flag.Arg(i))
//		if err != nil {
//			fmt.Fprint(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0],flag.Arg(i), err.Error())
//		}
//		cat(bufio.NewReader(f))
//	}
//}

//用切片读写文件
//func cat(f *os.File) {
//	const NBUF = 512
//	var buf [NBUF]byte
//	for {
//		switch nr, err := f.Read(buf[:]); true {
//		case nr < 0:
//			fmt.Fprintf(os.Stderr, "cat: error reading:%s\n", err.Error())
//			os.Exit(1)
//		case nr == 0: // EOF
//			return
//		case nr > 0:
//			if nw, ew := os.Stdout.Write(buf[0: nr]); nw != nr {
//				fmt.Fprintf(os.Stderr, "Cat: error writing: %s\n", ew.Error())
//			}
//
//		}
//	}
//}
//
//func main() {
//	flag.Parse() // scans the arg list and sets up flags
//	if flag.NArg() == 0 {
//		cat(os.Stdin)
//	}
//	for i := 0; i < flag.NArg(); i++ {
//		f, err := os.Open(flag.Arg(i))
//		if f == nil {
//			fmt.Fprintf(os.Stderr, "cat: can't open %s: error %s\n", flag.Arg(i), err)
//			os.Exit(1)
//		}
//		cat(f)
//		f.Close()
//	}
//}

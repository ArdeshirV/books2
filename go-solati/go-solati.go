// go-solati.go: My practices about "The Go programming language reference by Solati"
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

// Main entry point
func main() {
	defer func() {
		//waiteForEnter()
	}()
	PerformTitle()

	//mainChapterOne()
	//mainChapterTwo()
	//mainChapterThree()
	//mainReviewChapterOneAndTwo()
	//mainChannels()
	//mainChannels2()
	//mainChannels3()//
	//UsingRecover()
	//mainANewStepForward()
	//mainAncientAlphabets()
	//mainWriterReader()
	//mainWriterReader2()
	//mainFiles()
	mainSockets()
}

func mainSockets() {
	listner, err := net.Listen("tcp", "localhost:5060")
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	fmt.Println("Server started ...")

	i := 1
	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn, i)
		i++
	}
}

func handleConnection(conn net.Conn, i int) {
	fmt.Println("New connections number:", i)
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in reading the connection:%v", err)
		}
		fmt.Printf("Client %d: %s\n", i, string(buffer[:n]))
	}
}

func mainFiles() {
	fmt.Println("Files in Golang")

	filename := "/home/asohishn/d/sample-file.txt"
	fmt.Println("GetPageSize() = ", os.Getpagesize())

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(f, "Hello World")
	if err != nil {
		panic(err)
	}
	f.Close()

	f, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buff, err := io.ReadAll(f)
	fmt.Println(string(buff))
}

func mainWriterReader2() {
	fmt.Println("Reader & Writer 2")
	reader := strings.NewReader("This is a sample text")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func mainWriterReader() {
	fmt.Println("Reader & Writer")
	buff := []byte("Hello, World!")
	//os.Stdin.Read(buff)
	var s StringX
	s.Write(buff)
	fmt.Println(s)

	var ss String
	ss.Write(buff)
	fmt.Println(ss)

	r := io.TeeReader(&ss, os.Stdout)
	io.ReadAll(r)
	fmt.Println()

	fmt.Println()
	rx := bufio.NewReader(&ss)
	wx := bufio.NewWriter(&ss)

	wx.WriteString("Hello, World!")
	wx.Flush()
	str, err := rx.ReadString('\n')
	fmt.Println(str, err)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	fmt.Println("Hello", scanner.Text())
}

func (s StringX) ToString() string {
	return s.data
}

func mainAncientAlphabets() {
	WriteAlphabets(0x103A0, 0x103DF)
	WriteAlphabets(0x10B00, 0x10B3F)
	WriteAlphabets(0x10B40, 0x10B5F)
	WriteAlphabets(0x10B60, 0x10B7F)
	WriteAlphabets(0xFB50, 0xFDFF)
}

func WriteAlphabets(a, b rune) {
	for alph := a; alph <= b; alph++ {
		fmt.Print(string(alph), " ")
	}
	fmt.Println()
	fmt.Println()
}

func mainANewStepForward() {
	fmt.Println("A new step forward")
	str := "اردشیرThe Go Programming Language"
	buff := make([]byte, 20)
	var s String
	s.Write([]byte(str))
	s.Read(buff)
	fmt.Println(string(buff))
	for i, v := range buff {
		fmt.Printf("[%d]=%v ", i, string(v))
	}
	fmt.Println(s)
	fmt.Println()
	r := rune(buff[4]) | rune(buff[5])<<16
	fmt.Println(string(buff[0:12]))
	fmt.Println(string(r))
	fmt.Printf("%T\n", r)
	rx := rune(194) | rune(169)<<8
	fmt.Println("[", int(rx), "] =", string(rx))
}

type String struct {
	pos  int
	data string
}

type StringX = String

func (s String) String() string {
	return s.data
}

func (s *String) Write(b []byte) (n int, err error) {
	s.data += string(b)
	return len(b), nil
}

func (s *String) Read(b []byte) (n int, err error) {
	n = copy(b, s.data[s.pos:])
	s.pos += n
	if s.pos >= len(s.data) {
		err = io.EOF
	}
	return n, err
}

func UsingRecover() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Recovered from:", rec)
		}
	}()
	panic("Panic happened")
}

func mainChannels3() {
	n := 1000000
	RegisteredUsers := make(chan int, n)
	limit := make(chan bool, runtime.NumCPU())
	now := time.Now()

	for i := 1; i <= n; i++ {
		go func(in int) {
			limit <- true
			RegisteredUsers <- in
			<-limit
		}(i)
	}

	for n > 0 {
		<-RegisteredUsers
		n--
	}

	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Durations:", time.Since(now))
	fmt.Printf("%#T\n", time.Second)
}

func mainChannels2() {
	c := make(chan int, 7)
	go func() {
		for i := 100; i >= 0; i-- {
			time.Sleep(time.Millisecond * 300)
			c <- i
		}
	}()
	go func() {
		for i := 1000; i >= 0; i-- {
			time.Sleep(time.Millisecond * 100)
			c <- i
		}
	}()
	for data := range c {
		fmt.Print(data, " ")
		if data == 0 {
			//close(c)
			break
		}
	}
	fmt.Printf("\n%#T, len(c) = %v\n", c, len(c))
	c <- 10
	fmt.Printf("\n%#T, len(c) = %v\n", c, len(c))
	fmt.Println()
	close(c)
}

func mainChannels1() {
	c := make(chan int, 4)
	for i := range 100 {
		c <- i
		go DoHeavyTask(c)
	}
	close(c)
	fmt.Println()
}

func DoHeavyTask(c chan int) {
	time.Sleep(200 * time.Millisecond)
	data := <-c
	fmt.Printf("%v ", data)
}

// Chapter three
func mainChapterThree() {
	fmt.Printf("  Chapter Three: \n\n")
	fmt.Println("Hello World")
	Integrate("Hello from mainChapterThree +/-")
}

// Integrate ax^2+bx+c
func Integrate(expr string) float64 {
	fmt.Println(expr)
	return 0.0
}

func mainReviewChapterOneAndTwo() {
	printTitle("Review chapter one and two\n")

	const (
		a1 = (iota + 1) * 1024
		a2
		a3
		a4
		_
		a6
	)
	fmt.Printf("a1 = %v, a6 = %v\n", a1, a6)

	arr1 := []uint8{0: 1, 2: 2, 1: 3, 3: 4}
	fmt.Printf("%v\n", arr1)

	m1 := make(map[string]string)
	m1["شنبه"] = "saturday"
	m1["یکشنبه"] = "sunday"
	m1["دوشنبه"] = "monday"
	m1["سهشنبه"] = "tuesday"
	for k, v := range m1 {
		fmt.Printf("%v:%v\n", k, v)
	}
	k := "شنبه"
	fmt.Println(k, "=", m1[k])
	value, ok := m1["حمعه"]
	fmt.Printf("value = %v, ok = %v\n", value, ok)

	type MyType int
	var m MyType
	m = 9
	fmt.Println(m)

	type operator = func(int, int) int
	add := func(a, b int) int { return a + b }
	res := func(a, b int, op operator) int {
		return op(a, b)
	}(1, 2, add)
	fmt.Println("res = ", res)

	p1 := player{
		name: "nightworlf",
		hp:   100,
	}
	fmt.Println("before hit:", p1.hp)
	p1.hit()
	fmt.Println("After hit:", p1.hp)

	//zebel()
}
func zebel() {
	arr := []int{1, 20, 33, 1, 12, 332, 4, 239, 4, 23, 434, 9, 46, 90, 95, 439, 3, 3, 0, 99}
	fmt.Println(arr)
	total1 := 0
	now1 := time.Now()
	for range 1000 {
		total1 += addAllOne(arr)
	}
	fmt.Println("Total: ", total1)
	fmt.Println("Duration: ", time.Since(now1))
	total2 := 0
	now2 := time.Now()
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		for range 250 {
			total2 += addAllOne(arr)
		}
		wg.Done()
	}()
	go func() {
		for range 250 {
			total2 += addAllOne(arr)
		}
		wg.Done()
	}()
	go func() {
		for range 250 {
			total2 += addAllOne(arr)
		}
		wg.Done()
	}()
	go func() {
		for range 250 {
			total2 += addAllOne(arr)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Total: ", total2)
	fmt.Println("Duration: ", time.Since(now2))
}

func addAllOne(a []int) int {
	total := 0
	t := 0
	for range 1000000 {
		for _, v := range a {
			total += v
		}
		t += total
	}
	return total
}

type player struct {
	name string
	hp   int
}

func (p *player) hit() {
	p.hp--
}

// find n max of array
func findNMax(arr []int, n int) []int {
	if n > len(arr) {
		n = len(arr)
	}
	result := make([]int, n)
	copy(result, arr)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result[:n]
}

func findMaxPrime(n int) int {
	maxPrime := 2
	for i := 3; i <= n; i++ {
		if isPrime(i) && i > maxPrime {
			maxPrime = i
		}
	}
	return maxPrime
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func rev(text string) string {
	return text
}

func EvalMathematicExpr(expr string) float64 {
	fmt.Print("Hello")
	return 0.0
}

func integrate() {
	fmt.Println("&&&")
}

func mainChapterOne() {
	fmt.Printf("  Chapter One: \n\n")

	fmt.Println("Hello World")
	fmt.Println(`Hello, World!
in multilines.`)

	var age int
	var name string
	var movie string
	var score float64
	name = "Parmis"
	age = 13
	movie = "Hotel Transylvania"
	score = 7
	fmt.Println(name, "is a good student")
	fmt.Println(name, "is", age, "years old")
	fmt.Println(movie, name, "favorite movie score is", score)

	age = 39
	name = "اردشیر"
	fmt.Println(name, "برنامه نویس کامپیوتر است و ", age, "سال سن دارد")

	const (
		a = iota + 1
		b
		c
		_
		e
	)
	fmt.Println(a, b, c, e)

	var bit1 uint8 = 0b11101101
	fmt.Printf("%b, %b\n", bit1, ^bit1)

	str1 := "Hello"
	str2 := "World"
	str3 := fmt.Sprintf("%s, %s!", str1, str2)
	fmt.Println(str3)

	// TODO: Add your code here://
	fmt.Printf("%v\n", str1 != str2)
	const bit = 24
	fmt.Println(bit >> 2)
	const b1, b2 = 0b00110000, 0b00100001
	fmt.Printf("%b\n", b1^b2)

	gender := "male"
	if gender == "female" {
		fmt.Println("You are a woman")
	} else {
		fmt.Println("You are a man")
	}

	const day = 2
	dayName := ""

	switch day {
	case 0:
		dayName = "Saturday"
	case 1:
		dayName = "Sunday"
	case 2:
		dayName = "Monday"
	case 3:
		dayName = "Tuesday"
	case 4:
		dayName = "Wednsday"
	case 5:
		dayName = "Tursday"
	case 6:
		dayName = "Friday"
	default:
		dayName = "Xday"
	}

	arrDayNames := [7]string{"Saturday", "Sunday",
		"Monday", "Tuesday", "Wednsday", "Friday"}

	fmt.Println(dayName)
	fmt.Println(arrDayNames[day])

	goto next_statement
	fmt.Println("Khameneee is a dirty pig")
next_statement:
	fmt.Println("Yeah we are here")

	// Functions begins here:
	helloFunc()
	c1 := newCounter()
	c2 := newCounter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())

	expr := "2*3+4/2*2-1"
	operators := make(map[string]func(int, int) int)
	operators["+"] = add
	operators["-"] = sub
	operators["/"] = div
	operators["*"] = mul
	res := 0
	prev := 0
	newExpr := ""
	for i, v := range expr {
		item := string(v)
		if item == "*" || item == "/" {
			next := int(expr[i+1])
			res = operators[item](prev, next)
			newExpr += fmt.Sprintf("%d ", res)
		}
		prev = int(v)
	}
	for i, v := range expr {
		item := string(v)
		if item == "+" || item == "-" {
			next := int(expr[i+1])
			res = operators[item](prev, next)
			newExpr += fmt.Sprintf("%d ", res)
		}
		prev = int(v)
	}
	fmt.Println("\n", newExpr)

	// Stone 3
	fmt.Println("Stone 3")

}

func mul(a, b int) int {
	return a * b
}

func div(b, a int) int {
	return b / a
}

func add(a, b int) int {
	return a + b
}

func sub(b, a int) int {
	return b - a
}

func newCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func helloFunc() {
	fmt.Println("helloFunc")
}

func mainChapterTwo() {
	fmt.Printf("  Chapter Two: \n\n")
	fmt.Println("Hello World")
}

func PrintTitle() {
	blnColor := true
	strAppName := "go-solati"
	strAppYear := "2025"
	strAppDescription := "The Go Programming Language"
	strVersion := "1.0"
	strLicense := "GPLv3+"
	strCopyright := "https://github.com/ArdeshirV/book/go-solati"
	fmt.Print(FormatTitle(strAppName, strAppDescription, strVersion, blnColor))
	fmt.Print(FormatCopyright(strAppYear, strCopyright, strLicense, blnColor))
}

func FormatTitle(strAppName, strAppDescription, strVersion string, blnColor bool) string {
	NoneColored := "%v - %v Version %v\n"
	Colored := "\033[1;33m%v\033[0;33m - %v \033[1;33mVersion %v\033[0m\n"
	var strFormat string
	if blnColor {
		strFormat = Colored
	} else {
		strFormat = NoneColored
	}
	return sprintf(strFormat, strAppName, strAppDescription, strVersion)
}

func FormatCopyright(strAppYear, strCopyright, strLicense string, blnColor bool) string {
	NoneColored := "Copyright (c) %v %v, Licensed under %v\n"
	Colored := ("\033[0;33mCopyright (c) \033[1;33m%v \033[1;34m%v" +
		"\033[0;33m, \033[1;33m%v\033[0m\n")
	var strFormat string
	if blnColor {
		strFormat = Colored
	} else {
		strFormat = NoneColored
	}
	return sprintf(strFormat, strAppYear, strCopyright, strLicense)
}

func sprintf(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

func PerformTitle() {
	PrintTitle()
	bookName := "The Go programming language reference"
	title := "\n    %sMy Practices about \"%s%s%s\" %sʕ◔ϖ◔ʔ%s\n\n"
	fmt.Printf(title, MAGENTA, BMAGENTA, bookName, MAGENTA, BGREEN, TEAL)
}

func printTitle(title string) {
	fmt.Printf("\n  \033[1;36m%s\033[0;36m\n", title)
}

const (
	NORMAL   = "\033[0m"
	BOLD     = "\033[1m"
	RED      = "\033[0;31m"
	TEAL     = "\033[0;36m"
	WHITE    = "\033[0;37m"
	BLUE     = "\033[0;34m"
	GREEN    = "\033[0;32m"
	YELLOW   = "\033[0;33m"
	MAGENTA  = "\033[0;35m"
	BRED     = "\033[1;31m"
	BBLUE    = "\033[1;34m"
	BTEAL    = "\033[1;36m"
	BWHITE   = "\033[1;37m"
	BGREEN   = "\033[1;32m"
	BYELLOW  = "\033[1;33m"
	BMAGENTA = "\033[1;35m"
)

// Problems in book
// P29: sudo -C /usr/local -xzf flle-name --> tar -xvf file-name
// P44: پیدا کردن کوحکترین عدد برای یافتن بزرکترین عدد
// P55: true || false ==> false
// P54: a < b, a < b
// P79: تعریف تابغ به خط فارسی برعکس شده است

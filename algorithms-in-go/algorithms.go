package main

import (
	"fmt"
)

var (
	arr = []int{9, 2, 3, 7, 6, 1, 8, 4, 0, 5}
)

func main() {
	defer fmt.Print(NORMAL)
	title := "\n    %sAlgorithms in %sGo%slang %sʕ◔ϖ◔ʔ%s\n\n"
	fmt.Printf(title, BMAGENTA, BBLUE, BLUE, BGREEN, TEAL)

	mainBubbleSort1()
}

func mainBubbleSort1() {
	fmt.Println("mainBubbleSort1()")

	counter := 0
	fmt.Println(arr)
	for i := range len(arr) {
		for j := range i {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			counter++
			fmt.Print(NORMAL, "[", counter, "]=", arr, TEAL, "\n")
		}
	}
	fmt.Println(arr)
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

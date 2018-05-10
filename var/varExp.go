package main

import "fmt"

var (
	// usually used to define global var
	gvStr string = "global str"
	gvNum int    = 1000000
)

func main() {
	var vTest string
	vTest = "abc"

	var vTest2 string = "abcd"

	vTest3 := "abcde"

	var vNum1, vNum2, vNum3 = 1, 2, 3 // multi-definition

	// vTest4 = "abcdef"  error definition

	fmt.Println(vTest)
	fmt.Println(vTest2)
	fmt.Println(vTest3)

	fmt.Println(vNum1)
	fmt.Println(vNum2)
	fmt.Println(vNum3)

	fmt.Println(gvStr)
	fmt.Println(gvNum)
}

package utils

import "fmt"

func printInformationOfSlice[K int | float32 | float64](s *[]K, name string) {
	fmt.Println("The information is displayed for the slice ", name)
	fmt.Println("The values are ", *s)
	fmt.Printf("The length of this slice is %d and the capacity of this slice is %d", len(*s), cap(*s))
	fmt.Printf("The address of this slice is %p\n", *s)
	fmt.Println("===================================")

}

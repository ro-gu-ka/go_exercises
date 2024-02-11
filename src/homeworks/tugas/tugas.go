package main

import (
	"fmt"

	// "homework.com/tahap1"
	"homework.com/tahap2"
)

func main() {
	// tahap1.PrintDifference("Hello Sandbox HSI, in ", 2024)
	// tahap1.TestError()

	nik := tahap2.GenerateNik("female", 2009, 1)
	fmt.Println(nik)
}

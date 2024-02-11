package main

import (
	"fmt"
	"sort"

	// "homework.com/tahap1"
	"homework.com/tahap2"
)

func main() {
	// tahap1.PrintDifference("Hello Sandbox HSI, in ", 2024)
	// tahap1.TestError()

	// tugas 2
	fmt.Println("Random NIK: ")
	nikColl := tahap2.GenerateNikCollection("m", 2021, 10)
	fmt.Printf("%q\n", nikColl)

	fmt.Println("Consecutive NIK: ")
	nikConsecutive := tahap2.GenerateConsecutiveNextNik("ARN212-00122", 10)
	fmt.Printf("%q\n", nikConsecutive)

	fmt.Println("Sorted NIK: ")
	coll := []string{"ARN212-00813", "ARN192-00197", "ARN202-00770", "ARN212-00089", "ARN221-00892", "ARN242-00517", "ARN182-00825", "ARN191-00179", "ARN211-00542", "ARN212-00318"}
	sort.Strings(coll)
	fmt.Printf("%q\n", coll)
}

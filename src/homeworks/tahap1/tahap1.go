package tahap1

import (
	"errors"
	"fmt"
)

func PrintDifference(message string, year int) {
	// temporary buffer
	var temp string

	temp = fmt.Sprintln(message, year)

	// printing out the declared variables as a single string
	fmt.Print(temp)

}

var (
	fromErrorNew = errors.New("uh oh")
)

func validateValue(number int) error {
	if number == 1 {
		return fmt.Errorf("oops")
	} else if number == 2 {
		return fromErrorNew
	}
	return nil
}

func TestError() {
	for num := 1; num <= 3; num++ {
		fmt.Printf("validating %d... ", num)
		err := validateValue(num)
		if err == fromErrorNew {
			fmt.Println("This is from error.New using sentinel check.")
		} else if err != nil {
			fmt.Println("This is from fmt.Errorf with message:", err)
		} else {
			fmt.Println("valid!")
		}
	}
}

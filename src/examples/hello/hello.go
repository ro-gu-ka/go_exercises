package main

import (
	"fmt"
	// "log"

	// "example.com/greetings"

	"golang.org/x/example/hello/reverse"
)

func main() {
	// // Set properties of the predefined Logger, including
	// // the log entry prefix and a flag to disable printing
	// // the time, source file, and line number.
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// // Request a greeting message.
	// message, err := greetings.Hello("Rony")
	// // If an error was returned, print it to the console and
	// // exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // If no error was returned, print the returned message
	// // to the console.
	// fmt.Println(message)

	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}

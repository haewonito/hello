// Declare a main package. In Go, code 
// executed as an application must be in a main package.
package main

// Configure the log package to print the command name ("greetings: ") 
// at the start of its log messages, without a time stamp or source file information.
import (
    "fmt"
	"log"
    "example.com/greetings"
)
// 

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)

}
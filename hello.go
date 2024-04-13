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
    // Get a greeting message and print it.
    //original code before error handling:
	// message := greetings.Hello("Gladys")
    // fmt.Println(message)


	    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("Haewon")
    // If an error was returned, print it to the console and
    // exit the program.
	//Fatal function of log will print the error and stop the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)

}
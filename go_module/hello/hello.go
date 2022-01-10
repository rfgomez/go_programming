package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	// *** Name ***********
	// Request a greeing message.
	message1, err := greetings.Hello("Rodolfo")
	
	// If an error was returned, print it to the console and
	// exit the progrem.
	if err != nil {
		log.Fatal(err)
	}
	
	// if no error was returned, print returned message
	// to the console.
	fmt.Println(message1)
	
	// *** Names **********
	// A slice of names.
    names := []string{"Keith", "Steve", "John"}
	
	// Request a greeing message.
	message2, err := greetings.Hellos(names)
	
	// If an error was returned, print it to the console and
	// exit the progrem.
	if err != nil {
		log.Fatal(err)
	}
	
	// if no error was returned, print returned message
	// to the console.
	fmt.Println(message2)
}
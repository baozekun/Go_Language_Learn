package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Simple.
	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	//Multiple.
	names := []string{"张三", "李四", "王五"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	a := fmt.Sprintln("Hello, enter with your name")
	fmt.Println(a)
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(name)
	fmt.Println(Hello(name))

}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi %v, Welcome ", name)
	return message
}

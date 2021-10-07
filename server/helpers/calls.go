package helpers

import (
	"log"
)

// HelloWorld -- as the name goes
func HelloWorld() {
	log.Println("Hello, World!")
}

// Hello -- appending hello before the arg string
func Hello(n []interface{}) {
	name := n[0].(map[string]interface{})["name"]
	log.Printf("Hello, %s", name)
}

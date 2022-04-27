package main

import (
	"fmt"

	"github.com/kimj99/golang/util"
)

func main() {
	greetings := fmt.Sprintf("Hello %s", "Jiman")
	fmt.Println(greetings)

	fmt.Printf("Test length of %d characters", util.StringLength(greetings))
}

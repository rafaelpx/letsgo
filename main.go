package main

import (
	"fmt"

	"github.com/rafaelpx/letsgo/greetings"
	"github.com/rafaelpx/letsgo/integers"
)

func main() {
	fmt.Println(greetings.Hello("Rafa", "French"))
	fmt.Println(integers.Add(2, 2))
}

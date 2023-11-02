package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("My first GO Project! Hello world!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Println(text)
}

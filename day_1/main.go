package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	book := scanner.Text()
	fmt.Println(book, "- лучшая книга!")
}

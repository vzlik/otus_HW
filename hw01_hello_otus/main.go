package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	myStr := "Hello, OTUS!"

	fmt.Println(stringutil.Reverse(myStr))
}

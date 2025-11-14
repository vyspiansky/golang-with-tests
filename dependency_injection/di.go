package main

import (
	"fmt"
	"io"
)

// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Printf("Hello, %s", name)
// }

// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

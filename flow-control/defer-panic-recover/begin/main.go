// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	defer cleanup("123")
	defer cleanup("456")

	// defer recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from an error state:", err)
		}
	}()

	// panic
	panic("Error occurred")
}

// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 10
	// print the value of the local variable "val"
	fmt.Printf("%v=%T\n", val, val)
	// print the value of the package-level variable "val"
	pkgvar()
	// update the package-level variable "val" and print it again
	pkgvarupdate("global_update")
	pkgvar()

	// print the pointer address of the local variable "val"
	fmt.Printf("%v=%T\n", &val, &val)
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 11
	fmt.Println(val)
}

func pkgvar() {
	fmt.Println(val)
}

func pkgvarupdate(val_new string) {
	val = val_new
}

package main

import "fmt"

//
var name string = "Gopher นำแน่!!!"

func main() {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("type: %T\n", name)
}

func fucnum() {
	// จะ Error เมื่อตัวแปร
	fmt.Printf("name: %v\n", name)

}

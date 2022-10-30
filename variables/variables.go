package main

import "fmt"

//
var name string = "Gopher นำแน่!!!"

// สามารถประกาศตัวแปรโดย ไม่ใส่ type ได้
var name1 = "Gopher นำแน่!!!"

const day = "Monday"

func main() {
	// เมื่อตัวแปรไม่ได้ใช้ จะแจ้งเตือน ทันที่เมื่อ run
	num := 10
	fmt.Printf("name: %v\n", name)
	// %T แสดง type
	fmt.Printf("type: %T\n", num)
}

func fucnum() {

	// จะ Error เมื่อตัวแปร อยู่ใน function อื่น
	fmt.Printf("name: %v\n", name1)

}

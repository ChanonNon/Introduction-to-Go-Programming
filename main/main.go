// ไฟล์ทุกไฟล์ต้องเริ่มต้นที่ func main() ใน package main เสมอ
package main

import "fmt"

func main() {
	fmt.Println("hello Gopher!!!")

	// %s ใช้เฉพาะ string
	fmt.Printf("hello %s!!!!!\n", "Gopher")
	fmt.Println("next line")

	// %d ใช้เฉพาะ int
	fmt.Printf("hello %d!!!!!\n", 66)

	// %v สามารถทำได้ทั้ง string กับ int ได้
	fmt.Printf("hello %v!!!!!\n", "Gopher")
	fmt.Printf("hello %v!!!!!\n", 66)

	// %v สามารถต่อกันนะ
	fmt.Printf("hello %v : %v!!!!!\n", "Gopher", 66)

}

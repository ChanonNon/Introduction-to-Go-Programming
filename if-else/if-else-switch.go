package main

import (
	"fmt"
	"time"
)

func main() {
	num := 14

	if num%2 == 0 {
		fmt.Println("เลขคู่")
	} else if num%2 != 0 {
		fmt.Println("เลขคี่")
	} else {
		fmt.Println("ไม่สามารถระบุได้")
	}

	today := time.Wednesday
	switch today {
	case time.Monday:
		fmt.Println("today is Monday")
		// fallthrough สามารถ run case ถัดไปได้
		fallthrough
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Printf("สวัสดีวัน %v อยู่ดีมีแฮงเดย์\n", today)
	}

}

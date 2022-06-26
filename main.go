package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello World!")
	name := "DevOps"
	fmt.Printf("Type of name:= %T \n and size is:= %d \n", name, unsafe.Sizeof(name))

	fmt.Println(rectProps(4, 5))
	fmt.Println(rectPropsOne(4, 5))

	fmt.Println(geo.Area(3, 4))
}

func rectProps(length, width float32) (float32, float32) {
	area := length * width
	perimeter := (length + width) / 2
	return area, perimeter
}

func rectPropsOne(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) / 2
	return
}

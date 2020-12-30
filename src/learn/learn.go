package main

import (
	"fmt"
)

//文件单位
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Printf("KB %d\n", KB)
	fmt.Printf("MB %d\n", MB)
	fmt.Printf("GB %d\n", GB)
	fmt.Printf("TB %d\n", TB)
	fmt.Printf("PB %d\n", PB)
	InitArray()
}

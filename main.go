package main

import (
	"runtime"
)

func main() {
	println("Hello "+runtime.GOOS+"!")
}

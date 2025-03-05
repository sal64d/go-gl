package main

import (
	"go-rpg/internal/triangle"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	triangle.Main()
}

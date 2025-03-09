package main

import (
	"go-rpg/internal/home"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	home.Main()
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello flake")
	dispatchServiceProxy("0.0.0.0", 3000)
}

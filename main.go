package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello flake")
	proxy.dispatchServiceProxy("0.0.0.0", 3000)
}

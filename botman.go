package main

import (
	"fmt"
	routers "Backpack/routes"
)

func main() {
	fmt.Println("Main is running")
	routers.HandleReq()
}

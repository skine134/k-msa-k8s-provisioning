package main

import (
	"fmt"
	_ "project/createDeployment"
	_ "project/createNamespace"
	_ "project/createService"
)

func main() {
	fmt.Println("start go prog")
}

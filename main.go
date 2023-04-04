package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type deployStruct struct {
	name         string
	namespace    string
	replicas     string
	image        string
	imageVersion string
}

func main() {
	ns := bufio.NewReader(os.Stdin)

	fmt.Println("이름 입력")
	dname, err := ns.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	namespace := deployStruct{
		name: dname,
	}

	fmt.Println(namespace.name)
}

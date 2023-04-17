package main

import (
	"fmt"

	st "github.com/jjsair0412/embodiment"
)

func act(cd st.CreateDeployment) {
	cd.CreateDeployment()
}

func main() {

	deployTest := st.Deployment_st{
		name:         "test",
		namespace:    "tesetNamespace",
		replcas:      15,
		image:        "nginx",
		imageVersion: "latest",
	}

	act(deployTest)

	fmt.Println("success!")
}

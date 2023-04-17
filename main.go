package main

import (
	"fmt"

	st "github.com/jjsair0412/embodiment"
)

func act(cd st.CreateDeployment) {
	st.CreateDeployment()
}

func main() {

	deployTest := st.Deployment_st{"test", "tesetNamespace", 15, "nginx", "latest"}

	act(deployTest)

	fmt.Println("success!")
}

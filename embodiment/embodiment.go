package embodiment

import (
	"fmt"

	inter "github.com/jjsair0412/interfaces"
)

// deployment 생성
type CreateDeployment interface {
	CreateDeployment()
}

// namespace 생성
type CreateNamespace interface {
	Createnamespace()
}

type CreateService interface {
	CreateService()
}

func (D inter.Deployment_st) CreateDeployment(gname, gnamespace, gimage, gimageVersion string, greplcas int) {
	testDeploy := inter.Deployment_st{gname, gnamespace, greplcas, gimage, gimageVersion}
	fmt.Println(testDeploy)
}

func (N inter.Namespace_st) CreateNamespace(gname string) {
	testNamespace := inter.Namespace_st{gname}
	fmt.Println(testNamespace)
}

func (S inter.Service_st) CreateService(gname, gscType string, gexportPort int) {
	testService := inter.Service_st{gname, gscType, gexportPort}
	fmt.Print(testService)
}

package embodiment

import (
	"fmt"
)

func Test() {
	fmt.Println("hi")
}

// deployment 구조체
type Deployment_st struct {
	name         string
	namespace    string
	replcas      int
	image        string
	imageVersion string
}

// Namespace 구조체
type Namespace_st struct {
	name string
}

// service 구조체
type Service_st struct {
	name       string
	scType     string
	exportPort int
}

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

func (D *Deployment_st) CreateDeployment(gname, gnamespace, gimage, gimageVersion string, greplcas int) {
	testDeploy := Deployment_st{gname, gnamespace, greplcas, gimage, gimageVersion}
	fmt.Println(testDeploy)
}

func (N *Namespace_st) CreateNamespace(gname string) {
	testNamespace := Namespace_st{gname}
	fmt.Println(testNamespace)
}

func (S *Service_st) CreateService(gname, gscType string, gexportPort int) {
	testService := Service_st{gname, gscType, gexportPort}
	fmt.Print(testService)
}

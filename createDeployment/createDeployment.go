import _ "k8s.io/client-go/plugin/pkg/client/auth"



type DeployStruct struct {
	name string,
	namespace string,
	replicas string,
	image string,
	imageVersion string,
}

// Deployment 배포
func (c createDeployment) deployDeployment() {
	
}

// Deployment 구조 생성
func (c DeployStruct) createDeployment() string {
	
}
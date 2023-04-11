package createDeployment

type DeployStruct struct {
	name         string
	namespace    string
	replicas     string
	image        string
	imageVersion string
}

// Deployment 구조 생성
func CreateDeploymentStruct(name, namespace, replicas, image, imageVersion string) *DeployStruct {
	return &DeployStruct{
		name:         name,
		namespace:    namespace,
		replicas:     replicas,
		image:        image,
		imageVersion: imageVersion,
	}
}

// Deployment yaml 생성
func CreateDeployment(CreateDeploymentStruct func()) {

}

package createDeployment

type deployStruct struct {
	name         string
	namespace    string
	replicas     string
	image        string
	imageVersion string
}

// Deployment 구조 생성
func (d deployStruct) CreateDeployment(gname string) deployStruct

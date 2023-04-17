package interfaces

import (
	_ "fmt"
)

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

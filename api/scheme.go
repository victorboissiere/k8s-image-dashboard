package api

type NamespacePods struct {
	Namespace string
	Pods []Pod
}

type Pod struct {
	Name string
	Containers []Container
	Date string
}

type Container struct {
	Image string
	HasRepositoryURL bool
	RepositoryURL string
	Status ContainerStatus
}

type ContainerStatus struct {
	Name string
	Color string
	Message string
	HasMessage bool
	IsReady bool
}

type Node struct {
	Name string
	Labels map[string]string
	Annotations map[string]string
	IP string
	KubeVersion string
}

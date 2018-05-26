package api

import (
	"fmt"
	"k8s.io/api/core/v1"
	"os"
	"strings"
	"regexp"
)

func getContainerStatus(pod v1.Pod, index int) ContainerStatus {
	containerState := pod.Status.ContainerStatuses[index].State
	isReady := pod.Status.ContainerStatuses[index].Ready

	if containerState.Running != nil {
		return ContainerStatus{
			Name: "Running",
			Color: "success",
			Message: "",
			HasMessage: false,
			IsReady: isReady,
		}
	} else if containerState.Waiting != nil {
		return ContainerStatus{
			Name: "Waiting",
			Color: "warning",
			HasMessage: true,
			Message: fmt.Sprintf("%s %s", containerState.Waiting.Reason, containerState.Waiting.Message),
			IsReady: isReady,
		}
	} else {
		return ContainerStatus{
			Name: "Terminated",
			Color: "dark",
			HasMessage: true,
			Message: fmt.Sprintf("%s %s", containerState.Terminated.Reason, containerState.Terminated.Message),
			IsReady: isReady,
		}
	}
}

func getRepositoryURL(image string) (string, bool) {
	repoRegex, exists := os.LookupEnv("REPO_REGEX"); if exists {
		regexParts := strings.Split(repoRegex, "|")
		pattern := regexp.MustCompile(regexParts[0])
		if pattern.MatchString(image) {
			return pattern.ReplaceAllString(image, regexParts[1]), true
		}
	}

	return "", false
}

func getContainers(pod v1.Pod) []Container {
	containers := make([]Container, len(pod.Spec.Containers))
	for i := 0; i < len(containers); i++ {
		containerItem := pod.Spec.Containers[i]
		repositoryURL, hasURL := getRepositoryURL(containerItem.Image)

		containers[i] = Container{
			Image: containerItem.Image,
			Status: getContainerStatus(pod, i),
			HasRepositoryURL: hasURL,
			RepositoryURL: repositoryURL,
		}
	}

	return containers
}

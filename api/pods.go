package api

import (
	"github.com/justincampbell/timeago"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

func GetPod(namespace string, podName string) Pod {
	pod, err := client.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		return Pod{
			Name: "NotFound",
			Namespace: namespace,
			Containers: []Container{},
			Date: "",
		}
	}

	return getPodMapping(*pod)
}

func getPods(namespace string) []Pod {
	podList, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	checkError(err)

	pods := make([]Pod, len(podList.Items))
	for i := 0; i < len(pods); i++ {
		pods[i] = getPodMapping(podList.Items[i])
	}

	return pods
}

func getPodMapping(pod v1.Pod) Pod {
	return Pod{
		Name: pod.Name,
		Namespace: pod.Namespace,
		Containers: getContainers(pod),
		Date: timeago.FromTime(pod.CreationTimestamp.Time),
	}
}

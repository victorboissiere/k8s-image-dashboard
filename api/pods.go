package api

import (
	"github.com/justincampbell/timeago"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getPods(namespace string) []Pod {
	podList, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	checkError(err)

	pods := make([]Pod, len(podList.Items))
	for i := 0; i < len(pods); i++ {
		podItem := podList.Items[i]

		pods[i] = Pod{
			Name: podItem.Name,
			Containers: getContainers(podItem),
			Date: timeago.FromTime(podItem.CreationTimestamp.Time),
		}
	}

	return pods
}

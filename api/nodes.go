package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

func getNodeIP(node v1.Node) string {
	addresses := node.Status.Addresses
	for i := 0; i < len(addresses); i++ {
		if addresses[i].Type == "InternalIP" {
			return addresses[i].Address
		}
	}

	return ""
}

func GetNodes() []Node {
	nodeList, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
	checkError(err)

	nodes := make([]Node, len(nodeList.Items))
	for i := 0; i < len(nodeList.Items); i++ {
		nodeItem := nodeList.Items[i]

		memoryCapacity, _ := nodeItem.Status.Capacity.Memory().AsInt64()

		nodes[i] = Node{
			Name: nodeItem.Name,
			Annotations: nodeItem.Annotations,
			Labels: nodeItem.Labels,
			IP: getNodeIP(nodeItem),
			KubeVersion: nodeItem.Status.NodeInfo.KubeletVersion,
			OSImage: nodeItem.Status.NodeInfo.OSImage,
			MemoryCapacity: memoryCapacity / 1024 / 1024,
		}
	}

	return nodes
}
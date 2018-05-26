package api

import (
	"os"
	"strings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getNamespaces() []string {
	namespaceList, err := client.CoreV1().Namespaces().List(metav1.ListOptions{})
	checkError(err)

	excludeNamespaces := ""
	excludeNamespaceEnv, exists := os.LookupEnv("EXCLUDE_NAMESPACES"); if exists {
		excludeNamespaces = excludeNamespaceEnv + ","
	}

	namespaces := make([]string, 0)
	for i := 0; i < len(namespaceList.Items); i++ {
		namespaceName := namespaceList.Items[i].Name
		if !strings.Contains(excludeNamespaces, namespaceName + ",") {
			namespaces = append(namespaces, namespaceName)
		}
	}

	return namespaces
}

func GetNamespacePods() []NamespacePods {
	namespaces := getNamespaces()
	namespacePods := make([]NamespacePods, len(namespaces))

	for i := 0; i < len(namespaces); i++ {
		namespacePods[i] = NamespacePods{
			Namespace: namespaces[i],
			Pods: getPods(namespaces[i]),
		}
	}

	return namespacePods
}


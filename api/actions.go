package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"errors"
	"k8s.io/api/apps/v1"
	"time"
	v12 "k8s.io/api/core/v1"
)

const RollingUpdateEnvKey = "LAST_FORCED_ROLLING_UPDATE"

func ArbitraryRollingUpdate(namespace string, podName string) bool {
	pod, err := client.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	checkError(err)

	replicatSetName, err := getFirstParentName(pod.OwnerReferences, "ReplicatSet")
	checkError(err)

	replicatSet, err := client.AppsV1().ReplicaSets(namespace).Get(replicatSetName, metav1.GetOptions{})
	checkError(err)

	deploymentName, err := getFirstParentName(replicatSet.OwnerReferences, "Deployment")
	checkError(err)

	deployment, err := client.AppsV1().Deployments(namespace).Get(deploymentName, metav1.GetOptions{})
	checkError(err)

	patchEnv(deployment, RollingUpdateEnvKey, time.StampMilli)

	return false
}

func patchEnv(deployment *v1.Deployment, key string, value string) {
	envPos := findEnvPos(deployment.Spec.Template.Spec.Containers[0].Env, key)
	if envPos == -1 {
		deployment.Spec.Template.Spec.Containers[0].Env[envPos].Value = value
	} else {
		deployment.Spec.Template.Spec.Containers[0].Env = append(deployment.Spec.Template.Spec.Containers[0].Env, v12.EnvVar{
			Name: key,
			Value:value,
		})
	}
}

func findEnvPos(envs []v12.EnvVar, name string) int {
	for i := 0; i < len(envs); i++ {
		if envs[i].Name == name {
			return i
		}
	}

	return -1
}

func getFirstParentName(references[] metav1.OwnerReference, parentKind string) (string, error) {
	for i := 0; i < len(references); i++ {
		if references[i].Kind == parentKind {
			return references[i].Name, nil
		}
	}

	return "", errors.New("Cannot find parent" + parentKind)
}


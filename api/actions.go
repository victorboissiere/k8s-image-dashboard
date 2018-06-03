package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"errors"
	"k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"time"
)

const RollingUpdateEnvKey = "LAST_FORCED_ROLLING_UPDATE"

func TriggerArbitraryRollingUpdateOnDeployment(namespace string, podName string) {
	pod, err := client.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	checkError(err)

	replicatSetName, err := getFirstParentName(pod.OwnerReferences, "ReplicaSet")
	checkError(err)

	replicatSet, err := client.AppsV1().ReplicaSets(namespace).Get(replicatSetName, metav1.GetOptions{})
	checkError(err)

	deploymentName, err := getFirstParentName(replicatSet.OwnerReferences, "Deployment")
	checkError(err)

	deployment, err := client.AppsV1().Deployments(namespace).Get(deploymentName, metav1.GetOptions{})
	checkError(err)

	checkError(updateOrCreateEnv(deployment, RollingUpdateEnvKey, time.Now().String()))
}

func updateOrCreateEnv(deployment *v1.Deployment, key string, value string) error {
	containers := deployment.Spec.Template.Spec.Containers;
	if len(containers) == 0 {
		return errors.New("deployment has no containers specs")
	}

	container := containers[0]

	envPos := findEnvPos(container.Env, key)
	if envPos == -1 {
		deployment.Spec.Template.Spec.Containers[0].Env = append(container.Env, v12.EnvVar{
			Name: key,
			Value:value,
		})
	} else {
		container.Env[envPos].Value = value
	}

	_, err := client.AppsV1().Deployments(deployment.Namespace).Update(deployment)

	return err
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


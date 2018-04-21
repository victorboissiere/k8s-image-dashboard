# K8S Image Dashboard

Simple dashboard to list all pods and associated images to quickly check
which version of your images you are using.

This project aims to quickly check the state of your deployed images.

## Configuration

**Available environment variables**

(optional) `EXCLUDE_NAMESPACES`: Do not show pods in the selected namespaces

_Example_: `EXCLUDE_NAMESPACES:kube-system,kube-public`

## K8S deployment

You can pull the latest docker image at `gitcommit/k8s-image-dashboard`.

Then execute the following command:

```bash
kubectl create -f https://raw.githubusercontent.com/victorboissiere/k8s-image-dashboard/master/k8s/deployment.yaml
kubectl create -f https://raw.githubusercontent.com/victorboissiere/k8s-image-dashboard/master/k8s/service.yaml
```


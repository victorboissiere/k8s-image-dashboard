# K8S Image Dashboard

Simple dashboard to list all pods and associated images to quickly check
which version of your images you are using.

This project aims to quickly check the state of your deployed images.

![Overview](https://raw.githubusercontent.com/victorboissiere/k8s-image-dashboard/master/docs/overview.png)

## Configuration

**Available environment variables**

(optional) `PORT`: Running application port. Default: 3000

(optional) `EXCLUDE_NAMESPACES`: Do not show pods in the selected namespaces

_Example_: `EXCLUDE_NAMESPACES:kube-system,kube-public`

(optional) `REPO_REGEX`: Link to your own repository

Add a link to a specific image by using a matching regex.

_Example_: `REPO_REGEX=gitcommit/([a-zA-Z-_]+)|https://github.com/victorboissiere/$1`

## K8S deployment

You can pull the most recent docker image at `gitcommit/k8s-image-dashboard:v1.2.0`.

Or you can directly execute the following command:

```bash
kubectl create -f https://raw.githubusercontent.com/victorboissiere/k8s-image-dashboard/master/k8s/deployment.yaml
kubectl create -f https://raw.githubusercontent.com/victorboissiere/k8s-image-dashboard/master/k8s/service.yaml
```

The default service type is ClusterIP. Feel free to update
the yaml configuration files to your needs.


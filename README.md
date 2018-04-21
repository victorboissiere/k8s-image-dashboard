# K8S Image Dashboard

Simple dashboard to list all pods and associated images to quickly check
which version of your images you are using.

This project aims to quickly check the state of your deployed images.

![Overview](https://lh3.googleusercontent.com/7BZco6QmQvklkAWa68x-se5yQkjHco5LZEgtvo7CIyEDLh0W27G3vPkNtZdJ5an4sh8Xky26QlLnWve6T5RrPTtu97WdvcOmS6ml974SMgWpnbxjEsQp8dUME3yTLuRXSWf5flZ_oTT7hAphd0d2SOhhyA9X_FNWAn20Hq3K3EGJBvXjCsT1lxwNWwnmMGD2J1Slsq4-j8RrM2mWtxei_lDecx8oT2ZK60FnKavMvl8WzHZIZCrF1Ivylna9MyPyxGn78eqJYe_UyIg_LCfV7lApUlQCJ9a9GldAEziM184LZIQovocNyw8PQyCuxwfI80YXHHtOyJqh0zG0VJrblkkk34sw5wW4hZbIwHn6_2qAWCEMywQE26C0QSepQFWeda-XKx8WEkqb-I5u4d7lxyqUASxGbV3tWyMcSa9_lHswYLrMfjmPqdH2eGIUVFLeJvFwJ5zV0nskW77Q3I7ej8KwkcoEYUHtBsdun6tsuPsINMdpXDl06vUh89wdMq85MP4XdcIOhe6ilEoiym54R9E4dr1k4mY_T6h5RwuawqohcH8rHfRpIsuBiCktGpxVI4TDe6lzAXsM9y8Qu95YOm9oD_cmc5hR2XcZdWQt=w2680-h1548-no)

## Configuration

**Available environment variables**

(optional) `PORT`: Running application port. Default: 3000

(optional) `EXCLUDE_NAMESPACES`: Do not show pods in the selected namespaces

_Example_: `EXCLUDE_NAMESPACES:kube-system,kube-public`

(optional) `REPO_REGEX`: Link to your own repository

Add a link to a specific image by using a matching regex.

_Example_: `REPO_REGEX=gitcommit/([a-zA-Z-_]+)|https://github.com/victorboissiere/$1` 

## K8S deployment

You can pull the most recent docker image at `gitcommit/k8s-image-dashboard:v1.0.1`.

Or you can directly execute the following command:

```bash
kubectl create -f https://raw.githubusercontent.com/victorboissiere/k8s-image-dashboard/master/k8s/deployment.yaml
kubectl create -f https://raw.githubusercontent.com/victorboissiere/k8s-image-dashboard/master/k8s/service.yaml
```

The default service type is ClusterIP. Feel free to update
the yaml configuration files to your needs.


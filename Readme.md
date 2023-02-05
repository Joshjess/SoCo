# How to use the helm chart

### Generating the Chart
```helm create webapp```
### Installing the Chart
```helm install -f .kubernetes/values.yaml webapp ./kubernetes```
#### Uninstalling the Chart
```helm uninstall -f .kubernetes/values.yaml webapp ./kubernetes```
### Upgrading deployment the Chart
```helm upgrade  -f .kubernetes/values.yaml webapp ./kubernetes```


# How to build and push docker images

### Backend
```docker build ./backend -t <backend-tag>```

```docker push <backend-tag>```

### Frontend
```docker build ./frontend -t <frontend-tag>```

```docker push <frontend-tag>```

# Misc commands

```gcloud compute disks create <disk-name> --zone=<google-cloud-region> --size=<disk-size>```

[Link to Github](https://github.com/Joshjess/SoCo)
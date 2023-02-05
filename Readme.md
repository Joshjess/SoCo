# How to use the helm chart

### Generating the Chart
```helm create webapp```
### Installing with the Chart
```helm install -f ./values.yaml webapp .```
#### Uninstalling with the Chart
```helm uninstall -f ./values.yaml webapp .```
### Upgrading an image with the Chart
```helm upgrade -f ./values.yaml webapp .```
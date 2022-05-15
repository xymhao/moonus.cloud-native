### Install istio

```sh
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.12.1 sh -
cd istio-1.12.1
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y
```

### Istio monitoring

- Grafana dashboard `7639`

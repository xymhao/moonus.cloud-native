### install istio follow by module12
```
cd /root/istio-1.12.0
```
### install bookinfo
```
kubectl label namespace default istio-injection=enabled
kubectl apply -f <(istioctl kube-inject -f samples/bookinfo/platform/kube/bookinfo.yaml)
```

### expose the application via istio ingress gateway
```
kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml
```

### 查看istio ingress的端口,80端口为：30690
k get svc -n istio-system
```
istio-ingressgateway   LoadBalancer   10.96.164.31     <pending>     15021:30770/TCP,80:30690/TCP,443:32243/TCP,31400:32155/TCP,15443:31765/TCP   6d22h
```

### check ingress vip
```
export GATEWAY_URL=192.168.34.2:31783
```


### access the productpage via
```
http://192.168.34.2:31783/productpage
```


### 创建ns
```
kubectl create ns module12
```

### 添加istio label
```
kubectl label ns module12 istio-injection=enabled
```

### 创建httpserver
```
kubectl apply -f httpserver-istio.yaml -n module12
```

### 生成证书
```
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.io' -keyout cncamp.io.key -out cncamp.io.crt

kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt

```

### 创建 gateway 和 VirtualService
```
kubectl apply -f istio.yaml -n module12
```

### Access the httpserver via ingress

```sh
k get svc -nistio-system

export INGRESS_IP=10.96.164.31

curl --resolve httpsserver-istio.cncamp.io:443:$INGRESS_IP https://httpsserver-istio.cncamp.io/healthz -v -k
curl --resolve httpsserver-istio.cncamp.io:443:$INGRESS_IP https://httpsserver-istio.cncamp.io/metrics -v -k
curl --resolve httpsserver-istio.cncamp.io:443:$INGRESS_IP https://httpsserver-istio.cncamp.io/httpserver -v -k

curl -H "Host:httpsserver-istio.cncamp.io" $INGRESS_IP/hello -v
```

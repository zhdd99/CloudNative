## istio + tls + L7
```
kubectl create ns homework12
kubectl create -f httpserver-deployment.yaml -n homework12
kubectl create -f nginx.yaml -n homework12
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.io' -keyout cncamp.io.key -out cncamp.io.crt
kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt
kubectl create -f istio-httpserver-specs.yaml -n homework12
curl --resolve httpsserver.cncamp.io:443:10.110.5.139 https://httpsserver.cncamp.io/hello -v -k
```

## tracing 
```
直接使用老师的yaml跑了一遍
```


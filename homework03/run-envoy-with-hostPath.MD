## run envoy
```
$ kubectl create -f envoy-deploy.yaml
$ kubectl expose deploy envoy --selector run=envoy --port=10000 --type=NodePort
$ kubectl exec -it envoy-6d77fdbf47-2czjr bash
```
## access service
```
$ curl <NODE IP Address>:<NodePort>
Notices: Node IP Address base on your kubernetes cluster.

Example:
$ curl 192.168.34.2:<nodeport>
```
## scale up/down/failover
```
$ kubectl scale deploy <deployment-name> --replicas=<n>
```
## cascade delete
```
$ kubectl delete deploy envoy --cascade=orphan
$ kubectl delete rs envoy-6d77fdbf47 --cascade=orphan
$ kubectl delete po envoy-6d77fdbf47-2czjr envoy-6d77fdbf47-6jw7f envoy-6d77fdbf47-zsmsv
```
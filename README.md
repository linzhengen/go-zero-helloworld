# go-zero-helloworld

## Usage:
```
# setup
make deploy-etcd
make dev

# port forward
kubectl -n hello-world-api port-forward services/api-svc 8888:8888
Forwarding from 127.0.0.1:8888 -> 8888
Forwarding from [::1]:8888 -> 8888
Handling connection for 8888

# curl api
curl http://localhost:8888/sayHello
{"name":"test"}%   
```
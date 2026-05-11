# k8s-secret-header-auth

A minimal Go HTTP server that validates request headers against a value
sourced from a Kubernetes Secret.

Built to understand the service-side of KubeVela's `headerFromSecret` pattern,
where a workflow step sources a secret value and injects it as an HTTP request header.

## How it works

A Kubernetes Secret holds an API key. The Deployment injects it into the
container as an environment variable via `secretKeyRef`. The server reads
that value at startup and checks every incoming request for a matching
`X-Api-Key` header.

## Stack

- Go: HTTP server
- Docker: multi-stage build (builder + alpine final image)
- Kubernetes: Secret, Deployment (2 replicas), Service
- kind: local cluster

## Run it

```bash
kind create cluster --name keyed-demo
docker build -t header-validator:latest .
kind load docker-image header-validator:latest --name keyed-demo

kubectl apply -f k8s/secret.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

kubectl port-forward svc/header-validator-svc 8080:80
```

## Test it

```bash
curl -H "X-Api-Key: supersecret" http://localhost:8080   # 200 Authorized
curl -H "X-Api-Key: wrongkey"    http://localhost:8080   # 401 Unauthorized
curl                             http://localhost:8080   # 401 Unauthorized
```
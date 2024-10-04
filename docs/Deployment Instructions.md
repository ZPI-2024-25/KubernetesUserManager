# Local deployment with minikube

Instruction for local deployment with minikube

## Building and deploying for the first time
1. Start minikube

```
minikube start
```

Terminal should have administrator access

2. Point your terminal to minikube's Docker using minikube docker-env

```
# on Mac or Linux
eval $(minikube docker-env) 
```

```
# on Windows PowerShell
& minikube docker-env | Invoke-Expression
```

If you close your terminal, you will have to repeat step 2.

Running `minikube status` can be used to verify if your terminal is pointing to minikube's Docer. It should look like this. Last line should say `docker-env: in-use`

```
PS C:\Users\admin> minikube status
minikube
type: Control Plane
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured
docker-env: in-use
```

3. Build Docker image inside minikube
```
docker build -t client-go-test:v1 .
```

Used to skip uploading image to repository. **Remember to be in correct directory!**

4. Deploy to Kubernetes
```
kubectl apply -f backend/deploy/deployment.yaml
kubectl apply -f backend/deploy/service.yaml
```

5. Accees the deployed service in browser
```
minikube service client-go-test-svc
```

You will receive ip addres that can be used to access service

`$ minikube service list` lists all exposed serivces

## Redeploying after code change

1. Ensure you're in correct directory and still pointing to minikube's Docker
2. Delete the deployment and the image

```
kubectl delete -f backend/deploy/deployment.yaml
docker rmi client-go-test:v1
```

3. Rebuild the image and re-deploy to minikube
```
docker build -t client-go-test:v1 .
kubectl apply -f backend/deploy/deployment.yaml
```

4. Access service in browser
```
minikube service client-go-test-svc
```

## Troubleshooting

1. Insufficient permissions

Error example:
```
Error: pods is forbidden: User "system:serviceaccount:default:default" cannot list resource "pods" in ...
```

With RBAC enabled, use this to create role binding which will grant the default service account view permissions.
```
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
```
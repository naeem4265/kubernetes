# Running the API Server from In Cluster

Follow these steps to run the `api-server` within a Kubernetes cluster:

## Step 1: Start a Cluster with a Config File
    $ kind create cluster --config=./kind-config.yaml

## Step 2: Deploy Ingress Controller (nginx)
    $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

## Step 3: Create Role and Role Binding, Create role and role-binding YAML files, then apply them
     $ kubectl apply -f cluster-role.yaml 
     $ kubectl apply -f cluster-role-binding.yaml
## Step 4: Write Go Code and Build Docker Image
Write your Go code, and then navigate to the directory containing the code. Create a Dockerfile for your Go application.
    $ docker build -t naeem4265/api-server:1.0.5 .
## Step 5: Load Docker Image into Kind Cluster
Load your Docker image into the Kind cluster.
$ kind load docker-image naeem4265/api-server:1.0.5
## Step 6: Deploy the API Server Pod
Write a pod YAML file and apply it to create the API server pod.
    $ kubectl apply -f pod.yaml
## Step 7: Port Forwarding
Port forward to access the API server from outside the cluster.
$ kubectl port-forward api-server-pod 8080
## Step 8: Test with Postman
Test your API server using Postman or a similar tool, using the specific URL.

Now you can access and test your api-server running within the Kubernetes cluster. Have fun coding!

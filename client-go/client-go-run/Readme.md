* Create cluster:
    $ kind create cluster --config=./kind-config.yaml
* Run Ingress Controller(nginx)
    $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
* Run Role & Rolebinding if necessary. 
    $ kubectl apply -f cluster-role.yaml
    $ kubectl apply -f cluster-role-binding.yaml
* Run Ingress file which contail all pods
    & kubectl apply -f ingress.yaml
* Server is ready. check from browser or curl
    $ curl localhost/hello/check

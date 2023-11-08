# Define the target to apply the YAML file
apply:
	 kubectl apply -f yaml/my-configmap.yaml
	 kubectl apply -f yaml/my-secret.yaml
	 kubectl apply -f yaml/pv.yaml
	 kubectl apply -f yaml/pvc.yaml
	 kubectl apply -f my-deployment.yaml
	 kubectl apply -f service.yaml
# Define a target to delete the applied resources
clean:
	 kubectl delete -f my-deployment.yaml
	 kubectl delete -f service.yaml
	 kubectl delete -f yaml/pvc.yaml
	 kubectl delete -f yaml/pv.yaml
	 kubectl delete -f yaml/my-configmap.yaml
	 kubectl delete -f yaml/my-secret.yaml

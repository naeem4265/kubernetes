## Resouces: 
Video: [Volume, pv, pvc, storage class ](https://www.youtube.com/watch?v=LPy6Q-q1MVQ)

Blog: [Multiple values pass as arguments](https://www.baeldung.com/linux/kubernetes-pass-many-commands)
## Steps
In this api-server I created a pod named pod.yaml for access with passing multiple args. 
First passed a signin request and received a cookie that stored in /tmp/cookie.txt file,
then post request with this txt file so that I can insert book. For access api-server from 
different pod/envirnment I must need service as NodePort for external service.

## Run 
    yaml/my-configmap.yaml
    yaml/my-secret.yaml
    yaml/pv.yaml
    yaml/pvc.yaml
    my-deployment.yaml
    service.yaml
    yaml/pod.yaml
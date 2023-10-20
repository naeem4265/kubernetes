# CRD Controller
- Sample Controller - [Link](https://github.com/kubernetes/sample-controller)
- Guidelines for writing controllers - [Link](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-api-machinery/controllers.md)
- K8s Custom Controllers - [Link1](https://www.linkedin.com/pulse/kubernetes-custom-controllers-part-1-kritik-sachdeva/) [Link2](https://www.linkedin.com/pulse/kubernetes-custom-controller-part-2-kritik-sachdeva/)
- K8s Custom Controllers best - [Link](https://itnext.io/how-to-generate-client-codes-for-kubernetes-custom-resource-definitions-crd-b4b9907769ba)
# Steps
- Write doc.go, types.go and register.go in pkg/apis/group/version directory. 
- Install code generator for first time. 
  - $ go get k8s.io/code-generator && go get k8s.io/apimachinery
- Generate Client, Informer, Lister and Deepcopy method from project directory 
  - $ ./generate-groups.sh all "github.com/../.../pkg/client" "github.com/.../.../pkg/apis" groupname:version
- Install controller gen for first time from project directory.
  - $ go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.5.0
  - $ go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.5.0
- Create manifests file from project directory.
  - controller-gen rbac:roleName=controller-perms crd paths=./... output:crd:dir=./manifests output:stdout
- Write main.go and controller.go files for custom controller. 
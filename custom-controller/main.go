package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"time"

	"github.com/naeem4265/custom-controller/controller"
	cClientset "github.com/naeem4265/custom-controller/pkg/client/clientset/versioned"
	cInformer "github.com/naeem4265/custom-controller/pkg/client/informers/externalversions"
	informer "k8s.io/client-go/informers"
)

func main() {
	// Authentication
	kubeconfig := flag.String("kubeconfig", "/home/appscodepc/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("erorr %s building config from flags\n", err.Error())
		// If in cluster
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s, getting inclusterconfig", err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// handle error
		fmt.Printf("error %s, creating clientset\n", err.Error())
	}
	bookClientset, err := cClientset.NewForConfig(config)
	if err != nil {
		// handle error
		fmt.Printf("error %s, creating custom clientset\n", err.Error())
	}

	informer := informer.NewSharedInformerFactory(clientset, 10*time.Minute)
	bookInformer := cInformer.NewSharedInformerFactory(bookClientset, 10*time.Minute)

	ctrl := controller.NewController(clientset, bookClientset,
		informer.Apps().V1().Deployments(), bookInformer.Naeem4265().V1().Books())

	ch := make(chan struct{})

	bookInformer.Start(ch)
	if err := ctrl.Run(ch); err != nil {
		fmt.Printf("%s Error running controller\n", err.Error())
	}
}

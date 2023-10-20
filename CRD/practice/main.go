package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {
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

	informer := informers.NewSharedInformerFactory(clientset, 10*time.Minute)

	c := newController(clientset, informer.Apps().V1().Deployments())
	ch := make(chan struct{})
	informer.Start(ch)
	c.run(ch)
	fmt.Println(informer)
}

// More details: https://www.youtube.com/watch?v=lzoWSfvE2yA

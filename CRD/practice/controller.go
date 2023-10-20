package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"time"
)

type controller struct {
	clientset      kubernetes.Interface
	depLister      appslisters.DeploymentLister
	depCacheSynced cache.InformerSynced
	queue          workqueue.RateLimitingInterface
}

func newController(clientset kubernetes.Interface, depInformer appsinformers.DeploymentInformer) *controller {
	c := &controller{
		clientset:      clientset,
		depLister:      depInformer.Lister(),
		depCacheSynced: depInformer.Informer().HasSynced,
		queue:          workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "my-queue"),
	}
	depInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Println("added")
				c.queue.Add(obj)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Println("updated")
				c.queue.Add(newObj)
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Println("deleted")
				c.queue.Add(obj)
			},
		})
	return c
}

func (c *controller) run(ch <-chan struct{}) {
	fmt.Println("starting controller")
	if !cache.WaitForCacheSync(ch, c.depCacheSynced) {
		fmt.Printf("Waiting for cache to be synced\n")
	}
	go wait.Until(c.worker, 1*time.Second, ch)
	<-ch
}

func (c *controller) worker() {
	for c.processItem() {

	}
}

func (c *controller) processItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}

	defer c.queue.Forget(item)
	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		fmt.Printf("getting key form cache %s \n", err.Error())
		return false
	}

	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		fmt.Printf("Namespace error %s\n", err.Error())
		return false
	}

	err = c.syncDeployment(ns, name)
	if err != nil {
		fmt.Println("Syncing deployment %s\n", err.Error())
		return false
	}
	return true
}

func (c *controller) syncDeployment(ns, name string) error {
	dep, err := c.depLister.Deployments(ns).Get(name)
	svc := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dep.Name,
			Namespace: ns,
		},
		Spec: corev1.ServiceSpec{
			Selector: depLeabel(*dep),
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					Port: 8080,
				},
			},
		},
	}
	_, err = c.clientset.CoreV1().Services(ns).Create(context.TODO(), &svc, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Creating service %s\n", err.Error())
	}
	fmt.Printf("Service created successfully\n")
	return nil
}

func depLeabel(dep appsv1.Deployment) map[string]string {
	return dep.Spec.Template.Labels
}

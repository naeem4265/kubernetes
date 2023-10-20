package controller

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/intstr"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"log"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	controllerv1 "github.com/naeem4265/custom-controller/pkg/apis/naeem4265.com/v1"

	cClientset "github.com/naeem4265/custom-controller/pkg/client/clientset/versioned"

	cInformer "github.com/naeem4265/custom-controller/pkg/client/informers/externalversions/naeem4265.com/v1"
	appsInformers "k8s.io/client-go/informers/apps/v1"

	cListers "github.com/naeem4265/custom-controller/pkg/client/listers/naeem4265.com/v1"
	appsListers "k8s.io/client-go/listers/apps/v1"
)

type controller struct {
	// CustomClientset is a clientset for our own API group
	BookClientset cClientset.Interface
	Clientset     kubernetes.Interface

	DeploymentLister appsListers.DeploymentLister
	BookLister       cListers.BookLister

	DeploymentSynced cache.InformerSynced
	BookSynced       cache.InformerSynced

	WorkQueue workqueue.RateLimitingInterface
}

func NewController(clientset kubernetes.Interface, bookClientset cClientset.Interface,
	depInformer appsInformers.DeploymentInformer, bookInformer cInformer.BookInformer,
) *controller {

	ctrl := &controller{
		Clientset:     clientset,
		BookClientset: bookClientset,

		DeploymentLister: depInformer.Lister(),
		BookLister:       bookInformer.Lister(),

		DeploymentSynced: depInformer.Informer().HasSynced,
		BookSynced:       bookInformer.Informer().HasSynced,
		WorkQueue:        workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "my-queue"),
	}
	log.Println("Setting up event handlers")

	// Set up an event handler for when Book resources change
	bookInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Println("Custom resource added")
				ctrl.enqueueObj(obj)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Println("Custom resource updated")
				ctrl.enqueueObj(newObj)
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Println("Custom resource deleted")
				ctrl.enqueueObj(obj)
			},
		})
	return ctrl
}

// enqueueObj takes an Book resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than Book.
func (ctrl *controller) enqueueObj(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	ctrl.WorkQueue.AddRateLimited(key)
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shut down the workqueue and wait for
// workers to finish processing their current work items.
func (ctrl *controller) Run(ch <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer ctrl.WorkQueue.ShuttingDown()

	fmt.Println("starting controller")

	if ok := cache.WaitForCacheSync(ch, ctrl.DeploymentSynced, ctrl.BookSynced); !ok {
		return fmt.Errorf("Failed to wait for cache to sync")
	}

	go wait.Until(ctrl.worker, 1*time.Second, ch)
	log.Println("Worker Started")
	<-ch
	log.Println("Shutting down workers")

	return nil
}

// Worker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the workqueue.
func (ctrl *controller) worker() {
	for ctrl.processNextItem() {
	}
}

func (ctrl *controller) processNextItem() bool {
	item, shutdown := ctrl.WorkQueue.Get()
	if shutdown {
		return false
	}

	defer ctrl.WorkQueue.Forget(item)
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

	err = ctrl.syncHandler(ns, name)
	if err != nil {
		fmt.Println("Syncing deployment %s\n", err.Error())
		return false
	}
	return true
}

func (ctrl *controller) syncHandler(ns, name string) error {
	// Get the Book resource with this namespace/name
	resource, err := ctrl.BookLister.Books(ns).Get(name)
	if err != nil {
		fmt.Printf("Book type resource %s\n", err.Error())
		return nil
	}

	depName := resource.Spec.Name
	if depName == "" {
		// We choose to absorb the error here as the worker would requeue the
		// resource otherwise. Instead, the next time the resource is updated
		// the resource will be queued again.
		utilruntime.HandleError(fmt.Errorf("%s : deployment name must be specified", resource))
		return nil
	}

	// Get the deployment with the name specified in Book.spec
	// If the resource doesn't exist, we'll create it
	// If an error occurs during Get/Create, we'll requeue the item, so we can
	deployment, err := ctrl.DeploymentLister.Deployments(resource.Namespace).Get(depName)
	if errors.IsNotFound(err) {
		deployment, err = ctrl.Clientset.AppsV1().Deployments(resource.Namespace).Create(context.TODO(), newDeployment(resource), metav1.CreateOptions{})
	}
	if err != nil {
		return err
	}

	// Finally, we update the status block of the Book resource to reflect the current state of the world
	err = ctrl.updateStatus(resource, deployment)
	if err != nil {
		return err
	}

	// Get the deployment with the name specified in Book.spec
	// If the resource doesn't exist, we'll create it
	// If an error occurs during Get/Create, we'll requeue the item, so we can
	servicename := resource.Spec.Name + "-service"
	service, err := ctrl.Clientset.CoreV1().Services(resource.Namespace).Get(context.TODO(), servicename, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		service, err = ctrl.Clientset.CoreV1().Services(resource.Namespace).Create(context.TODO(), customService(resource), metav1.CreateOptions{})
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("\nservice %s created .....\n", service.Name)
	} else if err != nil {
		log.Println(err)
		return err
	}

	// For confirmation
	_, err = ctrl.Clientset.CoreV1().Services(resource.Namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ctrl *controller) updateStatus(resouce *controllerv1.Book, deployment *appsv1.Deployment) error {

	// NEVER modify objects from the store. It's a read-only, local cache.
	// You can use DeepCopy() to make a deep copy of original object and modify this copy
	// Or create a copy manually for better performance
	Copy := resouce.DeepCopy()
	Copy.Status.AvailableReplicas = deployment.Status.AvailableReplicas

	// If the CustomResourceSubresources feature gate is not enabled,
	// we must use Update instead of UpdateStatus to update the Status block of the Foo resource.
	// UpdateStatus will not allow changes to the Spec of the resource,
	// which is ideal for ensuring nothing other than resource status has been updated.
	_, err := ctrl.BookClientset.Naeem4265V1().Books(resouce.Namespace).Update(context.TODO(), Copy, metav1.UpdateOptions{})

	return err

}

// newDeployment creates a new Deployment for a Book resource. It also sets
// the appropriate OwnerReferences on the resource so handleObject can discover
// the Book resource that 'owns' it.
func newDeployment(resource *controllerv1.Book) *appsv1.Deployment {

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resource.Spec.Name,
			Namespace: resource.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: resource.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "my-app",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "my-app",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "my-app-pod",
							Image: resource.Spec.Container.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: resource.Spec.Container.Port,
								},
							},
						},
					},
				},
			},
		},
	}
}

// customService creates a new service for a custom resource
func customService(resource *controllerv1.Book) *corev1.Service {
	labels := map[string]string{
		"app": "my-app",
	}
	service := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: resource.Spec.Name + "-service",
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(resource, controllerv1.SchemeGroupVersion.WithKind("Book")),
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       resource.Spec.Container.Port,
					TargetPort: intstr.FromInt32(resource.Spec.Container.Port),
					NodePort:   30007,
				},
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}
	return service
}

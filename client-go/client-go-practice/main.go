package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/klog/v2"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "Path")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "Path")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	klog.Infof("Pod creating...")
	err = CreatePod(clientset.CoreV1().Pods("default"))
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	time.Sleep(10 * time.Minute)
}

func CreatePod(pc v1.PodInterface) error {
	pod := corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "api-server-pod",
			Labels: map[string]string{
				"app": "api-server-pod",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:            "api-server-pod",
					Image:           "naeem4265/api-server:1.0.4",
					Command:         nil,
					Args:            nil,
					ImagePullPolicy: corev1.PullIfNotPresent,
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 8080,
						},
					},
				},
			},
		},
		Status: corev1.PodStatus{},
	}
	_, err := pc.Get(context.TODO(), "api-server-pod", metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			p, err := pc.Create(context.TODO(), &pod, metav1.CreateOptions{})
			if err != nil {
				return err
			}
			klog.Infof("%s pod created !! ", p.Name)
			return nil
		} else {
			return err
		}
	}
	return nil
}

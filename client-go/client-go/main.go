/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"time"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	// Creating out cluster config

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	klog.Infof("All good..........................")

	err = CreateAPIServerPod(clientset.CoreV1().Pods("default"))
	if err != nil {
		panic(err.Error())
	}
	er := CreateAPIServerService(clientset.CoreV1().Services("default"))
	if er != nil {
		panic(er.Error())
	}
	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	time.Sleep(10 * time.Minute)
}

func CreateAPIServerPod(pc v1.PodInterface) error {
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

func CreateAPIServerService(sc v1.ServiceInterface) error {
	service := corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "api-server-service",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "api-server-pod",
			},
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       8080,
					TargetPort: intstr.FromInt32(8080),
				},
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}
	_, err := sc.Get(context.TODO(), "api-server-servics", metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			s, err := sc.Create(context.TODO(), &service, metav1.CreateOptions{})
			if err != nil {
				klog.Infof("%s pod created !! ", s.Name)
			}
			klog.Infof("%s Service created!", s.Name)
		} else {
			return err
		}
	}
	return nil
}

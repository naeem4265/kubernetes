// +groupName=naeem4265.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ContainerSpec struct {
	Image string `json:"image,omitempty"`
	Port  int32  `json:"port,omitempty"`
}

// BookSpec defines the desired state of Book CRD
type BookSpec struct {
	Name      string        `json:"name,omitempty"`
	Replicas  *int32        `json:"replicas"`
	Container ContainerSpec `json:"container,container"`
}

// BookStatus defines the observed state of Book CRD
type BookStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Book is the Schema for the book API
type Book struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Custom fields for my resource
	Spec   BookSpec   `json:"spec,omitempty"`
	Status BookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BookList contains a list of Book
type BookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Book `json:"items"`
}

/*
Copyright Ishtiaq Islam.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/ishtiaqhimel/crd-controller/pkg/apis/crd.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IshtiaqLister helps list Ishtiaqs.
// All objects returned here must be treated as read-only.
type IshtiaqLister interface {
	// List lists all Ishtiaqs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Ishtiaq, err error)
	// Ishtiaqs returns an object that can list and get Ishtiaqs.
	Ishtiaqs(namespace string) IshtiaqNamespaceLister
	IshtiaqListerExpansion
}

// ishtiaqLister implements the IshtiaqLister interface.
type ishtiaqLister struct {
	indexer cache.Indexer
}

// NewIshtiaqLister returns a new IshtiaqLister.
func NewIshtiaqLister(indexer cache.Indexer) IshtiaqLister {
	return &ishtiaqLister{indexer: indexer}
}

// List lists all Ishtiaqs in the indexer.
func (s *ishtiaqLister) List(selector labels.Selector) (ret []*v1.Ishtiaq, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Ishtiaq))
	})
	return ret, err
}

// Ishtiaqs returns an object that can list and get Ishtiaqs.
func (s *ishtiaqLister) Ishtiaqs(namespace string) IshtiaqNamespaceLister {
	return ishtiaqNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IshtiaqNamespaceLister helps list and get Ishtiaqs.
// All objects returned here must be treated as read-only.
type IshtiaqNamespaceLister interface {
	// List lists all Ishtiaqs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Ishtiaq, err error)
	// Get retrieves the Ishtiaq from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Ishtiaq, error)
	IshtiaqNamespaceListerExpansion
}

// ishtiaqNamespaceLister implements the IshtiaqNamespaceLister
// interface.
type ishtiaqNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ishtiaqs in the indexer for a given namespace.
func (s ishtiaqNamespaceLister) List(selector labels.Selector) (ret []*v1.Ishtiaq, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Ishtiaq))
	})
	return ret, err
}

// Get retrieves the Ishtiaq from the indexer for a given namespace and name.
func (s ishtiaqNamespaceLister) Get(name string) (*v1.Ishtiaq, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("ishtiaq"), name)
	}
	return obj.(*v1.Ishtiaq), nil
}
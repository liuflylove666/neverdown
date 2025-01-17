/*
Copyright The Kubernetes Authors.

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
	v1 "github.com/kzz45/neverdown/pkg/apis/openx/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OpenxLister helps list Openxes.
// All objects returned here must be treated as read-only.
type OpenxLister interface {
	// List lists all Openxes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Openx, err error)
	// Openxes returns an object that can list and get Openxes.
	Openxes(namespace string) OpenxNamespaceLister
	OpenxListerExpansion
}

// openxLister implements the OpenxLister interface.
type openxLister struct {
	indexer cache.Indexer
}

// NewOpenxLister returns a new OpenxLister.
func NewOpenxLister(indexer cache.Indexer) OpenxLister {
	return &openxLister{indexer: indexer}
}

// List lists all Openxes in the indexer.
func (s *openxLister) List(selector labels.Selector) (ret []*v1.Openx, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Openx))
	})
	return ret, err
}

// Openxes returns an object that can list and get Openxes.
func (s *openxLister) Openxes(namespace string) OpenxNamespaceLister {
	return openxNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OpenxNamespaceLister helps list and get Openxes.
// All objects returned here must be treated as read-only.
type OpenxNamespaceLister interface {
	// List lists all Openxes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Openx, err error)
	// Get retrieves the Openx from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Openx, error)
	OpenxNamespaceListerExpansion
}

// openxNamespaceLister implements the OpenxNamespaceLister
// interface.
type openxNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Openxes in the indexer for a given namespace.
func (s openxNamespaceLister) List(selector labels.Selector) (ret []*v1.Openx, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Openx))
	})
	return ret, err
}

// Get retrieves the Openx from the indexer for a given namespace and name.
func (s openxNamespaceLister) Get(name string) (*v1.Openx, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("openx"), name)
	}
	return obj.(*v1.Openx), nil
}

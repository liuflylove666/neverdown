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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "github.com/kzz45/neverdown/pkg/apis/openx/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=openx, Version=v1
	case v1.SchemeGroupVersion.WithResource("accesscontrols"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openx().V1().AccessControls().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("affinities"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openx().V1().Affinities().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("etcds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openx().V1().Etcds().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("loadbalancers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openx().V1().LoadBalancers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("mysqls"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openx().V1().Mysqls().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("nodeselectors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openx().V1().NodeSelectors().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("openxes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openx().V1().Openxes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("redises"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openx().V1().Redises().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tolerations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openx().V1().Tolerations().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

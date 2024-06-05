/*
Copyright The Discovery Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package kubernetes

import (

	auditv1 "github.com/kzz45/discovery/pkg/client-go/kubernetes/typed/audit/v1"
	jingxv1 "github.com/kzz45/discovery/pkg/client-go/kubernetes/typed/jingx/v1"
	rbacv1 "github.com/kzz45/discovery/pkg/client-go/kubernetes/typed/rbac/v1"
	discovery "github.com/kzz45/discovery/pkg/client-go/discovery"
	rest "github.com/kzz45/discovery/pkg/client-go/rest"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AuditV1() auditv1.AuditV1Interface
	JingxV1() jingxv1.JingxV1Interface
	RbacV1() rbacv1.RbacV1Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	auditV1 *auditv1.AuditV1Client
	jingxV1 *jingxv1.JingxV1Client
	rbacV1  *rbacv1.RbacV1Client
}

// AuditV1 retrieves the AuditV1Client
func (c *Clientset) AuditV1() auditv1.AuditV1Interface {
	return c.auditV1
}

// JingxV1 retrieves the JingxV1Client
func (c *Clientset) JingxV1() jingxv1.JingxV1Interface {
	return c.jingxV1
}

// RbacV1 retrieves the RbacV1Client
func (c *Clientset) RbacV1() rbacv1.RbacV1Interface {
	return c.rbacV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	var cs Clientset
	var err error
	
	cs.auditV1, err = auditv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.jingxV1, err = jingxv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.rbacV1, err = rbacv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}


// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.auditV1 = auditv1.NewForConfigOrDie(c)
	cs.jingxV1 = jingxv1.NewForConfigOrDie(c)
	cs.rbacV1 = rbacv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.auditV1 = auditv1.New(c)
	cs.jingxV1 = jingxv1.New(c)
	cs.rbacV1 = rbacv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
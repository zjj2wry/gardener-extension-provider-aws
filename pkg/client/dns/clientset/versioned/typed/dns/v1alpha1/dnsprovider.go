/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	scheme "github.com/gardener/external-dns-management/pkg/client/dns/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DNSProvidersGetter has a method to return a DNSProviderInterface.
// A group's client should implement this interface.
type DNSProvidersGetter interface {
	DNSProviders(namespace string) DNSProviderInterface
}

// DNSProviderInterface has methods to work with DNSProvider resources.
type DNSProviderInterface interface {
	Create(ctx context.Context, dNSProvider *v1alpha1.DNSProvider, opts v1.CreateOptions) (*v1alpha1.DNSProvider, error)
	Update(ctx context.Context, dNSProvider *v1alpha1.DNSProvider, opts v1.UpdateOptions) (*v1alpha1.DNSProvider, error)
	UpdateStatus(ctx context.Context, dNSProvider *v1alpha1.DNSProvider, opts v1.UpdateOptions) (*v1alpha1.DNSProvider, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DNSProvider, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DNSProviderList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DNSProvider, err error)
	DNSProviderExpansion
}

// dNSProviders implements DNSProviderInterface
type dNSProviders struct {
	client rest.Interface
	ns     string
}

// newDNSProviders returns a DNSProviders
func newDNSProviders(c *DnsV1alpha1Client, namespace string) *dNSProviders {
	return &dNSProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dNSProvider, and returns the corresponding dNSProvider object, and an error if there is any.
func (c *dNSProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DNSProvider, err error) {
	result = &v1alpha1.DNSProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DNSProviders that match those selectors.
func (c *dNSProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DNSProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DNSProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dNSProviders.
func (c *dNSProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnsproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dNSProvider and creates it.  Returns the server's representation of the dNSProvider, and an error, if there is any.
func (c *dNSProviders) Create(ctx context.Context, dNSProvider *v1alpha1.DNSProvider, opts v1.CreateOptions) (result *v1alpha1.DNSProvider, err error) {
	result = &v1alpha1.DNSProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnsproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNSProvider).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dNSProvider and updates it. Returns the server's representation of the dNSProvider, and an error, if there is any.
func (c *dNSProviders) Update(ctx context.Context, dNSProvider *v1alpha1.DNSProvider, opts v1.UpdateOptions) (result *v1alpha1.DNSProvider, err error) {
	result = &v1alpha1.DNSProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsproviders").
		Name(dNSProvider.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNSProvider).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dNSProviders) UpdateStatus(ctx context.Context, dNSProvider *v1alpha1.DNSProvider, opts v1.UpdateOptions) (result *v1alpha1.DNSProvider, err error) {
	result = &v1alpha1.DNSProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsproviders").
		Name(dNSProvider.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNSProvider).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dNSProvider and deletes it. Returns an error if one occurs.
func (c *dNSProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsproviders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dNSProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsproviders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dNSProvider.
func (c *dNSProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DNSProvider, err error) {
	result = &v1alpha1.DNSProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnsproviders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	scheme "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualMachineSetsGetter has a method to return a VirtualMachineSetInterface.
// A group's client should implement this interface.
type VirtualMachineSetsGetter interface {
	VirtualMachineSets() VirtualMachineSetInterface
}

// VirtualMachineSetInterface has methods to work with VirtualMachineSet resources.
type VirtualMachineSetInterface interface {
	Create(*v1.VirtualMachineSet) (*v1.VirtualMachineSet, error)
	Update(*v1.VirtualMachineSet) (*v1.VirtualMachineSet, error)
	UpdateStatus(*v1.VirtualMachineSet) (*v1.VirtualMachineSet, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.VirtualMachineSet, error)
	List(opts metav1.ListOptions) (*v1.VirtualMachineSetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.VirtualMachineSet, err error)
	VirtualMachineSetExpansion
}

// virtualMachineSets implements VirtualMachineSetInterface
type virtualMachineSets struct {
	client rest.Interface
}

// newVirtualMachineSets returns a VirtualMachineSets
func newVirtualMachineSets(c *HobbyfarmV1Client) *virtualMachineSets {
	return &virtualMachineSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the virtualMachineSet, and returns the corresponding virtualMachineSet object, and an error if there is any.
func (c *virtualMachineSets) Get(name string, options metav1.GetOptions) (result *v1.VirtualMachineSet, err error) {
	result = &v1.VirtualMachineSet{}
	err = c.client.Get().
		Resource("virtualmachinesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineSets that match those selectors.
func (c *virtualMachineSets) List(opts metav1.ListOptions) (result *v1.VirtualMachineSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualMachineSetList{}
	err = c.client.Get().
		Resource("virtualmachinesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineSets.
func (c *virtualMachineSets) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("virtualmachinesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualMachineSet and creates it.  Returns the server's representation of the virtualMachineSet, and an error, if there is any.
func (c *virtualMachineSets) Create(virtualMachineSet *v1.VirtualMachineSet) (result *v1.VirtualMachineSet, err error) {
	result = &v1.VirtualMachineSet{}
	err = c.client.Post().
		Resource("virtualmachinesets").
		Body(virtualMachineSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualMachineSet and updates it. Returns the server's representation of the virtualMachineSet, and an error, if there is any.
func (c *virtualMachineSets) Update(virtualMachineSet *v1.VirtualMachineSet) (result *v1.VirtualMachineSet, err error) {
	result = &v1.VirtualMachineSet{}
	err = c.client.Put().
		Resource("virtualmachinesets").
		Name(virtualMachineSet.Name).
		Body(virtualMachineSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualMachineSets) UpdateStatus(virtualMachineSet *v1.VirtualMachineSet) (result *v1.VirtualMachineSet, err error) {
	result = &v1.VirtualMachineSet{}
	err = c.client.Put().
		Resource("virtualmachinesets").
		Name(virtualMachineSet.Name).
		SubResource("status").
		Body(virtualMachineSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualMachineSet and deletes it. Returns an error if one occurs.
func (c *virtualMachineSets) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("virtualmachinesets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineSets) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("virtualmachinesets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualMachineSet.
func (c *virtualMachineSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.VirtualMachineSet, err error) {
	result = &v1.VirtualMachineSet{}
	err = c.client.Patch(pt).
		Resource("virtualmachinesets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

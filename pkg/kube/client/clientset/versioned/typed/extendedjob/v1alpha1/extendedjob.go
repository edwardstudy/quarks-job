/*

Don't alter this file, it was generated.

*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "code.cloudfoundry.org/quarks-job/pkg/kube/apis/extendedjob/v1alpha1"
	scheme "code.cloudfoundry.org/quarks-job/pkg/kube/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExtendedJobsGetter has a method to return a ExtendedJobInterface.
// A group's client should implement this interface.
type ExtendedJobsGetter interface {
	ExtendedJobs(namespace string) ExtendedJobInterface
}

// ExtendedJobInterface has methods to work with ExtendedJob resources.
type ExtendedJobInterface interface {
	Create(*v1alpha1.ExtendedJob) (*v1alpha1.ExtendedJob, error)
	Update(*v1alpha1.ExtendedJob) (*v1alpha1.ExtendedJob, error)
	UpdateStatus(*v1alpha1.ExtendedJob) (*v1alpha1.ExtendedJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ExtendedJob, error)
	List(opts v1.ListOptions) (*v1alpha1.ExtendedJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ExtendedJob, err error)
	ExtendedJobExpansion
}

// extendedJobs implements ExtendedJobInterface
type extendedJobs struct {
	client rest.Interface
	ns     string
}

// newExtendedJobs returns a ExtendedJobs
func newExtendedJobs(c *ExtendedjobV1alpha1Client, namespace string) *extendedJobs {
	return &extendedJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the extendedJob, and returns the corresponding extendedJob object, and an error if there is any.
func (c *extendedJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.ExtendedJob, err error) {
	result = &v1alpha1.ExtendedJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("extendedjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ExtendedJobs that match those selectors.
func (c *extendedJobs) List(opts v1.ListOptions) (result *v1alpha1.ExtendedJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ExtendedJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("extendedjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested extendedJobs.
func (c *extendedJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("extendedjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a extendedJob and creates it.  Returns the server's representation of the extendedJob, and an error, if there is any.
func (c *extendedJobs) Create(extendedJob *v1alpha1.ExtendedJob) (result *v1alpha1.ExtendedJob, err error) {
	result = &v1alpha1.ExtendedJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("extendedjobs").
		Body(extendedJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a extendedJob and updates it. Returns the server's representation of the extendedJob, and an error, if there is any.
func (c *extendedJobs) Update(extendedJob *v1alpha1.ExtendedJob) (result *v1alpha1.ExtendedJob, err error) {
	result = &v1alpha1.ExtendedJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("extendedjobs").
		Name(extendedJob.Name).
		Body(extendedJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *extendedJobs) UpdateStatus(extendedJob *v1alpha1.ExtendedJob) (result *v1alpha1.ExtendedJob, err error) {
	result = &v1alpha1.ExtendedJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("extendedjobs").
		Name(extendedJob.Name).
		SubResource("status").
		Body(extendedJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the extendedJob and deletes it. Returns an error if one occurs.
func (c *extendedJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("extendedjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *extendedJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("extendedjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched extendedJob.
func (c *extendedJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ExtendedJob, err error) {
	result = &v1alpha1.ExtendedJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("extendedjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

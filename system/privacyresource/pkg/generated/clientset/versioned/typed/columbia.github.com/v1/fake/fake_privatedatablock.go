/*
Copyright (c) 2020 Mingen Pan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	columbiagithubcomv1 "columbia.github.com/sage/privacyresource/pkg/apis/columbia.github.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrivateDataBlocks implements PrivateDataBlockInterface
type FakePrivateDataBlocks struct {
	Fake *FakeColumbiaV1
	ns   string
}

var privatedatablocksResource = schema.GroupVersionResource{Group: "columbia.github.com", Version: "v1", Resource: "privatedatablocks"}

var privatedatablocksKind = schema.GroupVersionKind{Group: "columbia.github.com", Version: "v1", Kind: "PrivateDataBlock"}

// Get takes name of the privateDataBlock, and returns the corresponding privateDataBlock object, and an error if there is any.
func (c *FakePrivateDataBlocks) Get(ctx context.Context, name string, options v1.GetOptions) (result *columbiagithubcomv1.PrivateDataBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatedatablocksResource, c.ns, name), &columbiagithubcomv1.PrivateDataBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*columbiagithubcomv1.PrivateDataBlock), err
}

// List takes label and field selectors, and returns the list of PrivateDataBlocks that match those selectors.
func (c *FakePrivateDataBlocks) List(ctx context.Context, opts v1.ListOptions) (result *columbiagithubcomv1.PrivateDataBlockList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatedatablocksResource, privatedatablocksKind, c.ns, opts), &columbiagithubcomv1.PrivateDataBlockList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &columbiagithubcomv1.PrivateDataBlockList{ListMeta: obj.(*columbiagithubcomv1.PrivateDataBlockList).ListMeta}
	for _, item := range obj.(*columbiagithubcomv1.PrivateDataBlockList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateDataBlocks.
func (c *FakePrivateDataBlocks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatedatablocksResource, c.ns, opts))

}

// Create takes the representation of a privateDataBlock and creates it.  Returns the server's representation of the privateDataBlock, and an error, if there is any.
func (c *FakePrivateDataBlocks) Create(ctx context.Context, privateDataBlock *columbiagithubcomv1.PrivateDataBlock, opts v1.CreateOptions) (result *columbiagithubcomv1.PrivateDataBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatedatablocksResource, c.ns, privateDataBlock), &columbiagithubcomv1.PrivateDataBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*columbiagithubcomv1.PrivateDataBlock), err
}

// Update takes the representation of a privateDataBlock and updates it. Returns the server's representation of the privateDataBlock, and an error, if there is any.
func (c *FakePrivateDataBlocks) Update(ctx context.Context, privateDataBlock *columbiagithubcomv1.PrivateDataBlock, opts v1.UpdateOptions) (result *columbiagithubcomv1.PrivateDataBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatedatablocksResource, c.ns, privateDataBlock), &columbiagithubcomv1.PrivateDataBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*columbiagithubcomv1.PrivateDataBlock), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateDataBlocks) UpdateStatus(ctx context.Context, privateDataBlock *columbiagithubcomv1.PrivateDataBlock, opts v1.UpdateOptions) (*columbiagithubcomv1.PrivateDataBlock, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatedatablocksResource, "status", c.ns, privateDataBlock), &columbiagithubcomv1.PrivateDataBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*columbiagithubcomv1.PrivateDataBlock), err
}

// Delete takes name of the privateDataBlock and deletes it. Returns an error if one occurs.
func (c *FakePrivateDataBlocks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(privatedatablocksResource, c.ns, name), &columbiagithubcomv1.PrivateDataBlock{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateDataBlocks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatedatablocksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &columbiagithubcomv1.PrivateDataBlockList{})
	return err
}

// Patch applies the patch and returns the patched privateDataBlock.
func (c *FakePrivateDataBlocks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *columbiagithubcomv1.PrivateDataBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatedatablocksResource, c.ns, name, pt, data, subresources...), &columbiagithubcomv1.PrivateDataBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*columbiagithubcomv1.PrivateDataBlock), err
}
/*
Copyright 2020 The Knative Authors

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
// Code generated by injection-gen. DO NOT EDIT.

package client

import (
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	discovery "k8s.io/client-go/discovery"
	dynamic "k8s.io/client-go/dynamic"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing-rabbitmq/pkg/apis/sources/v1alpha1"
	versioned "knative.dev/eventing-rabbitmq/pkg/client/clientset/versioned"
	typedsourcesv1alpha1 "knative.dev/eventing-rabbitmq/pkg/client/clientset/versioned/typed/sources/v1alpha1"
	injection "knative.dev/pkg/injection"
	dynamicclient "knative.dev/pkg/injection/clients/dynamicclient"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterClient(withClientFromConfig)
	injection.Default.RegisterClientFetcher(func(ctx context.Context) interface{} {
		return Get(ctx)
	})
	injection.Dynamic.RegisterDynamicClient(withClientFromDynamic)
}

// Key is used as the key for associating information with a context.Context.
type Key struct{}

func withClientFromConfig(ctx context.Context, cfg *rest.Config) context.Context {
	return context.WithValue(ctx, Key{}, versioned.NewForConfigOrDie(cfg))
}

func withClientFromDynamic(ctx context.Context) context.Context {
	return context.WithValue(ctx, Key{}, &wrapClient{dyn: dynamicclient.Get(ctx)})
}

// Get extracts the versioned.Interface client from the context.
func Get(ctx context.Context) versioned.Interface {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		if injection.GetConfig(ctx) == nil {
			logging.FromContext(ctx).Panic(
				"Unable to fetch knative.dev/eventing-rabbitmq/pkg/client/clientset/versioned.Interface from context. This context is not the application context (which is typically given to constructors via sharedmain).")
		} else {
			logging.FromContext(ctx).Panic(
				"Unable to fetch knative.dev/eventing-rabbitmq/pkg/client/clientset/versioned.Interface from context.")
		}
	}
	return untyped.(versioned.Interface)
}

type wrapClient struct {
	dyn dynamic.Interface
}

var _ versioned.Interface = (*wrapClient)(nil)

func (w *wrapClient) Discovery() discovery.DiscoveryInterface {
	panic("Discovery called on dynamic client!")
}

func convert(from interface{}, to runtime.Object) error {
	bs, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("Marshal() = %w", err)
	}
	if err := json.Unmarshal(bs, to); err != nil {
		return fmt.Errorf("Unmarshal() = %w", err)
	}
	return nil
}

// SourcesV1alpha1 retrieves the SourcesV1alpha1Client
func (w *wrapClient) SourcesV1alpha1() typedsourcesv1alpha1.SourcesV1alpha1Interface {
	return &wrapSourcesV1alpha1{
		dyn: w.dyn,
	}
}

type wrapSourcesV1alpha1 struct {
	dyn dynamic.Interface
}

func (w *wrapSourcesV1alpha1) RESTClient() rest.Interface {
	panic("RESTClient called on dynamic client!")
}

func (w *wrapSourcesV1alpha1) RabbitmqSources(namespace string) typedsourcesv1alpha1.RabbitmqSourceInterface {
	return &wrapSourcesV1alpha1RabbitmqSourceImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "sources.knative.dev",
			Version:  "v1alpha1",
			Resource: "rabbitmqsources",
		}),

		namespace: namespace,
	}
}

type wrapSourcesV1alpha1RabbitmqSourceImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedsourcesv1alpha1.RabbitmqSourceInterface = (*wrapSourcesV1alpha1RabbitmqSourceImpl)(nil)

func (w *wrapSourcesV1alpha1RabbitmqSourceImpl) Create(ctx context.Context, in *v1alpha1.RabbitmqSource, opts v1.CreateOptions) (*v1alpha1.RabbitmqSource, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "sources.knative.dev",
		Version: "v1alpha1",
		Kind:    "RabbitmqSource",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.RabbitmqSource{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapSourcesV1alpha1RabbitmqSourceImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapSourcesV1alpha1RabbitmqSourceImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapSourcesV1alpha1RabbitmqSourceImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.RabbitmqSource, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.RabbitmqSource{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapSourcesV1alpha1RabbitmqSourceImpl) List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RabbitmqSourceList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.RabbitmqSourceList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapSourcesV1alpha1RabbitmqSourceImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RabbitmqSource, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.RabbitmqSource{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapSourcesV1alpha1RabbitmqSourceImpl) Update(ctx context.Context, in *v1alpha1.RabbitmqSource, opts v1.UpdateOptions) (*v1alpha1.RabbitmqSource, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "sources.knative.dev",
		Version: "v1alpha1",
		Kind:    "RabbitmqSource",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.RabbitmqSource{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapSourcesV1alpha1RabbitmqSourceImpl) UpdateStatus(ctx context.Context, in *v1alpha1.RabbitmqSource, opts v1.UpdateOptions) (*v1alpha1.RabbitmqSource, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "sources.knative.dev",
		Version: "v1alpha1",
		Kind:    "RabbitmqSource",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.RabbitmqSource{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapSourcesV1alpha1RabbitmqSourceImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

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

package v1

import (
	time "time"

	nginxcontrollerv1 "github.com/gravitational/workshop/crd/controller/pkg/apis/nginxcontroller/v1"
	versioned "github.com/gravitational/workshop/crd/controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/gravitational/workshop/crd/controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/gravitational/workshop/crd/controller/pkg/generated/listers/nginxcontroller/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NginxInformer provides access to a shared informer and lister for
// Nginxes.
type NginxInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.NginxLister
}

type nginxInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNginxInformer constructs a new informer for Nginx type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNginxInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNginxInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNginxInformer constructs a new informer for Nginx type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNginxInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrainingV1().Nginxes(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrainingV1().Nginxes(namespace).Watch(options)
			},
		},
		&nginxcontrollerv1.Nginx{},
		resyncPeriod,
		indexers,
	)
}

func (f *nginxInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNginxInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nginxInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&nginxcontrollerv1.Nginx{}, f.defaultInformer)
}

func (f *nginxInformer) Lister() v1.NginxLister {
	return v1.NewNginxLister(f.Informer().GetIndexer())
}
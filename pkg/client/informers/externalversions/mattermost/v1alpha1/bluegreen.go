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

package v1alpha1

import (
	time "time"

	mattermostv1alpha1 "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1"
	versioned "github.com/mattermost/mattermost-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/mattermost/mattermost-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/mattermost/mattermost-operator/pkg/client/listers/mattermost/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BlueGreenInformer provides access to a shared informer and lister for
// BlueGreens.
type BlueGreenInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BlueGreenLister
}

type blueGreenInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBlueGreenInformer constructs a new informer for BlueGreen type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBlueGreenInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBlueGreenInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBlueGreenInformer constructs a new informer for BlueGreen type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBlueGreenInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MattermostV1alpha1().BlueGreens(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MattermostV1alpha1().BlueGreens(namespace).Watch(options)
			},
		},
		&mattermostv1alpha1.BlueGreen{},
		resyncPeriod,
		indexers,
	)
}

func (f *blueGreenInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBlueGreenInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *blueGreenInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mattermostv1alpha1.BlueGreen{}, f.defaultInformer)
}

func (f *blueGreenInformer) Lister() v1alpha1.BlueGreenLister {
	return v1alpha1.NewBlueGreenLister(f.Informer().GetIndexer())
}

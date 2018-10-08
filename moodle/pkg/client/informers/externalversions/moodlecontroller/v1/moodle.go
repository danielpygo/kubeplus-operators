/*
Copyright 2018 The Kubernetes Authors.

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

	moodlecontrollerv1 "github.com/cloud-ark/kubeplus-operators/moodle/pkg/apis/moodlecontroller/v1"
	versioned "github.com/cloud-ark/kubeplus-operators/moodle/pkg/client/clientset/versioned"
	internalinterfaces "github.com/cloud-ark/kubeplus-operators/moodle/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/cloud-ark/kubeplus-operators/moodle/pkg/client/listers/moodlecontroller/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MoodleInformer provides access to a shared informer and lister for
// Moodles.
type MoodleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MoodleLister
}

type moodleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMoodleInformer constructs a new informer for Moodle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMoodleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMoodleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMoodleInformer constructs a new informer for Moodle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMoodleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MoodlecontrollerV1().Moodles(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MoodlecontrollerV1().Moodles(namespace).Watch(options)
			},
		},
		&moodlecontrollerv1.Moodle{},
		resyncPeriod,
		indexers,
	)
}

func (f *moodleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMoodleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *moodleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&moodlecontrollerv1.Moodle{}, f.defaultInformer)
}

func (f *moodleInformer) Lister() v1.MoodleLister {
	return v1.NewMoodleLister(f.Informer().GetIndexer())
}

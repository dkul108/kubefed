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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	federatedscheduling_v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/federatedscheduling/v1alpha1"
	clientset "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/kubernetes-sigs/federation-v2/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/client/listers_generated/federatedscheduling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ReplicaSchedulingPreferenceInformer provides access to a shared informer and lister for
// ReplicaSchedulingPreferences.
type ReplicaSchedulingPreferenceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReplicaSchedulingPreferenceLister
}

type replicaSchedulingPreferenceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReplicaSchedulingPreferenceInformer constructs a new informer for ReplicaSchedulingPreference type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReplicaSchedulingPreferenceInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReplicaSchedulingPreferenceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReplicaSchedulingPreferenceInformer constructs a new informer for ReplicaSchedulingPreference type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReplicaSchedulingPreferenceInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederatedschedulingV1alpha1().ReplicaSchedulingPreferences(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederatedschedulingV1alpha1().ReplicaSchedulingPreferences(namespace).Watch(options)
			},
		},
		&federatedscheduling_v1alpha1.ReplicaSchedulingPreference{},
		resyncPeriod,
		indexers,
	)
}

func (f *replicaSchedulingPreferenceInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReplicaSchedulingPreferenceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *replicaSchedulingPreferenceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federatedscheduling_v1alpha1.ReplicaSchedulingPreference{}, f.defaultInformer)
}

func (f *replicaSchedulingPreferenceInformer) Lister() v1alpha1.ReplicaSchedulingPreferenceLister {
	return v1alpha1.NewReplicaSchedulingPreferenceLister(f.Informer().GetIndexer())
}

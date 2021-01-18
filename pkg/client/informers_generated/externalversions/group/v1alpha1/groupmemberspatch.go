/*
Copyright 2021 The Lynx Authors.

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
	"context"
	time "time"

	groupv1alpha1 "github.com/smartxworks/lynx/pkg/apis/group/v1alpha1"
	clientset "github.com/smartxworks/lynx/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/smartxworks/lynx/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/smartxworks/lynx/pkg/client/listers_generated/group/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GroupMembersPatchInformer provides access to a shared informer and lister for
// GroupMembersPatches.
type GroupMembersPatchInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GroupMembersPatchLister
}

type groupMembersPatchInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewGroupMembersPatchInformer constructs a new informer for GroupMembersPatch type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGroupMembersPatchInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGroupMembersPatchInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredGroupMembersPatchInformer constructs a new informer for GroupMembersPatch type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGroupMembersPatchInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GroupV1alpha1().GroupMembersPatches().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GroupV1alpha1().GroupMembersPatches().Watch(context.TODO(), options)
			},
		},
		&groupv1alpha1.GroupMembersPatch{},
		resyncPeriod,
		indexers,
	)
}

func (f *groupMembersPatchInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGroupMembersPatchInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *groupMembersPatchInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&groupv1alpha1.GroupMembersPatch{}, f.defaultInformer)
}

func (f *groupMembersPatchInformer) Lister() v1alpha1.GroupMembersPatchLister {
	return v1alpha1.NewGroupMembersPatchLister(f.Informer().GetIndexer())
}
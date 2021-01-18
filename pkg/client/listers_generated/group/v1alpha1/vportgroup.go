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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/smartxworks/lynx/pkg/apis/group/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VPortGroupLister helps list VPortGroups.
type VPortGroupLister interface {
	// List lists all VPortGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VPortGroup, err error)
	// Get retrieves the VPortGroup from the index for a given name.
	Get(name string) (*v1alpha1.VPortGroup, error)
	VPortGroupListerExpansion
}

// vPortGroupLister implements the VPortGroupLister interface.
type vPortGroupLister struct {
	indexer cache.Indexer
}

// NewVPortGroupLister returns a new VPortGroupLister.
func NewVPortGroupLister(indexer cache.Indexer) VPortGroupLister {
	return &vPortGroupLister{indexer: indexer}
}

// List lists all VPortGroups in the indexer.
func (s *vPortGroupLister) List(selector labels.Selector) (ret []*v1alpha1.VPortGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VPortGroup))
	})
	return ret, err
}

// Get retrieves the VPortGroup from the index for a given name.
func (s *vPortGroupLister) Get(name string) (*v1alpha1.VPortGroup, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vportgroup"), name)
	}
	return obj.(*v1alpha1.VPortGroup), nil
}
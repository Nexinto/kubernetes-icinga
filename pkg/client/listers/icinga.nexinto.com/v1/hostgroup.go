/*
Copyright 2018 Nexinto

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

package v1

import (
	v1 "github.com/Soluto-Private/kubernetes-icinga/pkg/apis/icinga.nexinto.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HostGroupLister helps list HostGroups.
type HostGroupLister interface {
	// List lists all HostGroups in the indexer.
	List(selector labels.Selector) (ret []*v1.HostGroup, err error)
	// HostGroups returns an object that can list and get HostGroups.
	HostGroups(namespace string) HostGroupNamespaceLister
	HostGroupListerExpansion
}

// hostGroupLister implements the HostGroupLister interface.
type hostGroupLister struct {
	indexer cache.Indexer
}

// NewHostGroupLister returns a new HostGroupLister.
func NewHostGroupLister(indexer cache.Indexer) HostGroupLister {
	return &hostGroupLister{indexer: indexer}
}

// List lists all HostGroups in the indexer.
func (s *hostGroupLister) List(selector labels.Selector) (ret []*v1.HostGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.HostGroup))
	})
	return ret, err
}

// HostGroups returns an object that can list and get HostGroups.
func (s *hostGroupLister) HostGroups(namespace string) HostGroupNamespaceLister {
	return hostGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HostGroupNamespaceLister helps list and get HostGroups.
type HostGroupNamespaceLister interface {
	// List lists all HostGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.HostGroup, err error)
	// Get retrieves the HostGroup from the indexer for a given namespace and name.
	Get(name string) (*v1.HostGroup, error)
	HostGroupNamespaceListerExpansion
}

// hostGroupNamespaceLister implements the HostGroupNamespaceLister
// interface.
type hostGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HostGroups in the indexer for a given namespace.
func (s hostGroupNamespaceLister) List(selector labels.Selector) (ret []*v1.HostGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.HostGroup))
	})
	return ret, err
}

// Get retrieves the HostGroup from the indexer for a given namespace and name.
func (s hostGroupNamespaceLister) Get(name string) (*v1.HostGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("hostgroup"), name)
	}
	return obj.(*v1.HostGroup), nil
}

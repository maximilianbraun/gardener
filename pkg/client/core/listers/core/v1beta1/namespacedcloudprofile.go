// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NamespacedCloudProfileLister helps list NamespacedCloudProfiles.
// All objects returned here must be treated as read-only.
type NamespacedCloudProfileLister interface {
	// List lists all NamespacedCloudProfiles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.NamespacedCloudProfile, err error)
	// NamespacedCloudProfiles returns an object that can list and get NamespacedCloudProfiles.
	NamespacedCloudProfiles(namespace string) NamespacedCloudProfileNamespaceLister
	NamespacedCloudProfileListerExpansion
}

// namespacedCloudProfileLister implements the NamespacedCloudProfileLister interface.
type namespacedCloudProfileLister struct {
	indexer cache.Indexer
}

// NewNamespacedCloudProfileLister returns a new NamespacedCloudProfileLister.
func NewNamespacedCloudProfileLister(indexer cache.Indexer) NamespacedCloudProfileLister {
	return &namespacedCloudProfileLister{indexer: indexer}
}

// List lists all NamespacedCloudProfiles in the indexer.
func (s *namespacedCloudProfileLister) List(selector labels.Selector) (ret []*v1beta1.NamespacedCloudProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.NamespacedCloudProfile))
	})
	return ret, err
}

// NamespacedCloudProfiles returns an object that can list and get NamespacedCloudProfiles.
func (s *namespacedCloudProfileLister) NamespacedCloudProfiles(namespace string) NamespacedCloudProfileNamespaceLister {
	return namespacedCloudProfileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NamespacedCloudProfileNamespaceLister helps list and get NamespacedCloudProfiles.
// All objects returned here must be treated as read-only.
type NamespacedCloudProfileNamespaceLister interface {
	// List lists all NamespacedCloudProfiles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.NamespacedCloudProfile, err error)
	// Get retrieves the NamespacedCloudProfile from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.NamespacedCloudProfile, error)
	NamespacedCloudProfileNamespaceListerExpansion
}

// namespacedCloudProfileNamespaceLister implements the NamespacedCloudProfileNamespaceLister
// interface.
type namespacedCloudProfileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NamespacedCloudProfiles in the indexer for a given namespace.
func (s namespacedCloudProfileNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.NamespacedCloudProfile, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.NamespacedCloudProfile))
	})
	return ret, err
}

// Get retrieves the NamespacedCloudProfile from the indexer for a given namespace and name.
func (s namespacedCloudProfileNamespaceLister) Get(name string) (*v1beta1.NamespacedCloudProfile, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("namespacedcloudprofile"), name)
	}
	return obj.(*v1beta1.NamespacedCloudProfile), nil
}
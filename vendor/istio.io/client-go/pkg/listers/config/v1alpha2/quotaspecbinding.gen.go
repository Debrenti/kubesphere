// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// QuotaSpecBindingLister helps list QuotaSpecBindings.
type QuotaSpecBindingLister interface {
	// List lists all QuotaSpecBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.QuotaSpecBinding, err error)
	// QuotaSpecBindings returns an object that can list and get QuotaSpecBindings.
	QuotaSpecBindings(namespace string) QuotaSpecBindingNamespaceLister
	QuotaSpecBindingListerExpansion
}

// quotaSpecBindingLister implements the QuotaSpecBindingLister interface.
type quotaSpecBindingLister struct {
	indexer cache.Indexer
}

// NewQuotaSpecBindingLister returns a new QuotaSpecBindingLister.
func NewQuotaSpecBindingLister(indexer cache.Indexer) QuotaSpecBindingLister {
	return &quotaSpecBindingLister{indexer: indexer}
}

// List lists all QuotaSpecBindings in the indexer.
func (s *quotaSpecBindingLister) List(selector labels.Selector) (ret []*v1alpha2.QuotaSpecBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.QuotaSpecBinding))
	})
	return ret, err
}

// QuotaSpecBindings returns an object that can list and get QuotaSpecBindings.
func (s *quotaSpecBindingLister) QuotaSpecBindings(namespace string) QuotaSpecBindingNamespaceLister {
	return quotaSpecBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// QuotaSpecBindingNamespaceLister helps list and get QuotaSpecBindings.
type QuotaSpecBindingNamespaceLister interface {
	// List lists all QuotaSpecBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.QuotaSpecBinding, err error)
	// Get retrieves the QuotaSpecBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.QuotaSpecBinding, error)
	QuotaSpecBindingNamespaceListerExpansion
}

// quotaSpecBindingNamespaceLister implements the QuotaSpecBindingNamespaceLister
// interface.
type quotaSpecBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all QuotaSpecBindings in the indexer for a given namespace.
func (s quotaSpecBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.QuotaSpecBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.QuotaSpecBinding))
	})
	return ret, err
}

// Get retrieves the QuotaSpecBinding from the indexer for a given namespace and name.
func (s quotaSpecBindingNamespaceLister) Get(name string) (*v1alpha2.QuotaSpecBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("quotaspecbinding"), name)
	}
	return obj.(*v1alpha2.QuotaSpecBinding), nil
}

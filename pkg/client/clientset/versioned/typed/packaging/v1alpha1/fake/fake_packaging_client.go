// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned/typed/packaging/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePackagingV1alpha1 struct {
	*testing.Fake
}

func (c *FakePackagingV1alpha1) PackageInstalls(namespace string) v1alpha1.PackageInstallInterface {
	return &FakePackageInstalls{c, namespace}
}

func (c *FakePackagingV1alpha1) PackageRepositories(namespace string) v1alpha1.PackageRepositoryInterface {
	return &FakePackageRepositories{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePackagingV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

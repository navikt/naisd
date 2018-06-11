package api

import (
	"fmt"
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	k8smeta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c clientHolder) createNamespace(name string) (*corev1.Namespace, error) {
	namespaceInterface := c.client.CoreV1().Namespaces()

	namespace, err := namespaceInterface.Get(name, k8smeta.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		glog.Error("Failed while getting namespace: unexpected error: %s", err)
		return nil, fmt.Errorf("unexpected error: %s", err)
	}

	if namespace != nil {
		glog.Info("Namespace %s already exists.", name)
		return namespace, nil
	}

	return namespaceInterface.Create(createNamespaceDef(name))
}

func createNamespaceDef(name string) *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: k8smeta.ObjectMeta{
			Name: name,
		},
	}
}
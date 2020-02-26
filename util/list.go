package util

import (
	v2 "github.com/kgysu/oc-wrapper/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListServerItems(name string, namespace string, types string, options string) error {


	project, err := v2.NewFromRemote(name, namespace, types, v1.ListOptions{})


}

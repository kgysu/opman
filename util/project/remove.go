package project

import (
	v2 "github.com/kgysu/oc-wrapper/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RemoveFromLocalProject(namespace string, folder string) error {
	project, err := v2.NewFromLocal(namespace, folder, false)
	if err != nil {
		return err
	}

	err = project.Delete(v1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

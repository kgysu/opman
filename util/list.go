package util

import (
	"fmt"
	v2 "github.com/kgysu/oc-wrapper/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListServerItems(name string, namespace string, types string, labelSelector string, limit int64) error {

	project, err := v2.NewFromRemote(name, namespace, parseTypes(types), v1.ListOptions{
		LabelSelector: labelSelector,
		Limit:         limit,
	})
	if err != nil {
		return err
	}

	fmt.Println("Objects: ")
	t := NewStandardTable(project.GetItems())
	t.Render()
	return nil
}

func parseTypes(input string) string {
	if input == "" {
		return v2.AllItemsKey
	}
	if input == "project" {
		return v2.ProjectObjectsKey
	}
	if input == "run" {
		return v2.ProjectMonitorItemsKey
	}
	return input
}

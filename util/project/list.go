package project

import (
	"fmt"
	v2 "github.com/kgysu/oc-wrapper/v2"
	"opman/util"
)

func ListLocalProjectFiles(namespace string, path string, debug bool) error {
	project, err := v2.NewFromLocal(namespace, path, debug)
	if err != nil {
		return err
	}

	fmt.Println("Objects: ")
	t := util.NewStandardTable(project.GetItems())
	t.Render()
	return nil
}

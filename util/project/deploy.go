package project

import v2 "github.com/kgysu/oc-wrapper/v2"

func DeployFromLocalProject(namespace string, folder string) error {
	project, err := v2.NewFromLocal(namespace, folder, false)
	if err != nil {
		return err
	}

	err = project.Create()
	if err != nil {
		return err
	}
	return nil
}

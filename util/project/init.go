package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func InitSampleProject(folder string) error {
	root := filepath.FromSlash(folder)
	deployments := filepath.FromSlash(folder + "/deployments")
	routes := filepath.FromSlash(folder + "/routes")
	secrets := filepath.FromSlash(folder + "/secrets")
	services := filepath.FromSlash(folder + "/services")
	// Create folders
	createDirIfNotExist(root)
	createDirIfNotExist(deployments)
	createDirIfNotExist(routes)
	createDirIfNotExist(secrets)
	createDirIfNotExist(services)
	// Create sample files
	err := ioutil.WriteFile(filepath.FromSlash(root+"/project.json"), SampleProjectJson, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.FromSlash(deployments+"/sample-deployment-config.json"), SampleDeploymentJson, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.FromSlash(routes+"/sample-route.json"), SampleRouteJson, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.FromSlash(secrets+"/sample-secret.json"), SampleSecretJson, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.FromSlash(services+"/sample-service.json"), SampleServiceJson, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("-> Project initialized! All Sample files were created successfully in subfolder '%s'", root)
	return nil
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

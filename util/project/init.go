package project

import (
	"fmt"
	"io/ioutil"
	"os"
)

func InitSampleProject() error {
	// Create folders
	createDirIfNotExist("items")
	createDirIfNotExist("items/deployments")
	createDirIfNotExist("items/routes")
	createDirIfNotExist("items/secrets")
	createDirIfNotExist("items/services")
	// Create sample files
	err := ioutil.WriteFile("./items/project.json", SampleProjectJson, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./items/deployments/sample-deployment.json", SampleDeploymentJson, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./items/routes/sample-route.json", SampleRouteJson, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./items/secrets/sample-secret.json", SampleSecretJson, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./items/services/sample-service.json", SampleServiceJson, 0644)
	if err != nil {
		return err
	}
	fmt.Println("-> Project initialized! All Sample files were created successfully in subfolder './items'")
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

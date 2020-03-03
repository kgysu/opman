# Openshift-Manager - opman
Lightweight go application to operate on an openshift project. Setup your Openshift-Project
locally and deploy it to remote server in minutes. This makes it easy to test new poc-setups
and also to manage large projects with many items.

## Usage
### Commands
You can use the following commands:
 - `list` - List items filtered by Type or LabelSelector.
 - `monitor` - Monitor items. (List with refresh interval)
 - `watch` - Watching specific items.
 - `project` - Work with local project. See Local Project.
 - `help` - Getting help.
 
For more details run: `opman <command> --help`

### Options
For every command running on remote Openshift-Cluster or working with local project files, 
**you need to specify the current namespace (openshift project) to operate**.

You can set the `namespace` with one of the following options:
 - `-n <namespace>` 
 - `--namespace=<namespace>` 

### Local Project
The command 'project' uses local files in folder `./items/` to run commands.
There you can add all openshift items as JSON files like DeploymentConfigs or other
types. These files are like templates wich can also contain stage dependend values. 
You can edit the location of these items with the folder option `--folder=/new/path`.

To start with local project you can initialize a sample setup with `opman project init`.
This creates the following sample files in `./items/`
 - `./items/project.json` Project configuration file.
 - `./items/deployments/sample-deployment-config.json` Sample DeploymentConfig.
 - `./items/services/sample-service.json` Sample Service.
 - `./items/routes/sample-route.json` Project configuration file.

After setting up your local project, you can now run the following commands:
 - `list` List all the local project items.
 - `deploy` Create all items on remote Openshift-Cluster.
 - `remove` - Delete all items in remote Openshift-Cluster.
 - `compare` - Show wich items are deployed on remote Openshift-Cluster.

#### Stage dependend values
In the local project item json-files you can specify any value as a stage dependen replacement.
Just name the value with a `$` "Dollar-Sign" prefix. Then you can add a new Stage(Namespace)
environment file like `./items/dev.env`. There you can specify all real values like 
`KEY=VALUE`. 

For Example:

Add a new key in a local item file:
 ```json
{
  "name": "$MY_EXAMPLE_NAME"
}
```
Add a new value to local env file:
```
MY_EXAMPLE_NAME=myname
```
You can verify the replacements with the `list` command and the `--debug`(`-d`) option enabled.
 
#### Types
Types are used to filter items in several commands like for example in `list`. 
And also as type recognition for local project items. The files must be named with a
type as suffix. For example: if you want to create a new `DeploymentConfig`, create a
file named `name-deployment-config.json`. It does not care about upper or lower case and 
ignores additional minus chars. 

The following types are possible (* = type can be used in local project):
 - DeploymentConfig*
 - StatefulSet*
 - Service*
 - Route*
 - ConfigMap*
 - PersistenVolume*
 - PersistenVolumeClaim*
 - ServiceAccount*
 - Role*
 - RoleBinding*
 - Pod
 - Event

#### ConfigMaps
To manage `ConfigMaps` you can also create one without any data in it. If opman finds an empty
ConfigMap-Data it will automatically add all files found in `./items/config/<namespace>/`
as data for this ConfigMap. Namespace must match with the folder name!

## Install

To install deployer from source you have to install go and then you can run the following commands:
```bash
git clone https://github.com/kgysu/opman.git
cd opman
go build
go install
```
You have to add `$GOPATH/bin` to your env PATH!

## Contribute
tbd

## License
Openshift-Manager(Opman) is licensed under the Apache License 2.0.
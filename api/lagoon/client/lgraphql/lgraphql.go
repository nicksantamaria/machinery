// Code generated for package lgraphql by go-bindata DO NOT EDIT. (@generated)
// sources:
// _lgraphql/addEnvVariable.graphql
// _lgraphql/addNotificationEmail.graphql
// _lgraphql/addNotificationMicrosoftTeams.graphql
// _lgraphql/addNotificationRocketChat.graphql
// _lgraphql/addNotificationSlack.graphql
// _lgraphql/addNotificationToProject.graphql
// _lgraphql/lagoonSchema.graphql
// _lgraphql/lagoonVersion.graphql
// _lgraphql/sshEndpointsByProject.graphql
// _lgraphql/taskByID.graphql
// _lgraphql/deployments/deployEnvironmentBranch.graphql
// _lgraphql/deployments/deployEnvironmentLatest.graphql
// _lgraphql/deployments/deployEnvironmentPromote.graphql
// _lgraphql/deployments/deployEnvironmentPullrequest.graphql
// _lgraphql/deployments/deploymentByName.graphql
// _lgraphql/deployments/deploymentByRemoteID.graphql
// _lgraphql/deployments/getDeploymentsByBulkID.graphql
// _lgraphql/deployments/getDeploymentsForEnvironment.graphql
// _lgraphql/deployments/updateDeployment.graphql
// _lgraphql/deploytargets/addDeployTarget.graphql
// _lgraphql/deploytargets/deleteDeployTarget.graphql
// _lgraphql/deploytargets/listDeployTargets.graphql
// _lgraphql/deploytargets/updateDeployTarget.graphql
// _lgraphql/deploytargetconfigs/addDeployTargetConfig.graphql
// _lgraphql/deploytargetconfigs/deleteDeployTargetConfig.graphql
// _lgraphql/deploytargetconfigs/deployTargetConfigsByProjectId.graphql
// _lgraphql/deploytargetconfigs/updateDeployTargetConfig.graphql
// _lgraphql/projects/addProject.graphql
// _lgraphql/projects/minimalProjectByName.graphql
// _lgraphql/projects/projectByName.graphql
// _lgraphql/projects/projectByNameMetadata.graphql
// _lgraphql/projects/projectsByMetadata.graphql
// _lgraphql/projects/removeProjectMetadataByKey.graphql
// _lgraphql/projects/updateProject.graphql
// _lgraphql/projects/updateProjectMetadata.graphql
// _lgraphql/environments/addOrUpdateEnvironment.graphql
// _lgraphql/environments/addOrUpdateEnvironmentStorage.graphql
// _lgraphql/environments/addRestore.graphql
// _lgraphql/environments/backupsForEnvironmentByName.graphql
// _lgraphql/environments/deleteEnvironment.graphql
// _lgraphql/environments/environmentByName.graphql
// _lgraphql/environments/environmentByNamespace.graphql
// _lgraphql/environments/setEnvironmentServices.graphql
// _lgraphql/environments/updateEnvironment.graphql
// _lgraphql/tasks/switchActiveStandby.graphql
// _lgraphql/tasks/updateTask.graphql
// _lgraphql/usergroups/addGroup.graphql
// _lgraphql/usergroups/addGroupsToProject.graphql
// _lgraphql/usergroups/addSshKey.graphql
// _lgraphql/usergroups/addUser.graphql
// _lgraphql/usergroups/addUserToGroup.graphql
// _lgraphql/usergroups/allGroupMembers.graphql
// _lgraphql/usergroups/allUsers.graphql
// _lgraphql/usergroups/me.graphql
// _lgraphql/usergroups/removeGroupsFromProject.graphql
// _lgraphql/usergroups/removeUserFromGroup.graphql
// _lgraphql/usergroups/userByEmail.graphql
package lgraphql

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __lgraphqlAddenvvariableGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\xb1\xca\xc3\x20\x14\x46\xf7\x3c\xc5\xf7\xc3\x3f\x24\x90\x27\x70\xef\x90\x39\xa5\xfb\x6d\x95\x22\x24\x37\x21\xbd\x06\x42\xe9\xbb\x17\xbd\xda\x4a\x07\x97\x73\xd4\xef\xcc\x41\x48\xfc\xc2\x68\x1b\xe0\x5f\x8e\xd5\x19\x9c\x78\xbf\xd0\xe6\xe9\x3a\xb9\xf3\xb1\xba\xbf\xbe\xa8\xc1\x1a\x0c\x2c\x0a\x1e\xb7\xe5\xe7\xf2\x18\x89\x4a\xa6\xd9\x19\x8c\xb2\x79\xbe\x2b\xd9\x69\x0a\x5f\xd4\xe1\xd9\x00\x00\x59\x5b\x7d\xd0\x7a\x5e\x83\x98\xec\x00\xcd\x49\xd3\x7d\x85\x62\x46\xee\x29\x38\xc7\x68\x54\x81\x1a\x91\x5a\x0a\xca\x15\x5a\x93\xd8\xab\xfb\xcc\x79\x5b\x3d\x54\xd9\xc4\xf3\x0e\x00\x00\xff\xff\x1d\x4f\x13\xd7\x24\x01\x00\x00")

func _lgraphqlAddenvvariableGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddenvvariableGraphql,
		"_lgraphql/addEnvVariable.graphql",
	)
}

func _lgraphqlAddenvvariableGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddenvvariableGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addEnvVariable.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\xa4\xe6\x26\x66\xe6\x38\xa6\xa4\x14\xa5\x16\x17\xc3\x65\x34\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\xfc\xf2\x4b\x32\xd3\x32\x93\xc1\x66\xb8\x82\xd4\x6a\x64\xe6\x15\x94\x96\x58\x41\x55\x28\x28\x40\x8c\x04\x9b\xac\x03\x15\x42\x35\x13\xc5\x0a\xb0\x8a\x5a\x4d\xb8\xee\xcc\x14\x24\x63\x20\x92\x5c\x20\x0c\x08\x00\x00\xff\xff\x7d\x3b\x64\xad\xb7\x00\x00\x00")

func _lgraphqlAddnotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationemailGraphql,
		"_lgraphql/addNotificationEmail.graphql",
	)
}

func _lgraphqlAddnotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\x94\xa7\x26\x65\xe4\xe7\x67\xc3\x05\x35\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\xfc\xf2\x4b\x32\xd3\x32\x93\xc1\xda\x7d\x33\x93\x8b\xf2\x8b\xf3\xd3\x4a\x42\x52\x13\x73\x8b\x35\x32\xf3\x0a\x4a\x4b\xac\xa0\x4a\x15\x14\x20\xc6\x82\x4d\xd7\x81\x0a\xc1\xcd\x85\xd9\x00\x16\xaf\xd5\x84\xeb\xc9\x4c\x41\xd2\x0c\x91\xe4\x02\x61\x40\x00\x00\x00\xff\xff\x97\x56\x46\x93\xb1\x00\x00\x00")

func _lgraphqlAddnotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationmicrosoftteamsGraphql,
		"_lgraphql/addNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlAddnotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x31\x0e\xc2\x30\x0c\x45\xf7\x9c\xe2\x23\x31\xb4\x52\x4f\xd0\x95\x9d\x01\x4e\x60\x92\x40\xac\x52\x1b\x21\x57\x0c\x88\xbb\xa3\xa6\x49\x04\x83\x97\xf7\xad\xf7\xe6\xc5\xc8\x58\x05\x9d\x03\xf6\x42\x73\x1c\x71\xb6\x27\xcb\x6d\x37\xac\xc4\x27\x12\x89\xf7\x7f\xf8\x8a\x97\xa4\x3a\x35\xd8\xe3\xed\x00\x80\x42\x38\xaa\xf1\x95\x7d\x76\x9e\xd4\x4f\xd1\x0e\x89\xac\x63\x79\x2c\x36\x96\x37\x60\xeb\xe4\xdc\x50\x50\x0b\xd5\x64\x1d\x5a\xac\x66\x33\xff\xf4\x4d\xc6\xe1\xc7\xba\x8d\x6e\xbd\x6f\x00\x00\x00\xff\xff\x53\x0f\xa2\x30\xdb\x00\x00\x00")

func _lgraphqlAddnotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationrocketchatGraphql,
		"_lgraphql/addNotificationRocketChat.graphql",
	)
}

func _lgraphqlAddnotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x41\x0e\x82\x40\x0c\x45\xf7\x73\x8a\x6f\xe2\x02\x12\x4e\xc0\x21\xdc\x70\x82\x3a\x33\x4a\x03\xb4\xc6\x94\xb8\x30\xde\xdd\x30\x30\x8d\x2e\xba\x79\xbf\x79\x6f\x59\x8d\x8c\x55\xd0\x04\xe0\x2c\xb4\xe4\x1e\x83\x3d\x59\xee\xa7\x6e\x23\x71\x24\x91\x3c\xff\xc3\x57\xbe\x8e\xaa\x93\xc3\x16\xef\x00\x00\x94\xd2\x45\x8d\x6f\x1c\x8b\x73\x98\x29\x4e\x0d\xcb\x63\xb5\xfe\xf8\x00\xf6\x44\x29\x75\x07\xf2\x46\xad\xd5\xc1\x3b\xb5\x58\xf8\xa7\x75\x19\xa7\x1f\xeb\x3e\x86\xed\xbe\x01\x00\x00\xff\xff\xd9\x6e\xb4\xda\xd6\x00\x00\x00")

func _lgraphqlAddnotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationslackGraphql,
		"_lgraphql/addNotificationSlack.graphql",
	)
}

func _lgraphqlAddnotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationtoprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x09\xe6\xe5\x97\x64\xa6\x65\x26\x83\x95\x85\x54\x16\xa4\x5a\x29\xf8\xa1\x89\x60\xaa\xf3\x4b\xcc\x4d\x85\x9b\xa2\xa9\x50\xcd\xa5\xa0\xa0\xa0\x90\x98\x92\x82\xa2\x33\x3f\x00\x62\x9d\x46\x66\x5e\x41\x69\x89\x15\x54\x95\x82\x02\xdc\x15\x30\xf7\x40\xc5\x31\x1d\x82\xe1\x36\x2c\x2a\x21\x4e\xc1\x70\x1d\x58\x65\xad\x26\xdc\xd2\xcc\x14\x98\x5e\xb8\x24\x17\x08\x03\x02\x00\x00\xff\xff\x5d\xba\xcd\xcd\x21\x01\x00\x00")

func _lgraphqlAddnotificationtoprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationtoprojectGraphql,
		"_lgraphql/addNotificationToProject.graphql",
	)
}

func _lgraphqlAddnotificationtoprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationtoprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationToProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlLagoonschemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x88\x8f\x2f\x4e\xce\x48\xcd\x4d\x04\x73\x14\x14\x4a\x2a\x0b\x52\x8b\xa1\x6c\x05\x85\xec\xcc\xbc\x14\x28\x33\x2f\x31\x37\x15\xca\x4c\xcb\x4c\xcd\x49\x29\xd6\xc8\xcc\x4b\xce\x29\x4d\x49\x75\x49\x2d\x28\x4a\x4d\x4e\x2c\x49\x4d\xb1\x2a\x29\x2a\x4d\xd5\x84\x6b\x46\xd1\x53\xcb\x05\x23\x6b\xb9\x6a\x01\x01\x00\x00\xff\xff\x29\x07\x39\xef\x7e\x00\x00\x00")

func _lgraphqlLagoonschemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlLagoonschemaGraphql,
		"_lgraphql/lagoonSchema.graphql",
	)
}

func _lgraphqlLagoonschemaGraphql() (*asset, error) {
	bytes, err := _lgraphqlLagoonschemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/lagoonSchema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlLagoonversionGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\xc8\x49\x4c\xcf\xcf\xcf\x0b\x4b\x2d\x2a\xce\xcc\xcf\xe3\xaa\x05\x04\x00\x00\xff\xff\x42\xb4\x77\x45\x19\x00\x00\x00")

func _lgraphqlLagoonversionGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlLagoonversionGraphql,
		"_lgraphql/lagoonVersion.graphql",
	)
}

func _lgraphqlLagoonversionGraphql() (*asset, error) {
	bytes, err := _lgraphqlLagoonversionGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/lagoonVersion.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlSshendpointsbyprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x00\xab\x83\x48\x2b\x28\x64\xa6\x80\x29\x90\x10\x98\x91\x9a\x57\x96\x59\x94\x9f\x97\x9b\x9a\x57\x52\x0c\x55\x02\x57\x84\xa4\x4c\x41\x21\xbf\x20\x35\xaf\x38\x23\x33\xad\x24\x00\x62\x85\x1f\x16\x29\xb8\x01\x48\x46\x28\x28\x14\x17\x67\x78\xe4\x17\x97\x20\xf3\x03\xf2\x8b\x60\xfc\x5a\x2e\x18\x59\xcb\x55\xcb\x05\x08\x00\x00\xff\xff\x1f\x6e\x6a\x1a\xdb\x00\x00\x00")

func _lgraphqlSshendpointsbyprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlSshendpointsbyprojectGraphql,
		"_lgraphql/sshEndpointsByProject.graphql",
	)
}

func _lgraphqlSshendpointsbyprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlSshendpointsbyprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/sshEndpointsByProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlTaskbyidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8c\x41\x0e\xc2\x30\x0c\x04\xef\x79\xc5\x22\xf5\x50\xbe\xc0\x91\x5b\x9e\x61\xd5\x16\x8a\xa8\x13\x70\x5c\xa4\x0a\xf5\xef\xa8\x29\xe9\xc9\xe3\xdd\xd5\xbc\x17\xb1\x15\x63\x00\x86\xc4\x37\xc4\xec\x97\x00\x5c\xbf\x01\x70\xaa\xcf\xfb\x1a\x79\xdc\x8b\x21\x71\x0b\x81\xc4\xed\x64\x52\x69\x30\x15\x55\xca\x47\x58\x9d\x7c\xa9\x1d\xcd\x85\xfb\xe4\x35\xcb\xf9\x99\x50\xe7\xb9\x3c\x8e\xb9\x89\x16\x97\xf8\xd7\x88\x7d\xd2\xb4\xeb\xb7\xb0\xfd\x02\x00\x00\xff\xff\x74\xe4\xc2\x1e\xa2\x00\x00\x00")

func _lgraphqlTaskbyidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlTaskbyidGraphql,
		"_lgraphql/taskByID.graphql",
	)
}

func _lgraphqlTaskbyidGraphql() (*asset, error) {
	bytes, err := _lgraphqlTaskbyidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/taskByID.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeployenvironmentbranchGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x2f\xb8\x68\xa1\x27\xc8\xb2\xe8\xd6\x85\x9e\x20\xd6\x51\x23\xed\xa4\x0c\x53\x41\xa4\x77\x97\x36\x89\x56\x70\x35\xcc\xfb\x9f\xcf\xeb\x47\x75\xea\x03\xa3\x34\xc0\x66\x90\x70\xa7\x56\x2d\x8e\x2a\x9e\xaf\x45\x8d\x99\x9e\xc4\x71\x7b\xfb\x0b\x0f\x74\xc9\xbc\x9e\xa9\x90\x8e\xc2\x5b\xa7\xce\xa2\x09\xa1\x23\xc7\x45\x85\x97\x01\x80\x33\x0d\x5d\x78\xee\xf8\xe1\x25\x70\x4f\xac\xcd\x32\x51\x7a\x1e\x46\xb5\xa9\x04\x64\x87\xfc\x03\xec\x7a\xb2\x1f\xb9\x84\xa7\x74\xa3\xc7\x3e\x56\xe2\xf3\x93\x2c\x86\x5f\xdb\x94\xad\x3d\x57\xd2\x26\x2f\x57\x66\x7a\x07\x00\x00\xff\xff\xb1\xcb\x5e\x73\x1a\x01\x00\x00")

func _lgraphqlDeploymentsDeployenvironmentbranchGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeployenvironmentbranchGraphql,
		"_lgraphql/deployments/deployEnvironmentBranch.graphql",
	)
}

func _lgraphqlDeploymentsDeployenvironmentbranchGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeployenvironmentbranchGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deployEnvironmentBranch.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeployenvironmentlatestGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\xb1\x6a\xc4\x30\x0c\x86\x77\x3f\x85\x0a\x1d\x7a\xaf\xe0\xb9\x37\x84\x96\x2e\x85\x5b\x4a\x07\x1d\x11\x45\xd4\x91\x83\x22\x07\x42\xc9\xbb\x17\x27\x4e\xe2\x9c\x37\xff\xfa\x2c\x7f\x7f\x97\x0c\x8d\xa3\xc0\x8b\x03\x78\x26\x19\x59\xa3\x74\x24\xe6\xe1\x7a\x5c\x1a\xe9\x93\x3d\x65\xe2\x9e\xc2\x6f\xd3\x7a\xf8\x34\x65\xf9\xd9\x92\x0f\xec\xa8\xce\x7a\xe5\xa8\x6c\x93\x87\x46\x6c\x85\x38\xb4\x37\x54\xc6\x7b\xa0\xc1\xc3\xd7\x55\xc6\x37\x9a\x6e\x18\x12\x2d\xcb\xbf\xdd\x05\xfe\x1c\x40\x4b\x7d\x88\x53\xf5\xf5\x3b\x1a\x0d\x96\xed\x00\x38\x93\x7e\xe1\xf2\x39\xd9\xd6\xee\x65\xbe\xb9\x16\xe9\x2a\x5d\x7d\x77\xf5\x32\x39\xac\xf7\x02\xfb\x9b\xb3\xfe\x43\x9f\x42\x29\x59\x52\x79\x45\x43\x0f\xa6\x69\x5d\x3b\x3b\x80\x8b\x9b\xdd\x7f\x00\x00\x00\xff\xff\xf3\x83\xe7\x9d\x69\x01\x00\x00")

func _lgraphqlDeploymentsDeployenvironmentlatestGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeployenvironmentlatestGraphql,
		"_lgraphql/deployments/deployEnvironmentLatest.graphql",
	)
}

func _lgraphqlDeploymentsDeployenvironmentlatestGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeployenvironmentlatestGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deployEnvironmentLatest.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeployenvironmentpromoteGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\xaa\x83\x30\x10\x45\xf7\xf9\x8a\x2b\xbc\x85\xc2\xfb\x82\x2c\x1f\xaf\xfb\x42\xbf\x20\xe8\x50\x52\xcc\x8c\xc4\x49\xa1\x14\xff\xbd\xd8\x1a\x6b\xab\xae\x02\xb9\xe7\x0c\x33\x37\x24\x75\xea\x85\x51\x1a\xe0\xa7\x8b\x72\xa1\x5a\x2d\x4e\x1a\x3d\x9f\x8b\x5f\x8c\xbf\xbd\xa4\x58\xd3\x81\xaf\x3e\x0a\x07\xe2\xef\xbc\xa1\x5e\x3d\x3f\xc7\x6c\x42\x23\x13\x49\x53\xe4\x7f\xa7\xce\xe2\x4f\xa4\x25\xc7\x45\x85\xbb\x01\x80\x86\xba\x56\x6e\x0b\xf5\x18\x25\x88\x52\xe9\xb9\x4b\x6a\x5f\x10\xb0\x5e\x23\x27\x00\xbb\x40\x76\x63\xd3\x19\xc8\x97\xbd\x95\x59\x9a\xa2\x39\x18\xcc\xe7\xbb\x56\x37\xc5\x8c\xef\xb5\xb1\x53\xd3\x64\x2d\xfb\x59\x94\x65\xf2\xe4\xca\x0c\x8f\x00\x00\x00\xff\xff\xcd\xc9\x92\xe1\xab\x01\x00\x00")

func _lgraphqlDeploymentsDeployenvironmentpromoteGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeployenvironmentpromoteGraphql,
		"_lgraphql/deployments/deployEnvironmentPromote.graphql",
	)
}

func _lgraphqlDeploymentsDeployenvironmentpromoteGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeployenvironmentpromoteGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deployEnvironmentPromote.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeployenvironmentpullrequestGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x41\x4e\xc3\x30\x10\x45\xf7\x39\xc5\x8f\xc4\xa2\x95\x7a\x02\x2f\x2b\x58\x74\x83\x2a\x38\xc1\xb4\x1d\xa8\x91\x33\x0e\xd3\x31\x12\x42\xbd\x3b\x4a\x33\x85\xc6\x51\x77\xc9\xff\x7f\x2c\xbd\xd7\x15\x23\x8b\x59\xb0\x68\x80\x87\x5e\xf3\x07\xef\x2d\x60\x3b\x7e\x6c\xa4\x2f\xd6\xae\x30\x74\x52\xba\x1d\x6b\xc0\x46\xac\x5d\x0d\x81\x45\x4b\x1c\xf0\x6a\x1a\xe5\x7d\x8c\x76\x74\xe2\xb5\x92\xec\x8f\xcf\xd4\xdd\xed\x5e\xf8\x6d\x5a\x1d\x99\x0e\xf7\xce\xfe\xbb\xd9\x99\xb2\x15\x95\x47\x32\x0a\x58\xe7\x9c\x98\xa4\x5d\xe2\xa7\x01\x80\x03\xf7\x29\x7f\x3f\xc9\x57\xd4\x2c\x1d\x8b\x6d\x4b\x4a\xca\x9f\x85\x4f\xb6\x88\x03\x56\xf0\x25\xf0\x47\x7d\xe5\xf7\xfc\x4a\xec\xe8\x9e\x3a\xf6\x88\xef\x59\xcd\x5d\x89\x98\xad\x2e\x28\x53\x23\xbe\xa9\x55\x54\x6e\x66\xab\xf1\xa5\xc9\xbf\x6f\x6e\xed\xdc\xa8\xba\xb4\xe7\x06\x58\x36\xe7\xdf\x00\x00\x00\xff\xff\x49\x97\xf5\x67\xfd\x01\x00\x00")

func _lgraphqlDeploymentsDeployenvironmentpullrequestGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeployenvironmentpullrequestGraphql,
		"_lgraphql/deployments/deployEnvironmentPullrequest.graphql",
	)
}

func _lgraphqlDeploymentsDeployenvironmentpullrequestGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeployenvironmentpullrequestGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deployEnvironmentPullrequest.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeploymentbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\xae\xc2\x20\x10\x86\xf7\x3d\x05\xaf\x79\x8b\xbe\x2b\x74\xf9\x76\x26\x46\x4d\x3c\x01\xa1\xa3\xa2\x65\xc0\x61\x6a\xd2\x34\xdc\xdd\x50\x63\x5b\x68\x67\x35\x7c\x0c\xdf\xfc\xe1\xd9\x01\xf5\xa2\x01\xd7\xda\xde\x00\xf2\x7f\x7f\x90\x06\xaa\x5f\x94\x06\xbc\x93\x0a\x6a\x71\x66\xd2\x78\xfd\xf9\x13\x43\x21\x84\x58\xcd\xd6\x80\x2f\x4d\x16\x3f\xe4\xe8\x00\xfd\x4d\x5f\xf8\x44\xf6\x0e\x8a\x47\x9b\xdd\x80\xb5\x98\x77\x7c\xd5\xb1\x74\x33\xb5\xf1\x7e\x3a\xcc\x6b\x7d\x85\xe3\xfb\x72\x18\x66\x18\x95\x21\x94\x4b\x53\x66\x5b\x19\x63\x79\x96\xdc\xf9\x04\x29\x02\xc9\xd0\xe4\x63\x94\x33\x65\x8d\x6b\x21\xa7\x04\xc6\x32\xec\x52\xd8\xe9\xbd\xc6\x47\x82\x16\x9f\x96\x45\xde\xcc\x19\x8a\xb4\x0b\x45\x78\x07\x00\x00\xff\xff\x45\xbf\xf5\x0a\xba\x01\x00\x00")

func _lgraphqlDeploymentsDeploymentbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeploymentbynameGraphql,
		"_lgraphql/deployments/deploymentByName.graphql",
	)
}

func _lgraphqlDeploymentsDeploymentbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeploymentbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deploymentByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeploymentbyremoteidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\x48\x49\x2d\xc8\xc9\xaf\xcc\x4d\xcd\x2b\x71\xaa\x0c\x4a\xcd\xcd\x2f\x49\xf5\x4c\xd1\x50\x29\x82\xb2\xac\x14\x82\x4b\x8a\x32\xf3\xd2\x15\x35\x15\xaa\xb9\x14\x14\x14\xb0\x2b\xcf\x4c\xb1\x52\x80\x6b\x81\xa9\x04\x81\xcc\x14\x38\x33\x2f\x31\x37\x15\xce\x29\x2e\x49\x2c\x29\x2d\x86\x73\x93\x8b\x52\x13\x4b\x52\x53\x90\xa5\x8b\x90\xf9\xc9\xf9\xb9\x05\x39\xa9\xc8\x22\x30\xcb\xe0\x02\xa5\x99\x3e\x99\x79\xd9\x70\x6e\x6a\x5e\x59\x66\x51\x7e\x1e\xc8\x9d\x48\xce\xc1\x70\x47\x2d\x17\x84\xac\x05\x04\x00\x00\xff\xff\x9f\x95\x0e\x67\x0c\x01\x00\x00")

func _lgraphqlDeploymentsDeploymentbyremoteidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeploymentbyremoteidGraphql,
		"_lgraphql/deployments/deploymentByRemoteID.graphql",
	)
}

func _lgraphqlDeploymentsDeploymentbyremoteidGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeploymentbyremoteidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deploymentByRemoteID.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsGetdeploymentsbybulkidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\x49\x2a\xcd\xc9\xf6\x4c\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\x2d\xc8\xc9\xaf\xcc\x4d\xcd\x2b\x29\x76\xaa\x74\x02\xab\xd1\x80\x29\x85\xea\xd1\x04\x2b\x84\x28\x07\x81\xcc\x14\x38\x33\x2f\x31\x37\x15\xce\x29\x4a\xcd\xcd\x2f\x49\xf5\x44\xc8\x16\x97\x24\x96\x94\x16\xc3\xb9\x49\xa5\x99\x39\x29\x3e\xf9\xe9\x70\x81\xd4\xbc\xb2\xcc\xa2\xfc\x3c\x90\xe5\x08\xd3\xd1\x6c\xc0\xb0\x05\x04\x0a\x8a\xf2\xb3\x52\x93\xd1\x34\x61\xd1\x88\x55\x73\x2d\x17\x2a\xab\x96\xab\x16\x10\x00\x00\xff\xff\x26\x1b\x9b\xfe\x24\x01\x00\x00")

func _lgraphqlDeploymentsGetdeploymentsbybulkidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsGetdeploymentsbybulkidGraphql,
		"_lgraphql/deployments/getDeploymentsByBulkID.graphql",
	)
}

func _lgraphqlDeploymentsGetdeploymentsbybulkidGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsGetdeploymentsbybulkidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/getDeploymentsByBulkID.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8c\x41\x0a\x83\x30\x10\x45\xf7\x9e\x62\x04\x17\x16\x7a\x02\x97\xdd\xb9\xe9\xa6\x27\x08\xfa\x29\x16\x33\xb1\x93\xb1\x10\xc4\xbb\x17\x6d\x2d\x4e\xfd\x8b\x40\xde\x4b\xde\x73\x84\x24\x2a\x8b\x41\xc2\x03\x8d\x56\x54\xb3\xe6\x67\x2a\xc0\xaf\x4e\x02\x7b\xb0\x56\x74\x53\xe9\xf8\x9e\x9f\xa6\x8c\x88\x68\xa7\x2e\xe9\xea\x3c\xca\x15\x6f\xfb\x95\xb6\xa6\xb1\xec\x3c\x2a\x93\x5f\xf5\x37\xbd\xac\xc5\xd0\x87\xb4\x98\x38\x1d\xbe\x1a\xd0\xb5\xe6\x2a\xf0\x41\x51\x5b\x18\xd5\xe9\x18\x0d\x6a\x04\x4e\x71\x78\x26\xff\xac\x09\x7e\xe8\xb1\xa7\x73\xf6\x39\xe7\x77\x00\x00\x00\xff\xff\x88\xb9\x03\x28\x36\x01\x00\x00")

func _lgraphqlDeploymentsGetdeploymentsforenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql,
		"_lgraphql/deployments/getDeploymentsForEnvironment.graphql",
	)
}

func _lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsGetdeploymentsforenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/getDeploymentsForEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsUpdatedeploymentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\x31\x8e\x84\x30\x0c\x45\x7b\x4e\x61\xa4\x2d\x96\x2b\xa4\xde\x06\x69\x8b\x69\xe6\x00\x19\x62\x69\x2c\x88\x13\x05\x67\x24\x84\x72\xf7\x51\x00\x85\x40\x2a\xff\xe7\x6f\xff\xd8\x46\xd1\x42\x8e\xe1\xb7\x01\x00\xf8\x21\xa3\xa0\x67\x69\x77\xe5\xb5\x0c\x6f\x05\x4f\x6f\xb4\xe0\x1f\xfa\xc9\x2d\x16\x59\x1e\x19\xf7\xec\xa3\xb4\xdd\xba\x39\xe3\xcd\xb1\x6f\xcb\x8f\xb2\x4d\xc1\x5a\xc0\x06\x8d\xca\x51\x17\x76\x64\xed\x99\xa5\x93\xb6\xaa\x3b\xc7\xab\x29\xd6\x16\x8b\x98\x45\x4b\x9c\x8b\x1c\x02\x6a\x41\x53\xb7\x43\xad\x07\x67\xfd\x84\x35\x09\x68\x9d\x60\x7f\x82\x48\xff\xc4\x63\x91\xc8\x1f\x0a\x8e\xf3\x75\xb7\x6b\x2e\xff\x48\xa5\xf2\x81\x5c\x20\x59\x0a\x78\xc5\x69\x3c\xf6\xa7\x26\x7d\x03\x00\x00\xff\xff\xdd\x19\x4b\x1e\x7a\x01\x00\x00")

func _lgraphqlDeploymentsUpdatedeploymentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsUpdatedeploymentGraphql,
		"_lgraphql/deployments/updateDeployment.graphql",
	)
}

func _lgraphqlDeploymentsUpdatedeploymentGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsUpdatedeploymentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/updateDeployment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsAdddeploytargetGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xd1\xcf\x4e\xc3\x30\x0c\x06\xf0\x7b\x9f\xc2\x48\x1c\x36\x69\x4f\xd0\x2b\x1c\x80\xc3\x56\x31\x78\x80\x40\xdc\x2e\x22\x8d\x2b\xd7\x45\x9a\x10\xef\x8e\xfa\x27\x55\xbc\x94\x43\x0f\xfd\xbe\x44\xb2\x7f\x69\x07\x31\xe2\x28\xc0\xae\x00\xb8\x0f\xa6\xc5\x12\xce\xc2\x2e\x34\x77\x87\x31\xf9\xa4\xd0\x93\xc7\x77\xf6\x3a\x17\xfa\xc2\x10\xa3\x29\x61\x1a\x04\xb9\x32\x22\xc8\xba\xe9\xfb\xcb\x13\xf5\x72\x9b\x55\xc4\x3a\xfb\x18\x9c\xb7\xcf\xad\x69\x50\xc5\x2d\x05\x27\x34\xfe\x3e\x50\xa8\x5d\x53\xc2\xcb\xf9\x74\x9c\xaa\x9a\x1d\x06\xeb\xaf\xc7\x64\xee\x79\x6c\x4f\x83\xad\x98\xbe\x9d\x45\xce\x9b\x57\x6c\x1c\xad\x43\xee\xe1\xa7\x00\x00\x30\xd6\x3e\x62\xe7\xe9\xfa\x66\xb8\x41\x29\xc7\xe0\xd4\x61\xe8\x2f\xae\x96\x9d\x0b\xdd\x20\xe5\x72\x14\x60\xa6\x9a\xc4\x0e\x4b\x94\x5a\x25\x70\xb1\x5e\xc8\x66\xba\x25\xbb\x41\xd3\x88\xf1\xe2\xea\x17\x25\x93\x62\x46\x8c\x9c\xb1\x48\x25\x13\xd6\x58\xe7\xa2\x19\x72\x3c\xaa\x85\x15\xf8\xba\xb6\xb6\xd6\xf6\xea\x50\x64\x4f\x1f\x61\xea\x7f\xf7\xab\xab\xb3\x09\x70\xbc\xcc\x68\x04\x6d\xea\x98\x91\x6b\x2b\x0d\xb4\x35\x68\x3e\xd7\xc6\xc6\x5b\xaf\x94\x11\xff\x83\x3a\x2f\x56\x8c\xdf\x5f\x00\x00\x00\xff\xff\x0d\x3b\x84\x12\x65\x03\x00\x00")

func _lgraphqlDeploytargetsAdddeploytargetGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsAdddeploytargetGraphql,
		"_lgraphql/deploytargets/addDeployTarget.graphql",
	)
}

func _lgraphqlDeploytargetsAdddeploytargetGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsAdddeploytargetGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/addDeployTarget.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsDeletedeploytargetGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\xd2\x54\xa8\xe6\x52\x50\x48\x49\xcd\x49\x2d\x49\x75\x49\x2d\xc8\xc9\xaf\x0c\x49\x2c\x4a\x4f\x2d\xb1\x82\x8a\xf9\x17\xa4\xe6\x15\x67\x64\xa6\x95\x68\x64\xe6\x15\x94\x96\x58\x81\x95\x2b\x28\x40\x0c\x01\x9b\xc5\xa5\xa0\x50\xab\xc9\x55\xcb\x05\x08\x00\x00\xff\xff\x8e\x88\xf8\xcc\x66\x00\x00\x00")

func _lgraphqlDeploytargetsDeletedeploytargetGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsDeletedeploytargetGraphql,
		"_lgraphql/deploytargets/deleteDeployTarget.graphql",
	)
}

func _lgraphqlDeploytargetsDeletedeploytargetGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsDeletedeploytargetGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/deleteDeployTarget.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsListdeploytargetsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8c\x51\x4a\xc4\x40\x10\x44\xff\x73\x8a\x3e\x87\xbf\xfa\xa1\x08\x12\x44\x0f\x30\xc9\x54\xc6\xc6\x4e\xb7\xf6\xf4\x2c\x84\x25\x77\x5f\x36\x43\xd8\xaf\x7a\x55\x3c\xea\xbf\xc1\x37\xba\x0e\x44\xc2\x35\x5e\xf0\x27\xb6\x7d\x25\x2f\x88\xfa\x44\x49\xe4\xbd\x4d\x70\x45\xa0\xde\x1d\x22\xce\x47\x68\x5a\x71\xc0\xec\x48\x81\x3e\xce\xa6\xd5\x04\xdf\x2e\xbd\x8a\xb5\xfc\x89\xc2\xa6\x8f\x3e\xba\x5d\x38\xc3\x8f\x65\x71\x86\x66\xd9\x3e\xce\xb7\xb0\x5f\x74\xb9\xd6\x9f\x57\xab\x71\xf2\x68\xde\xd9\xad\x05\x7c\x4c\x11\xf0\x6e\x4e\x8d\x25\xbf\xad\xa9\xf4\x8b\xd5\x94\xc3\x9c\xb5\x3c\x9b\x2e\x5c\x06\xa2\x7d\xd8\x6f\x01\x00\x00\xff\xff\x7c\xcb\xe1\x0e\xe7\x00\x00\x00")

func _lgraphqlDeploytargetsListdeploytargetsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsListdeploytargetsGraphql,
		"_lgraphql/deploytargets/listDeployTargets.graphql",
	)
}

func _lgraphqlDeploytargetsListdeploytargetsGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsListdeploytargetsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/listDeployTargets.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsUpdatedeploytargetGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x41\x6e\xea\x30\x10\x86\xf7\x73\x8a\x79\x12\x0b\x90\x38\x41\xb6\xaf\x8b\xd2\x05\xa0\xd2\x1e\xc0\xc5\x43\xb0\x9a\x78\x22\x67\x52\x09\x55\xdc\xbd\x0a\x24\xb1\x27\xce\x2a\xf1\xff\x8f\xa5\xf1\xf7\xd5\x9d\x18\x71\xec\x71\x0d\x88\x2b\x67\x0b\xdc\x79\xf9\xd7\xff\x7b\x53\x53\x81\x27\x09\xce\x97\xdb\x3e\x38\xb3\x6f\xb9\xa2\xcf\x50\xa9\x58\xf8\x9b\xbc\x4a\x02\x77\x42\xe1\x68\x44\x28\xe8\xa6\x6d\xaf\xaf\xdc\xca\x3c\x3b\x72\xd0\xd9\x57\xe7\x2a\xbb\xab\x4d\xa9\x37\xa8\xd9\x3b\xe1\xfe\xf8\x9f\xfd\xc5\x95\x05\xbe\x9d\x0e\xfb\x47\x75\x09\x8e\xbc\xad\x6e\xfb\x6c\xeb\x8a\x3b\x7b\x0c\xfc\xe3\x2c\x85\xbc\x79\xa7\xd2\xf1\xb4\xe4\x06\x7f\x01\x11\xb1\x6b\xac\x11\x7a\xa1\xa6\xe2\xdb\x87\x09\x25\x49\x31\x64\x87\x86\x7c\x7b\x75\x17\x59\x3b\xdf\x74\x52\x0c\x17\x10\x7b\x74\x2b\x67\x87\x53\x63\xe4\x7c\x8d\x25\xe2\x93\xe6\x03\xea\x76\x0a\x53\xa2\x09\xde\x38\x30\xb0\x7d\x32\x9e\xd2\x19\x5f\xcd\x3b\x5e\x9e\x60\x8f\xd8\x55\xf5\x64\x3e\xd2\x8f\x55\x8a\x3e\xf1\x10\x07\x72\x09\x99\x97\x38\xac\xb5\x28\x4b\x09\x06\xad\x48\x2b\x9b\x8d\x8d\xbe\x52\x7b\xc3\xc4\xfd\xf1\xbd\x6f\x12\x25\x10\xe1\x0f\xbf\xe7\x40\x46\x68\x2c\x52\xae\x91\x3f\x28\x7e\xa0\x90\xc1\xc2\xca\x90\xed\x07\xf9\xeb\x61\x8e\x17\x16\x64\xc2\x32\x62\x18\x9f\x77\x87\xbf\x00\x00\x00\xff\xff\x63\x54\xb3\x44\xb0\x03\x00\x00")

func _lgraphqlDeploytargetsUpdatedeploytargetGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsUpdatedeploytargetGraphql,
		"_lgraphql/deploytargets/updateDeployTarget.graphql",
	)
}

func _lgraphqlDeploytargetsUpdatedeploytargetGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsUpdatedeploytargetGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/updateDeployTarget.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xc1\x4e\xc3\x30\x0c\x86\xef\x7d\x8a\x20\x71\xd8\x5e\xa1\x57\xb8\x70\x41\x15\xf0\x02\xa1\xf1\x32\xa3\xcc\x29\x9e\x03\x9a\x50\xdf\x1d\x6d\x59\x92\x25\xad\x96\xe3\xef\xaf\x76\xbf\xff\x10\x44\x0b\x7a\x52\x9b\x4e\x29\xa5\x1e\x27\xf6\x5f\x30\x4a\xaf\x5e\x48\x1e\x62\xf4\x0b\x68\xf7\x31\x89\xc1\x27\x6b\x1a\xf7\x70\xec\xd5\xbb\x30\x92\xbd\x72\x53\x70\x8e\xe1\x3b\xc0\x51\xda\x91\x81\xc9\xf9\xd3\x87\x66\x0b\xd5\xea\xdb\x7c\x88\x97\x07\x2d\x02\x4c\x69\x41\xb7\xfd\xbb\xa0\xda\x98\xe7\x1b\xf8\xc9\xd3\x0e\xed\x06\x69\x0a\xd2\x47\xe2\xfc\xea\x3b\xd5\xfa\xcc\x24\x9d\xab\x57\xce\x8b\x55\x16\xcc\xb3\x5a\xad\x32\x2d\x4c\x6a\x2e\x75\xb8\xfa\x57\xad\xe5\x9d\x0a\x2e\xdf\xcf\xdb\x62\x87\x66\x75\x65\x01\x1a\xe8\xfc\x48\x1f\xa0\x0a\x76\x8c\x40\xc6\x9d\x5e\xdb\xc1\xe8\x7c\x30\x03\xfb\x1f\x34\xc0\xcb\xc9\x1b\x58\xf4\x94\xf3\x79\xd1\xdb\x6a\x59\x4d\xeb\x51\xa9\x9b\xff\x03\x00\x00\xff\xff\x2a\xf3\x98\x73\x76\x02\x00\x00")

func _lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql,
		"_lgraphql/deploytargetconfigs/addDeployTargetConfig.graphql",
	)
}

func _lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargetconfigs/addDeployTargetConfig.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\x84\xf0\x0a\x8a\xf2\xb3\x52\x93\x4b\xa0\x42\x9a\xd5\x60\xd1\x94\xd4\x9c\xd4\x92\x54\x97\xd4\x82\x9c\xfc\xca\x90\xc4\xa2\xf4\xd4\x12\xe7\xfc\xbc\xb4\xcc\x74\x8d\xcc\xbc\x82\xd2\x12\x2b\x88\x22\x10\x00\x99\xa5\x92\x99\x02\xe7\xc3\x4d\x83\x99\x0b\x96\xa9\xd5\xe4\xaa\x05\x04\x00\x00\xff\xff\x66\x8a\x2f\xa2\x86\x00\x00\x00")

func _lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql,
		"_lgraphql/deploytargetconfigs/deleteDeployTargetConfig.graphql",
	)
}

func _lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargetconfigs/deleteDeployTargetConfig.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\xca\xc2\x30\x10\x84\xef\x79\x8a\xfd\xe1\x3f\xe8\x2b\x78\xd4\x53\x2f\x52\xc4\x17\xa8\xcd\x34\x8d\xc4\x6c\xbb\x4d\x94\x20\x7d\x77\x69\x82\xa5\xe6\x90\x65\xbf\xd9\x19\x66\x8c\x90\x44\x3b\x45\xf4\x3f\x08\xdf\xd1\x86\x03\x55\x3e\xfc\xa9\xfd\x5b\x11\x69\x0c\x8e\xd3\xb5\x11\x83\x70\x62\xdf\x59\x33\x1d\x53\x5d\xee\x2a\xbd\xb8\x88\x56\xdb\x37\x40\x11\x65\x33\x91\xd5\x79\x6c\x53\x8a\xb0\x11\x97\xe7\x9b\x07\xd6\xa5\x13\x0b\xaf\x5d\x3a\x6f\x61\xeb\x38\xea\x5a\xf8\x69\x35\xe4\x97\x5e\x60\x2c\xfb\xcc\xe6\xfc\xdf\xa4\xf1\x6d\x8f\xa9\xb4\x8b\xce\x09\xc6\x88\x29\x14\xf0\x82\x35\xfd\xd2\x71\x56\xf3\x27\x00\x00\xff\xff\xba\x92\xa1\x8b\xfd\x00\x00\x00")

func _lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql,
		"_lgraphql/deploytargetconfigs/deployTargetConfigsByProjectId.graphql",
	)
}

func _lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargetconfigs/deployTargetConfigsByProjectId.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\x4e\xf4\x30\x0c\x85\xf7\x39\x45\x7e\x69\x16\x33\x57\xe8\xf6\x67\xc3\x06\x55\xc0\x05\x42\xe3\x49\x8d\x32\x4e\xf0\x38\xa0\x11\xea\xdd\x51\x1b\x28\x71\x19\xf0\xaa\x79\xef\xd5\x7a\xfe\x4e\x45\x9c\x60\x22\xbb\x37\xd6\x5a\xbb\x43\xdf\xd9\x5b\x92\x7f\xf5\xf5\x06\x18\x46\x59\x94\x2a\x3c\xb1\xa3\x61\x84\x73\x67\x1f\x84\x91\x42\x55\x73\x89\x91\xe1\xa5\xc0\x59\x36\x8e\x87\x1c\xd3\xe5\xd1\x71\x80\x76\x4d\x2b\xf7\x9c\x9e\x61\x90\xde\x89\x00\xd3\xfa\xfb\xe1\x7d\x89\x96\xec\x9d\xc0\x4d\x93\xff\x9f\xe8\x88\x61\x8f\x94\x8b\x74\x35\x34\xcf\x5c\x7c\x87\x7e\x7d\x67\x27\xc3\xd8\xf8\xf3\xe8\x36\xaa\x85\xca\x7d\x9d\xfd\x79\xbf\xf2\xbe\x09\xac\x30\x94\xaf\x59\x28\x34\xbf\x76\xd9\x22\xf8\x83\xcf\xba\x63\x5a\xbe\xa6\x43\x4b\xc0\x5c\x5b\xae\x11\x34\xa1\x79\xc8\x9d\x40\x09\x47\x46\x20\x1f\x2f\x77\x5b\x63\x88\xa9\xf8\x9e\xd3\x2b\x7a\xe0\x9f\xce\x3d\x04\x4c\xdb\x76\x2d\x2f\x73\x0d\x90\xd1\xc4\xeb\x49\x66\xfa\x08\x00\x00\xff\xff\x82\xe7\x64\x80\x95\x02\x00\x00")

func _lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql,
		"_lgraphql/deploytargetconfigs/updateDeployTargetConfig.graphql",
	)
}

func _lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargetconfigs/updateDeployTargetConfig.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsAddprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4b\x6a\xc3\x40\x0c\x86\xf7\x3e\x85\x02\x5d\x24\x90\x13\x78\x5b\xba\x08\xed\x22\x50\x7a\x80\x89\x47\x71\xa6\x8c\x35\xae\xac\x31\x94\xd2\xbb\x17\x3f\xe6\x95\x26\x59\x78\xf3\xff\x92\x2d\x7f\x5f\xe7\x45\x89\x71\x04\xdb\x0a\xe0\x89\x54\x87\x35\xbc\x0b\x1b\x6a\x37\xfb\x29\x69\x8d\x7c\xb0\x2d\xb3\xc1\x9f\xce\xce\x6a\xe4\x10\xcf\xa9\xeb\x91\x86\x8b\x39\x4b\x0d\x07\x92\x4d\x99\x1d\xd9\x7d\x62\x23\x47\x25\x82\x4c\xc5\xde\x89\x15\x35\x17\x1c\x8a\xb0\xf7\xd6\x32\x7e\x79\x1c\xe4\xaa\x60\xa7\x7d\x33\x5d\xfc\x42\xa3\x61\x47\x1d\x92\x94\xe7\x29\x2f\xee\xa0\x2d\xce\x77\x2c\x07\x8b\x63\xd5\xe2\xb3\xb2\x4d\x0a\x35\x8e\x68\x5d\x3f\xed\x67\xaf\x1a\xde\x4c\x67\x24\x4d\xf5\x6c\x46\x25\xf8\x8a\xdf\xe1\x23\x3b\xf8\xa9\x00\x00\x94\xd6\xeb\x5f\x6d\x0d\xf5\x5e\xea\x35\x07\x58\x28\xce\x30\xf7\x6b\x14\x30\xae\x3c\x43\x9c\x91\x4c\x54\x43\x99\x01\x4d\x20\xff\x95\xd7\x64\xef\x31\x0f\x8b\x09\x77\x24\x1f\xaa\x12\x7a\xe1\x20\x8e\xdc\xc6\x7f\x5b\x4b\x58\x4a\x42\xa2\x9b\x08\x20\x37\x93\x7b\x0a\x03\x8f\x2d\x3d\x94\x98\x6e\x4e\x0a\x33\x9f\x73\xfb\xbb\x8b\xd6\x8c\xce\xf4\x2d\x65\x35\x3d\x7f\x01\x00\x00\xff\xff\x70\x99\xe2\x23\x22\x03\x00\x00")

func _lgraphqlProjectsAddprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsAddprojectGraphql,
		"_lgraphql/projects/addProject.graphql",
	)
}

func _lgraphqlProjectsAddprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsAddprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/addProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsMinimalprojectbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x4a\xc6\x30\x10\x85\xf7\x39\xc5\x08\x2e\xea\x15\x5c\x0a\x2e\x04\x91\x82\x78\x80\xd8\x8c\xed\x48\x32\x49\x27\x93\x42\x91\xde\x5d\xda\x62\x52\x17\x7f\x36\x49\xbe\x37\x2f\x79\x6f\x2e\x28\x2b\x74\x06\xe0\x9e\x6d\xc0\x47\x78\x57\x21\x1e\xef\x1e\xe0\xc7\x00\x00\x24\x89\xdf\x38\xe8\xd3\xfa\x66\x03\x76\x07\x02\x38\x27\x0f\xc3\xdf\xdc\xbe\xc8\xd5\xe3\x2e\xd5\x8b\x2d\x1a\x5f\x9c\x6f\xe0\x53\x2c\x0f\x13\xe6\x0a\x52\xf1\x5e\x70\x2e\x98\xf5\x02\x25\xba\x32\x28\x45\x7e\xe6\x85\x24\x72\x40\xd6\xaa\xc6\x84\x9c\x27\xfa\xd2\xfe\x4c\xd8\x5b\x55\x14\xae\xba\xc3\x05\x7d\x4c\xbb\xe7\x62\xcf\xaf\x14\xa8\x3d\x32\x92\x7e\x88\xbf\x1d\xb4\x7e\xd2\x4a\xfe\xab\xb9\x99\xb6\x6f\xe6\x37\x00\x00\xff\xff\x43\xb8\x3d\x3c\x4c\x01\x00\x00")

func _lgraphqlProjectsMinimalprojectbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsMinimalprojectbynameGraphql,
		"_lgraphql/projects/minimalProjectByName.graphql",
	)
}

func _lgraphqlProjectsMinimalprojectbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsMinimalprojectbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/minimalProjectByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsProjectbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\xcd\x6e\xdb\x30\x0c\xbe\xe7\x29\x38\x60\x87\xee\x92\x07\xd8\x6d\x2d\x8a\x6d\xe8\x16\x14\x6d\xd6\x6b\xc1\xc8\x74\xac\x45\x16\x5d\x92\xf6\x60\x0c\x7d\xf7\xc1\x09\xe6\x5f\x35\xeb\x61\x40\x75\x89\xf2\x7d\xf4\x47\x8a\x7f\x4f\x35\x49\x0b\x17\x2b\x80\xf7\x11\x4b\xfa\x08\xf7\x26\x3e\xee\xdf\x7d\x80\xdf\x2b\x00\x80\x4a\xf8\x27\x39\xbb\x6c\x37\x58\xd2\xc5\x11\x02\x38\x59\x1e\x3f\xf8\x6b\xd7\x1d\x9f\xf5\xd7\x8e\xea\xff\x60\x6d\xfc\x35\x0b\x03\xb0\x13\x8c\xae\x20\xed\x81\xaa\x0e\x41\xe8\xa9\x26\xb5\x11\x28\xbe\x41\xa3\x1b\x6a\x47\x10\x67\xb5\x33\xcf\xf1\x3a\x36\x5e\x38\x96\x14\xad\x67\xd5\x58\x70\x4f\x57\x18\x5c\x8f\x71\x45\x51\x0b\x9f\xdb\xed\xe9\x21\xb7\x68\x46\x12\x7b\x3e\xa3\x86\x02\x57\x9d\xce\x48\x52\xbf\xf9\xd2\x0f\xc2\x7b\x6f\x3f\x24\xbc\xfc\x9e\xbd\x70\x5d\xe9\x90\x08\x80\xf5\x7a\x0d\x1c\xe1\x73\x47\xc0\x98\x00\x78\x7c\xb4\xb6\xa2\x49\x82\x16\x19\xeb\x4e\x49\xe5\x8e\x44\xa7\x1f\x03\xd4\x4a\x32\xc7\x00\xa8\x44\x1f\x16\xa8\x6a\x71\x43\xed\x42\x22\xe9\xee\x74\x0e\xd4\x6e\xdb\xea\x05\xe6\x01\x43\xbd\xa4\x9e\x17\x48\xee\x45\x6d\x93\x72\x10\x30\x49\xcc\x25\x84\xc3\xd4\x64\x6c\x30\xdc\x87\x5b\x64\xf3\xb9\x77\xd8\x75\x86\x42\xa2\x10\x9b\x91\xc1\x7d\x40\x77\x78\x5d\x51\x7e\xd1\xae\x60\x3e\x9c\x2f\x94\x2b\x30\x46\x0a\xc9\x10\x93\x01\xdc\xb1\x3b\x90\x5d\x15\x68\x6f\x19\xc5\x75\xd7\x32\xaf\x0b\xe0\xd8\x5d\x9f\xb2\x4c\x48\xf5\x5c\x14\xff\xf0\xf8\xdd\x3b\x61\xe5\xdc\xb6\x84\xa5\xfe\xb7\xb7\xa7\x1a\xa2\x1f\xfb\xb1\x93\xd1\x7e\x1a\x2c\x29\x36\x0f\x28\x1e\x77\x81\xa6\x21\xcd\xbc\xa8\xe3\xc9\x5c\x34\x93\x61\x98\xe8\xf5\x5b\xe4\x9c\x5e\x46\x55\xe0\xf9\xb0\x9d\xc0\x4b\x54\xba\xa3\x7c\x81\x7f\x21\xcc\x52\xf8\xd6\xdb\x64\x5e\x46\x21\xcc\xf4\xe7\xdb\x70\x36\x8b\x8b\xc5\x76\x26\x3d\x89\x16\x9c\xa7\x68\x9e\xa4\x54\xa9\xba\xdf\xe7\xd5\x9f\x00\x00\x00\xff\xff\x0a\x9a\xaa\x65\x86\x06\x00\x00")

func _lgraphqlProjectsProjectbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsProjectbynameGraphql,
		"_lgraphql/projects/projectByName.graphql",
	)
}

func _lgraphqlProjectsProjectbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsProjectbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/projectByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsProjectbynamemetadataGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x00\x0b\x29\x28\x40\x54\x82\x35\xc0\xd4\x81\x40\x66\x0a\x9c\x09\x92\x82\x73\x72\x53\x4b\x12\x53\x12\x4b\x12\xa1\x02\xb5\x5c\x20\x0c\x08\x00\x00\xff\xff\xd4\x75\x13\x1d\x79\x00\x00\x00")

func _lgraphqlProjectsProjectbynamemetadataGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsProjectbynamemetadataGraphql,
		"_lgraphql/projects/projectByNameMetadata.graphql",
	)
}

func _lgraphqlProjectsProjectbynamemetadataGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsProjectbynamemetadataGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/projectByNameMetadata.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsProjectsbymetadataGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xcf\x4a\x04\x31\x0c\xc6\xef\xfb\x14\x15\x3c\xb8\xaf\xb0\x47\xc1\x83\xa0\xb0\x20\x3e\x40\x9c\xc6\xdd\x68\x9b\x76\xd3\x74\x60\x58\xe6\xdd\xa5\xd3\xb1\x1d\x91\xcd\x29\xfc\x92\xef\xcb\x9f\x4b\x46\x99\xcc\xc3\xce\x18\x63\xee\xbf\x71\x3a\x98\x37\x15\xe2\xd3\x5d\x25\x23\xb8\x8c\xbf\x6c\x7f\x5d\x60\x94\xf0\x85\x83\xa6\xc7\xe9\x15\x15\x2c\x28\x54\x79\x09\xbf\x92\x83\xb9\x36\x56\x62\x71\x2e\xfe\x7f\xe8\xea\x5e\xa7\xb4\xca\xbc\x64\xfb\x6e\x40\xb6\xa5\x0c\xbe\x37\x42\xd6\xf0\x6c\x5d\x07\x1f\x02\x3c\x9c\x31\x35\x10\xb3\x73\x82\x97\x8c\x49\xd3\xbf\x1d\x7b\x97\x04\x9b\x07\xa5\xc0\x4f\x3c\x92\x04\xf6\xc8\xda\xaa\x21\x22\xa7\x33\x7d\xea\xb1\xde\x7d\x04\x55\x14\x6e\x75\x8b\x23\xba\x10\x8b\x66\x23\x4f\x2f\xe4\xa9\x9b\x9c\x48\xdf\xc5\xdd\xde\xbc\x0d\xd9\xbe\x6d\x73\x77\xfd\xc9\xbc\x9b\x7f\x02\x00\x00\xff\xff\xf0\x75\x45\x7d\xb0\x01\x00\x00")

func _lgraphqlProjectsProjectsbymetadataGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsProjectsbymetadataGraphql,
		"_lgraphql/projects/projectsByMetadata.graphql",
	)
}

func _lgraphqlProjectsProjectsbymetadataGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsProjectsbymetadataGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/projectsByMetadata.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsRemoveprojectmetadatabykeyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\x84\xf0\xb2\x53\x2b\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x15\x35\xab\xc1\x62\x45\xa9\xb9\xf9\x65\xa9\x01\x45\xf9\x59\xa9\xc9\x25\xbe\xa9\x25\x89\x29\x89\x25\x89\x4e\x95\xde\xa9\x95\x10\x13\x40\x20\x33\xaf\xa0\xb4\xc4\xaa\x1a\xce\x07\x8b\xa5\x58\x81\x4c\x47\x11\x03\x1b\x0f\xb2\x04\x2e\x5a\x0b\x66\x69\x22\xb4\x22\xe9\xc8\x4b\xcc\x4d\x85\x73\x72\xa1\x36\x73\x41\x74\xd5\x02\x02\x00\x00\xff\xff\xfc\x0e\xa7\xd1\xc8\x00\x00\x00")

func _lgraphqlProjectsRemoveprojectmetadatabykeyGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsRemoveprojectmetadatabykeyGraphql,
		"_lgraphql/projects/removeProjectMetadataByKey.graphql",
	)
}

func _lgraphqlProjectsRemoveprojectmetadatabykeyGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsRemoveprojectmetadatabykeyGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/removeProjectMetadataByKey.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsUpdateprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\x84\xf0\x0a\x12\x4b\x92\x33\xac\x14\x42\x0b\x52\x12\x4b\x52\x03\x8a\xf2\xb3\x52\x93\x4b\x02\x40\x62\x9e\x79\x05\xa5\x25\x8a\x9a\xd5\x60\x65\xa5\xc8\xd2\x10\x73\x40\x20\x13\xa4\xc6\x4a\xa1\x1a\x2e\x00\x16\x4c\xb1\x02\x59\x82\x22\x06\xb5\x05\x62\x1b\x5c\xa6\x16\xcc\xd2\x44\x68\x47\xd2\x95\x97\x98\x9b\xca\x05\x51\x54\x0b\x08\x00\x00\xff\xff\x58\x7f\xbf\xd7\xc2\x00\x00\x00")

func _lgraphqlProjectsUpdateprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsUpdateprojectGraphql,
		"_lgraphql/projects/updateProject.graphql",
	)
}

func _lgraphqlProjectsUpdateprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsUpdateprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/updateProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsUpdateprojectmetadataGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x4d\xca\x02\x31\x0c\x86\xf7\x3d\x45\x3e\xf8\x16\xce\x15\x7a\x03\x17\x82\xe0\x09\xc2\x34\x68\xd5\x66\xca\xf0\x56\x90\x61\xee\x2e\xfd\xd1\xb1\x98\x55\xf3\xa4\xc9\xfb\x84\x04\x86\x9f\x94\x76\x86\x88\xe8\xdf\x3b\x4b\x7b\xc5\x5f\xed\x6e\xf2\xb4\x74\xc2\xec\xf5\xdc\xc8\x83\xef\x49\xde\x6c\x58\x0a\x4c\xd1\x31\xe4\x38\x4f\x57\x19\x71\x10\xb0\x63\x70\xbd\x97\xcb\x6b\x4c\xb0\xcb\xa7\x2f\xcc\xd9\x9c\xd5\xb1\xc8\x18\x2f\x96\xfa\x8f\xb9\x8a\x45\x76\xf9\x99\x34\x9b\x6a\xd5\x4d\x57\xd3\xbf\x86\xed\xec\x57\xac\x72\xd8\xd6\x42\x33\x37\x75\x6b\x7d\x05\x00\x00\xff\xff\xd8\xa2\xd3\x0b\x1b\x01\x00\x00")

func _lgraphqlProjectsUpdateprojectmetadataGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsUpdateprojectmetadataGraphql,
		"_lgraphql/projects/updateProjectMetadata.graphql",
	)
}

func _lgraphqlProjectsUpdateprojectmetadataGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsUpdateprojectmetadataGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/updateProjectMetadata.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsAddorupdateenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4d\x6a\xc4\x30\x0c\x85\xf7\x39\x85\x0a\x5d\xcc\xc0\x9c\x20\xcb\xd2\x81\x76\xd3\x96\xfe\x1c\xc0\xc4\x9a\xd6\x25\x91\x8d\xad\x04\x86\x92\xbb\x97\xc6\xb1\x23\x27\xa1\x9b\x20\xe9\x29\xd2\xd3\xe7\xae\x67\xc5\xc6\x12\x1c\x2a\x80\x5b\x52\x1d\xd6\xf0\xc6\xde\xd0\xe7\xcd\xe9\xaf\xe2\xbc\xfd\xc6\x86\x6b\x78\x24\x8e\x15\x8d\xae\xb5\xd7\xf7\xab\xc3\x1a\xee\x73\x2c\xb5\x3b\x15\xf0\x15\x2f\xe5\xa0\x28\x3d\xa0\xd2\x42\x92\x03\x0d\xb7\x58\xd4\x91\x06\xe3\x2d\x75\x48\x1c\xb7\x9d\x69\x58\x56\x59\x87\x14\xbe\xcc\x85\x5f\xa2\xc3\x27\x69\xfd\x08\x3f\x15\x00\x80\xd2\xfa\xd9\x7f\x38\xad\x18\xcf\xcb\xb4\x83\x21\xd7\x73\x3d\xf7\x00\xc4\xab\xa7\xe3\x4f\x73\x29\x9f\x9d\x00\x24\x41\x5e\x2f\x50\x94\x72\x06\x50\x02\x29\x9b\x32\x8a\x12\xcd\x6a\x51\xa4\x22\x19\xa5\x86\x0d\x9e\x35\xb0\xd4\xb8\x4f\x6a\x17\xe0\xf4\xc7\x78\xcc\x64\x8c\x16\x88\xe6\xd0\xdb\x9e\x8b\x38\x6c\xd0\xec\x3b\xfc\xc7\xce\x2c\xf5\xd3\x4b\xa5\x9d\x8d\x47\x91\x69\x6c\x71\xc9\x02\xfa\xc1\x34\x18\x92\xd1\xc2\xec\xca\xf0\x58\xa5\xef\x58\xfd\x06\x00\x00\xff\xff\x6d\xb2\xb5\x91\xef\x02\x00\x00")

func _lgraphqlEnvironmentsAddorupdateenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsAddorupdateenvironmentGraphql,
		"_lgraphql/environments/addOrUpdateEnvironment.graphql",
	)
}

func _lgraphqlEnvironmentsAddorupdateenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsAddorupdateenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/addOrUpdateEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x3f\x0a\xc2\x30\x14\xc6\xf7\x9e\xe2\x13\x1c\xea\x15\xb2\x8a\x83\x93\x43\xe9\x01\x22\x79\x94\x07\xe6\x25\x24\xaf\x82\x48\xee\x2e\xa9\xc1\xd6\xa1\xeb\xf7\xf7\xe7\x67\xb5\xca\x41\xd0\x77\xc0\x91\xe4\xc9\x29\x88\x27\x51\x83\xab\xe8\xa1\x8a\x91\x52\xe6\xac\x24\x3a\x68\x48\x76\xa2\xf3\xc3\xb2\x37\x18\x34\xb1\x4c\x4b\xe4\xfe\x52\xca\x63\x26\xf7\x6b\x9d\xf0\xee\x00\xc0\x3a\x77\x4b\x63\x74\x56\xe9\xb2\x8e\xb7\xa1\x9e\x25\xce\x6a\x5a\x14\xf8\xbb\xdf\xc2\x34\x7f\x8f\x64\x07\xb1\xb5\x36\x70\x2b\xe8\xe2\x95\x0a\xfa\x3d\xe7\xaa\x94\xae\x7c\x02\x00\x00\xff\xff\x68\x1d\xbe\x72\x10\x01\x00\x00")

func _lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql,
		"_lgraphql/environments/addOrUpdateEnvironmentStorage.graphql",
	)
}

func _lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/addOrUpdateEnvironmentStorage.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsAddrestoreGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x49\x4a\x4c\xce\x2e\x2d\xc8\x4c\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x4c\x49\x09\x4a\x2d\x2e\xc9\x2f\x4a\xd5\xc8\xcc\x2b\x28\x2d\xb1\x52\xa8\x86\xa8\xf5\x4c\xb1\x42\x68\xab\x85\x29\x57\x50\xc8\x4c\x01\x33\x6a\xb9\x40\x18\x10\x00\x00\xff\xff\xdc\x18\x6c\xe7\x65\x00\x00\x00")

func _lgraphqlEnvironmentsAddrestoreGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsAddrestoreGraphql,
		"_lgraphql/environments/addRestore.graphql",
	)
}

func _lgraphqlEnvironmentsAddrestoreGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsAddrestoreGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/addRestore.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4d\x6e\x84\x30\x0c\x85\xf7\x39\xc5\x43\xea\x02\x24\x4e\xc0\xb2\x3b\xa4\xaa\x9b\x9e\x20\x0d\x56\x95\xb6\x24\x8c\xe3\x8c\x84\x10\x77\x1f\xf1\x33\x99\xc0\xcc\xca\xce\x67\xbf\xcf\xb9\x44\xe2\x11\xa5\x02\xde\x9c\xee\xa9\xc1\x97\xb0\x75\x3f\x45\xbd\x90\x81\xfd\x2f\x19\x69\xd0\x3a\x29\x2a\x4c\x0a\x00\xc8\x5d\x2d\x7b\xd7\x93\x93\xf7\xf1\x53\xf7\x54\xae\x18\xd8\xf2\xab\xa6\xde\x51\x12\xdc\x55\xd5\x3e\x98\xf6\x0a\xd8\x2e\xb5\x4b\x32\x3d\xbe\xb5\xf9\x8b\x43\x78\x2c\x1e\x56\x81\xe0\x23\x1b\xca\xc0\x16\x68\xf3\x1d\xc3\xa4\x85\x72\xd2\xd1\x3f\x1d\x09\x53\x10\xcf\x94\xdf\x39\x5d\x02\x82\x68\x89\xe1\x80\x9e\xd5\x2f\x7f\x90\xfc\x1f\xde\x68\xb1\xde\x65\xb3\x59\x9d\xbb\xa5\xce\xea\x16\x00\x00\xff\xff\x0c\x15\x4e\x54\x93\x01\x00\x00")

func _lgraphqlEnvironmentsBackupsforenvironmentbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql,
		"_lgraphql/environments/backupsForEnvironmentByName.graphql",
	)
}

func _lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsBackupsforenvironmentbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/backupsForEnvironmentByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsDeleteenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x08\x2e\xec\x15\xb2\x14\x3c\x81\x27\x08\xf5\x23\x91\x66\xa6\x84\x1f\x11\x4a\xef\x2e\x4d\x68\xa0\xfd\xab\xf0\x1e\xe4\x4d\x2a\x0c\x8c\xa6\x72\x73\x22\x22\x57\x0d\x09\x5e\x9e\xcc\x51\xdf\x97\x86\xe6\x6c\x1f\x8c\x3c\x51\xfc\x30\x16\xc2\xcb\xdd\x6c\x42\x50\x37\x2c\x55\xbc\x30\x81\x78\xe8\x37\x66\xd3\x04\x65\xfb\x78\x5b\xd4\xb9\xd0\xcb\xd2\xc1\xb6\x16\xac\xdd\x03\xef\xd5\xbd\x7f\xb0\xbd\xbe\xdf\xd1\xed\x5a\x5f\x83\x5b\xff\x01\x00\x00\xff\xff\x02\x42\x6d\x97\xda\x00\x00\x00")

func _lgraphqlEnvironmentsDeleteenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsDeleteenvironmentGraphql,
		"_lgraphql/environments/deleteEnvironment.graphql",
	)
}

func _lgraphqlEnvironmentsDeleteenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsDeleteenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/deleteEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsEnvironmentbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4d\x0a\xc2\x30\x10\x85\xf7\x39\xc5\x14\x5c\x28\xf4\x04\x5d\xba\x73\x53\x04\xbd\x40\x69\x9e\x1a\x69\x93\x38\x99\x16\x82\xf4\xee\xd2\x84\x86\xba\x18\xf8\x78\xc3\xfb\xf9\x4c\xe0\x48\x47\x45\x74\xb0\xdd\x88\x86\x6e\xc2\xc6\x3e\xab\x7a\x55\x3c\xbb\x37\x7a\x69\xe8\x62\xa5\x52\xa7\xaf\x22\x82\x9d\x0d\x3b\x3b\xc2\xca\x39\xb6\xdd\x88\xd5\x4a\x94\xbd\x29\xa2\x4e\x42\xb1\x6e\x21\x8a\x28\x05\x10\x19\x5d\x2c\x09\xd8\x4d\xb2\xa3\x90\x50\xc3\x0f\x2e\xde\xa3\xcf\x9f\x5d\x6d\xd1\x9c\x87\x0d\x2f\xf3\x90\x6b\x6e\x68\xb7\xc0\xc9\xeb\x4e\x90\x5b\x7a\x46\x61\x8d\x01\x1b\x07\xf0\x6c\x7a\x84\x3c\x69\x37\xeb\x6f\xda\xa2\xd6\x5b\xd4\x2f\x00\x00\xff\xff\x5b\x2b\x68\x2b\x26\x01\x00\x00")

func _lgraphqlEnvironmentsEnvironmentbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsEnvironmentbynameGraphql,
		"_lgraphql/environments/environmentByName.graphql",
	)
}

func _lgraphqlEnvironmentsEnvironmentbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsEnvironmentbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/environmentByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsEnvironmentbynamespaceGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xc1\x8a\xc3\x20\x10\x86\xef\x3e\xc5\x2c\xec\x61\xf7\x15\x72\xdc\xeb\x42\x28\xb4\x2f\x90\xea\xdf\xd6\x36\x51\x3b\x6a\x40\x4a\xde\xbd\x68\x88\x18\xe8\x41\xfc\x98\x7f\x98\xff\x7b\x46\x70\xa2\x1f\x41\xf4\x6d\x86\x09\xde\x0d\x12\x1d\x1d\x03\x6b\x73\xfd\x12\xbf\x2f\x41\x04\x33\x6b\xb6\x66\x82\x09\x7f\xa9\xaf\x4b\xbb\xf1\x7f\x3c\x83\x0d\x02\x7c\x5d\xc8\x90\xef\x12\x3d\x3e\x87\x5d\x53\x29\x88\x4a\x17\x91\x56\xe5\xcb\x49\x01\xb6\x31\x34\xe4\x0b\x2a\xb8\xd1\xa6\x53\x72\x6b\xd2\xa8\xd4\x99\x75\x30\xfe\xa6\x2f\xe1\xc0\xf6\x0e\x19\xfa\xed\x60\x74\x6a\x08\x58\x5b\x24\xa3\xb2\xc2\x88\x8d\x3d\x78\xd6\x12\x7e\x55\x6a\xb4\x76\x6a\x8b\xc8\x6f\x11\xef\x00\x00\x00\xff\xff\x81\x48\x8a\x83\x44\x01\x00\x00")

func _lgraphqlEnvironmentsEnvironmentbynamespaceGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsEnvironmentbynamespaceGraphql,
		"_lgraphql/environments/environmentByNamespace.graphql",
	)
}

func _lgraphqlEnvironmentsEnvironmentbynamespaceGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsEnvironmentbynamespaceGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/environmentByNamespace.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsSetenvironmentservicesGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\x84\xf0\x8a\x53\x8b\xca\x32\x93\x53\x8b\xad\x14\xa2\x83\x4b\x8a\x32\xf3\xd2\x63\x15\x35\xab\xc1\x52\xc5\xa9\x25\xae\x79\x65\x99\x45\xf9\x79\xb9\xa9\x79\x25\xc1\x50\x75\x10\x43\x40\x20\x33\xaf\xa0\xb4\xc4\x4a\xa1\x1a\x2e\x00\x02\xa9\x08\x1d\x56\x20\xab\x50\x24\x11\x76\xc1\xad\x85\xcb\xd7\x82\x59\x9a\x08\xd3\x90\xf4\xe6\x25\xe6\xa6\x72\x41\x14\xd5\x02\x02\x00\x00\xff\xff\xfc\x33\x0b\xcb\xce\x00\x00\x00")

func _lgraphqlEnvironmentsSetenvironmentservicesGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsSetenvironmentservicesGraphql,
		"_lgraphql/environments/setEnvironmentServices.graphql",
	)
}

func _lgraphqlEnvironmentsSetenvironmentservicesGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsSetenvironmentservicesGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/setEnvironmentServices.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsUpdateenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\x41\x6e\x86\x20\x10\x85\xf7\x9e\x62\x4c\xba\xa8\x57\x60\xdf\x85\x9b\xc6\x45\x7b\x00\x02\xd3\x48\xa3\x03\x81\xc1\xc4\x18\xef\xfe\x07\xf9\x03\x22\x2b\xbe\xc7\x1b\xde\xbc\x35\xb2\x64\x63\x09\x3e\x3b\x00\x80\x0f\xa3\x05\x8c\xc4\x7d\x26\x27\x59\xcd\x02\x7e\x9d\x96\x8c\x5f\xb4\x19\x6f\x69\x45\xe2\x29\xe9\x23\xb9\xc8\xfd\x70\x5c\xd6\xf8\xb4\xe4\xff\xd2\x31\xc9\x27\xe0\x28\xc2\x25\x6a\x91\xc2\x1a\xed\x9d\x96\x53\xcb\xcb\x79\xdd\x86\x3a\x7e\x9b\x22\xb9\x62\x01\x6f\x23\x3f\x28\x14\xd4\xe8\x16\xbb\xff\xec\xae\x3a\xb0\x2e\xdb\xe8\xd6\x21\x85\xd9\xfc\xf1\xe4\xed\x3f\x2a\xfe\xbe\x87\xe4\x9e\x75\x03\xe5\xb1\x61\x8d\x0b\xde\x39\xa0\xdf\x8c\xc2\xf0\x6c\xdf\x60\x53\x23\xf7\x3d\xbb\xf3\x15\x00\x00\xff\xff\x5a\xa6\xf7\x47\x9b\x01\x00\x00")

func _lgraphqlEnvironmentsUpdateenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsUpdateenvironmentGraphql,
		"_lgraphql/environments/updateEnvironment.graphql",
	)
}

func _lgraphqlEnvironmentsUpdateenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsUpdateenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/updateEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlTasksSwitchactivestandbyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\x52\x50\xd0\x54\xa8\xe6\x52\x50\x28\x2e\xcf\x2c\x49\xce\x70\x4c\x2e\xc9\x2c\x4b\x0d\x2e\x49\xcc\x4b\x49\xaa\x04\x69\x50\x50\xc8\xcc\x2b\x28\x2d\xb1\x02\xab\x01\x01\x98\x01\x30\xbe\x82\x42\x5e\x62\x6e\xaa\x15\xdc\x64\xa8\x70\x2d\x17\x8c\xd4\x84\x6a\xcd\x4c\xe1\x02\x09\xd4\x02\x02\x00\x00\xff\xff\x76\xc4\x5e\xef\x90\x00\x00\x00")

func _lgraphqlTasksSwitchactivestandbyGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlTasksSwitchactivestandbyGraphql,
		"_lgraphql/tasks/switchActiveStandby.graphql",
	)
}

func _lgraphqlTasksSwitchactivestandbyGraphql() (*asset, error) {
	bytes, err := _lgraphqlTasksSwitchactivestandbyGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/tasks/switchActiveStandby.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlTasksUpdatetaskGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x4d\x0e\x02\x21\x0c\x85\xf7\x73\x8a\x4e\xe2\xc2\xb9\x02\x37\x60\xe7\x42\x0f\xd0\x40\x13\x89\xf2\x13\x28\xab\x09\x77\x37\x80\xa9\xc8\x8a\xef\x7b\xed\xab\xaf\x8c\xec\x62\x80\xeb\x06\x00\x70\x71\x56\x81\x0e\xbc\x4f\x4a\xc8\xe6\xa9\xe0\x91\x2c\x32\xdd\xb1\xbc\x6e\x5d\xe8\x90\x2a\xef\xc7\x39\x66\xaa\x64\xb3\xa1\x3f\xd7\x07\x14\x9c\x22\x86\xb4\xaa\xd7\xff\xb9\x6f\xff\xbc\x23\x49\x1b\xbf\xe3\xb7\xbe\x6c\x05\xf4\x24\x50\x18\xb9\x16\x41\x93\x09\x99\xec\x1a\xe7\x95\x4d\xf4\xe9\x4d\xab\xc9\xe4\x23\x93\x9e\xa2\x6d\xed\x13\x00\x00\xff\xff\x25\xa0\x63\x9d\x0e\x01\x00\x00")

func _lgraphqlTasksUpdatetaskGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlTasksUpdatetaskGraphql,
		"_lgraphql/tasks/updateTask.graphql",
	)
}

func _lgraphqlTasksUpdatetaskGraphql() (*asset, error) {
	bytes, err := _lgraphqlTasksUpdatetaskGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/tasks/updateTask.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddgroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x4c\x49\x71\x2f\xca\x2f\x2d\xd0\xc8\xcc\x2b\x28\x2d\xb1\x82\x8a\x2a\x28\x40\x14\x83\xf5\x80\x45\x6a\x35\xe1\x52\x99\x29\x48\x6a\x20\x92\x5c\x20\x0c\x08\x00\x00\xff\xff\x33\xcc\xe8\xad\x6e\x00\x00\x00")

func _lgraphqlUsergroupsAddgroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddgroupGraphql,
		"_lgraphql/usergroups/addGroup.graphql",
	)
}

func _lgraphqlUsergroupsAddgroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddgroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddgroupstoprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x80\x30\x3c\xf3\x0a\x4a\x4b\x14\x75\x40\x52\xe9\x45\xf9\xa5\x05\xc5\x56\x0a\xd1\xee\x20\x06\x44\x22\x56\x51\x53\xa1\x9a\x4b\x41\x41\x41\x21\x31\x25\x05\x2c\x5e\x1c\x92\x0f\xd5\xaa\x91\x09\x52\x62\x05\x95\x57\x50\x80\x1b\x0d\xb3\x04\x2a\x0e\x33\x17\x6a\x01\x58\xb4\x56\x13\xae\x2d\x33\x05\xca\xc8\x4b\xcc\x4d\x85\x48\x72\x81\x30\x20\x00\x00\xff\xff\x8b\x70\xd6\xd3\xb8\x00\x00\x00")

func _lgraphqlUsergroupsAddgroupstoprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddgroupstoprojectGraphql,
		"_lgraphql/usergroups/addGroupsToProject.graphql",
	)
}

func _lgraphqlUsergroupsAddgroupstoprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddgroupstoprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addGroupsToProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddsshkeyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xc1\x0a\xc2\x30\x10\x44\xef\xf9\x8a\x11\x3c\xb4\xe0\x17\xf4\xee\xc9\xa3\xe2\x7d\x21\x41\x17\x9b\x58\xda\xe4\x50\x24\xff\x2e\xe9\xa6\x59\xc1\x53\xc2\x9b\x61\xdf\xf8\x14\x29\xf2\x3b\xa0\x33\xc0\x31\x90\x77\x03\xae\x71\xe6\xf0\x38\x9c\x0a\x79\xb9\xf5\x4e\x63\xfa\xa7\xb7\x75\x2a\x70\x79\x5e\xe4\x2f\x41\x5a\xdc\x7c\xf6\xc4\x63\xeb\xf7\xf8\x18\x00\x20\x6b\xa5\xdc\x71\x98\x52\x1c\x2a\x06\xc4\xb9\xa9\x2b\x51\x67\xd3\x6b\x22\xde\x7d\x41\xe5\x45\xab\x17\x01\x27\x13\x74\x4e\x4d\xf2\xf6\xe6\xbe\x55\xd9\xfe\xac\x30\x7b\x25\x9b\x6f\x00\x00\x00\xff\xff\xbf\xff\xb7\x89\x17\x01\x00\x00")

func _lgraphqlUsergroupsAddsshkeyGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddsshkeyGraphql,
		"_lgraphql/usergroups/addSshKey.graphql",
	)
}

func _lgraphqlUsergroupsAddsshkeyGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddsshkeyGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addSshKey.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAdduserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x0e\xc2\x30\x0c\x85\xe1\x3d\xa7\x78\x48\x0c\xad\xc4\x09\x7a\x08\x16\xc4\x01\x2c\x12\x90\xa5\xda\x45\xa9\x3b\x21\xee\x8e\xda\x3a\x69\x06\x06\x2f\xff\x1b\x3e\xcb\x62\x64\x3c\x29\xba\x00\x9c\x93\x10\x8f\x03\x6e\x96\x59\x5f\xa7\xcb\x9a\x9e\x9c\x67\xbb\x92\xa4\x92\xb7\x3a\xd2\x9f\xf8\x98\x44\x92\x5a\x69\x3d\x3e\x01\x00\x28\xc6\xfb\x9c\x72\xc7\xfa\x5e\x6c\xf0\x08\xb8\xb5\x9b\xde\x1a\xec\x80\x7d\x3b\xc8\xaa\xfb\x52\xdd\xf2\xc1\xd6\xbf\x7d\xa5\x38\xb6\xe6\xbe\x86\xf5\x7e\x01\x00\x00\xff\xff\xf7\xf8\x9f\x0d\xfe\x00\x00\x00")

func _lgraphqlUsergroupsAdduserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAdduserGraphql,
		"_lgraphql/usergroups/addUser.graphql",
	)
}

func _lgraphqlUsergroupsAdduserGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAdduserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addUser.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddusertogroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xbd\x0a\xc3\x30\x0c\x84\xf7\x3c\xc5\x05\x3a\x24\xd0\x27\xf0\x5e\xba\x75\xe8\xcf\x03\x18\x6c\x82\x20\xb6\x83\x62\x4f\x21\xef\x5e\xa4\xb8\xf1\xd2\x41\x20\xee\x8e\xfb\x2e\x94\x6c\x33\xa5\x88\xa1\x03\x2e\x65\xf5\x7c\x0b\x96\x66\x83\x57\x66\x8a\x53\x7f\x15\x79\xe2\x54\x96\x87\x0d\xfe\x8f\xfc\x4c\xb3\x37\xb8\xff\xde\x7e\xc4\xd6\x01\x80\x75\xee\xb3\x7a\x7e\x27\xb5\x06\x8a\x4b\xc9\xa6\x7a\x80\x80\x0c\x36\xf8\x03\xd6\xc0\xd8\x6b\x42\xcb\x25\x12\x95\xdb\x36\x9c\x09\x56\x72\x5b\xa1\xf2\x3e\x9e\x0c\x72\xf5\x91\x86\xc3\xec\xe4\xbe\x01\x00\x00\xff\xff\x00\xb0\x61\x57\xf3\x00\x00\x00")

func _lgraphqlUsergroupsAddusertogroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddusertogroupGraphql,
		"_lgraphql/usergroups/addUserToGroup.graphql",
	)
}

func _lgraphqlUsergroupsAddusertogroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddusertogroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addUserToGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAllgroupmembersGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x71\x2f\xca\x2f\x2d\xf0\x4d\xcd\x4d\x4a\x2d\x2a\xb6\x82\x0b\x14\x43\x15\x80\x40\x66\x0a\x9c\x99\x97\x98\x9b\x0a\xe7\xe4\x42\x34\x21\x14\x82\x40\x69\x71\x6a\x11\xaa\x08\x08\xa4\xe6\x26\x66\xe6\x60\x88\xa6\x67\x96\xe4\x24\x26\x79\xa6\xa0\x48\xd4\xa2\xf0\x8a\xf2\x73\x10\x16\x42\xa4\x6a\xb9\x6a\xb9\x00\x01\x00\x00\xff\xff\x02\xdf\xd4\xb3\xc4\x00\x00\x00")

func _lgraphqlUsergroupsAllgroupmembersGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAllgroupmembersGraphql,
		"_lgraphql/usergroups/allGroupMembers.graphql",
	)
}

func _lgraphqlUsergroupsAllgroupmembersGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAllgroupmembersGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/allGroupMembers.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAllusersGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\x49\xcd\x4d\xcc\xcc\xb1\x52\x08\x2e\x29\xca\xcc\x4b\xd7\x51\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\xd1\x51\x50\x49\xcf\x2c\xc9\x49\x4c\xf2\x84\xf0\x35\x15\xaa\xb9\x14\x14\x14\x14\x12\x73\x72\x42\x8b\x53\x8b\x8a\x35\xa0\xfa\x20\xfa\x75\x14\x40\xda\x54\x32\x53\x74\x14\x10\xba\xe0\x06\xc0\xf4\x82\x40\x66\x0a\x9c\x09\xd6\x09\xe7\xa5\x65\x16\x15\x97\xf8\x25\xe6\xa6\xc2\x45\x72\x12\xd1\x04\x92\xf3\x73\x73\x53\xf3\x4a\xc0\xfc\x5a\xae\x5a\x40\x00\x00\x00\xff\xff\x09\x21\x37\x61\xc2\x00\x00\x00")

func _lgraphqlUsergroupsAllusersGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAllusersGraphql,
		"_lgraphql/usergroups/allUsers.graphql",
	)
}

func _lgraphqlUsergroupsAllusersGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAllusersGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/allUsers.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsMeGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xb1\xae\x82\x21\x14\x83\xf7\xff\x29\x78\x9a\xbb\xdc\xc4\xc1\x18\x77\xa2\xf5\x97\x78\x0e\xe0\x01\x06\x62\x78\x77\x83\x31\x47\x02\x9d\xda\x6f\x68\xfb\x2c\x90\x6a\x5e\x9b\x31\xc6\x30\xbe\xa6\xcb\x5d\xd5\x82\xad\x23\x4d\x37\x27\x29\x1f\x2c\x43\x09\xd9\x09\x5c\x02\x33\x7c\xd6\x9c\xd2\xfd\x1f\x35\x0d\xed\xd3\x42\x97\x1f\x1b\xba\x1e\xa8\xa7\x1a\x17\x76\xb6\x54\x16\xf8\xe7\xfc\x0e\x89\xe2\x86\xd5\xcf\x13\x81\xcd\xf8\x0d\x35\x75\xbb\x84\x12\x8f\x81\x30\xdf\x5a\x7e\x48\x20\x4c\x05\x6d\x6b\xef\x00\x00\x00\xff\xff\x80\x58\xa8\xdd\x39\x01\x00\x00")

func _lgraphqlUsergroupsMeGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsMeGraphql,
		"_lgraphql/usergroups/me.graphql",
	)
}

func _lgraphqlUsergroupsMeGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsMeGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/me.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsRemovegroupsfromprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x80\x30\x3c\xf3\x0a\x4a\x4b\x14\x75\x40\x52\xe9\x45\xf9\xa5\x05\xc5\x56\x0a\xd1\xee\x20\x06\x44\x22\x56\x51\x53\xa1\x9a\x4b\x41\x41\x41\xa1\x28\x35\x37\xbf\x2c\x15\x2c\x55\xec\x56\x94\x9f\x0b\xd5\xaf\x91\x09\x52\x67\x05\x55\xa4\xa0\x00\x37\x1f\x66\x13\x54\x1c\x66\x38\xd4\x16\xb0\x68\xad\x26\x5c\x5b\x66\x0a\x94\x91\x97\x98\x9b\x0a\x91\xe4\x02\x61\x40\x00\x00\x00\xff\xff\x38\x23\xd0\x35\xbd\x00\x00\x00")

func _lgraphqlUsergroupsRemovegroupsfromprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsRemovegroupsfromprojectGraphql,
		"_lgraphql/usergroups/removeGroupsFromProject.graphql",
	)
}

func _lgraphqlUsergroupsRemovegroupsfromprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsRemovegroupsfromprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/removeGroupsFromProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsRemoveuserfromgroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\x31\x0a\xc3\x30\x0c\x45\x77\x9f\xe2\x17\x3a\x24\xd0\x13\x78\x6f\xbb\x75\x29\x3d\x80\xa1\x22\x08\x2a\x3b\x28\x76\x97\x90\xbb\x07\x25\x22\x19\x0c\xe6\xeb\xf1\x9e\xb4\x9a\x2a\x97\x8c\x2e\x00\xd7\x36\x91\xde\x25\xf1\x2f\xe2\x5d\x95\xf3\x70\xb9\xd9\x3c\x68\x69\xe3\x2b\x09\x1d\x73\x8f\x39\x00\x80\x92\x94\x3f\x7d\x26\xd2\x87\x16\x79\x1a\xd7\x71\x1e\x5b\x8d\x0e\x00\xe6\x8c\x98\x41\xbb\xf7\x6c\x60\x71\x62\xd3\x1b\x92\xb7\xc4\x99\x73\x62\xe9\x0f\x19\x7f\xfd\x63\xe8\x7e\x0c\xf6\xd6\x00\x00\x00\xff\xff\x8a\xee\x6f\x4c\xc7\x00\x00\x00")

func _lgraphqlUsergroupsRemoveuserfromgroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsRemoveuserfromgroupGraphql,
		"_lgraphql/usergroups/removeUserFromGroup.graphql",
	)
}

func _lgraphqlUsergroupsRemoveuserfromgroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsRemoveuserfromgroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/removeUserFromGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsUserbyemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x4e\xbd\xce\xc2\x30\x10\xdb\xfb\x14\xf7\x49\xdf\x50\x5e\x81\x11\x09\x16\x24\x06\x40\xec\x11\x98\x72\x22\x3f\xe5\x92\x0c\x11\xea\xbb\xa3\x14\x08\x55\xea\xe9\xce\xb6\x6c\x3f\x22\x24\x51\xfb\x0f\xa3\x58\x2f\xe9\x10\x84\x6d\xf7\xb7\xa0\x67\x43\x44\x14\x3d\x64\x95\xd6\x59\x6b\x3f\x8e\xb7\xf3\x6b\xc8\xe0\x4b\x39\x47\xad\x7c\x57\x16\x1f\x76\xca\xa0\x30\x5a\x55\xc4\xd9\x19\x03\x1b\xca\xef\xfd\x6d\x8b\xe4\x27\xe9\x55\x43\x86\x9d\x26\x64\xdc\x91\x8e\xa9\x9f\x71\x27\xa5\xe3\x8c\xdc\xb0\xed\x20\xbd\xf0\xa4\x75\x5c\x22\x50\x01\xbf\xa2\xa1\x5c\x9d\xb8\xd8\xef\x9d\x46\x3d\x6b\xb6\x43\x9c\x46\x15\x30\x34\xc3\x2b\x00\x00\xff\xff\xa1\x66\xca\x90\x63\x01\x00\x00")

func _lgraphqlUsergroupsUserbyemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsUserbyemailGraphql,
		"_lgraphql/usergroups/userByEmail.graphql",
	)
}

func _lgraphqlUsergroupsUserbyemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsUserbyemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/userByEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"_lgraphql/addEnvVariable.graphql":                                     _lgraphqlAddenvvariableGraphql,
	"_lgraphql/addNotificationEmail.graphql":                               _lgraphqlAddnotificationemailGraphql,
	"_lgraphql/addNotificationMicrosoftTeams.graphql":                      _lgraphqlAddnotificationmicrosoftteamsGraphql,
	"_lgraphql/addNotificationRocketChat.graphql":                          _lgraphqlAddnotificationrocketchatGraphql,
	"_lgraphql/addNotificationSlack.graphql":                               _lgraphqlAddnotificationslackGraphql,
	"_lgraphql/addNotificationToProject.graphql":                           _lgraphqlAddnotificationtoprojectGraphql,
	"_lgraphql/lagoonSchema.graphql":                                       _lgraphqlLagoonschemaGraphql,
	"_lgraphql/lagoonVersion.graphql":                                      _lgraphqlLagoonversionGraphql,
	"_lgraphql/sshEndpointsByProject.graphql":                              _lgraphqlSshendpointsbyprojectGraphql,
	"_lgraphql/taskByID.graphql":                                           _lgraphqlTaskbyidGraphql,
	"_lgraphql/deployments/deployEnvironmentBranch.graphql":                _lgraphqlDeploymentsDeployenvironmentbranchGraphql,
	"_lgraphql/deployments/deployEnvironmentLatest.graphql":                _lgraphqlDeploymentsDeployenvironmentlatestGraphql,
	"_lgraphql/deployments/deployEnvironmentPromote.graphql":               _lgraphqlDeploymentsDeployenvironmentpromoteGraphql,
	"_lgraphql/deployments/deployEnvironmentPullrequest.graphql":           _lgraphqlDeploymentsDeployenvironmentpullrequestGraphql,
	"_lgraphql/deployments/deploymentByName.graphql":                       _lgraphqlDeploymentsDeploymentbynameGraphql,
	"_lgraphql/deployments/deploymentByRemoteID.graphql":                   _lgraphqlDeploymentsDeploymentbyremoteidGraphql,
	"_lgraphql/deployments/getDeploymentsByBulkID.graphql":                 _lgraphqlDeploymentsGetdeploymentsbybulkidGraphql,
	"_lgraphql/deployments/getDeploymentsForEnvironment.graphql":           _lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql,
	"_lgraphql/deployments/updateDeployment.graphql":                       _lgraphqlDeploymentsUpdatedeploymentGraphql,
	"_lgraphql/deploytargets/addDeployTarget.graphql":                      _lgraphqlDeploytargetsAdddeploytargetGraphql,
	"_lgraphql/deploytargets/deleteDeployTarget.graphql":                   _lgraphqlDeploytargetsDeletedeploytargetGraphql,
	"_lgraphql/deploytargets/listDeployTargets.graphql":                    _lgraphqlDeploytargetsListdeploytargetsGraphql,
	"_lgraphql/deploytargets/updateDeployTarget.graphql":                   _lgraphqlDeploytargetsUpdatedeploytargetGraphql,
	"_lgraphql/deploytargetconfigs/addDeployTargetConfig.graphql":          _lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql,
	"_lgraphql/deploytargetconfigs/deleteDeployTargetConfig.graphql":       _lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql,
	"_lgraphql/deploytargetconfigs/deployTargetConfigsByProjectId.graphql": _lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql,
	"_lgraphql/deploytargetconfigs/updateDeployTargetConfig.graphql":       _lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql,
	"_lgraphql/projects/addProject.graphql":                                _lgraphqlProjectsAddprojectGraphql,
	"_lgraphql/projects/minimalProjectByName.graphql":                      _lgraphqlProjectsMinimalprojectbynameGraphql,
	"_lgraphql/projects/projectByName.graphql":                             _lgraphqlProjectsProjectbynameGraphql,
	"_lgraphql/projects/projectByNameMetadata.graphql":                     _lgraphqlProjectsProjectbynamemetadataGraphql,
	"_lgraphql/projects/projectsByMetadata.graphql":                        _lgraphqlProjectsProjectsbymetadataGraphql,
	"_lgraphql/projects/removeProjectMetadataByKey.graphql":                _lgraphqlProjectsRemoveprojectmetadatabykeyGraphql,
	"_lgraphql/projects/updateProject.graphql":                             _lgraphqlProjectsUpdateprojectGraphql,
	"_lgraphql/projects/updateProjectMetadata.graphql":                     _lgraphqlProjectsUpdateprojectmetadataGraphql,
	"_lgraphql/environments/addOrUpdateEnvironment.graphql":                _lgraphqlEnvironmentsAddorupdateenvironmentGraphql,
	"_lgraphql/environments/addOrUpdateEnvironmentStorage.graphql":         _lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql,
	"_lgraphql/environments/addRestore.graphql":                            _lgraphqlEnvironmentsAddrestoreGraphql,
	"_lgraphql/environments/backupsForEnvironmentByName.graphql":           _lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql,
	"_lgraphql/environments/deleteEnvironment.graphql":                     _lgraphqlEnvironmentsDeleteenvironmentGraphql,
	"_lgraphql/environments/environmentByName.graphql":                     _lgraphqlEnvironmentsEnvironmentbynameGraphql,
	"_lgraphql/environments/environmentByNamespace.graphql":                _lgraphqlEnvironmentsEnvironmentbynamespaceGraphql,
	"_lgraphql/environments/setEnvironmentServices.graphql":                _lgraphqlEnvironmentsSetenvironmentservicesGraphql,
	"_lgraphql/environments/updateEnvironment.graphql":                     _lgraphqlEnvironmentsUpdateenvironmentGraphql,
	"_lgraphql/tasks/switchActiveStandby.graphql":                          _lgraphqlTasksSwitchactivestandbyGraphql,
	"_lgraphql/tasks/updateTask.graphql":                                   _lgraphqlTasksUpdatetaskGraphql,
	"_lgraphql/usergroups/addGroup.graphql":                                _lgraphqlUsergroupsAddgroupGraphql,
	"_lgraphql/usergroups/addGroupsToProject.graphql":                      _lgraphqlUsergroupsAddgroupstoprojectGraphql,
	"_lgraphql/usergroups/addSshKey.graphql":                               _lgraphqlUsergroupsAddsshkeyGraphql,
	"_lgraphql/usergroups/addUser.graphql":                                 _lgraphqlUsergroupsAdduserGraphql,
	"_lgraphql/usergroups/addUserToGroup.graphql":                          _lgraphqlUsergroupsAddusertogroupGraphql,
	"_lgraphql/usergroups/allGroupMembers.graphql":                         _lgraphqlUsergroupsAllgroupmembersGraphql,
	"_lgraphql/usergroups/allUsers.graphql":                                _lgraphqlUsergroupsAllusersGraphql,
	"_lgraphql/usergroups/me.graphql":                                      _lgraphqlUsergroupsMeGraphql,
	"_lgraphql/usergroups/removeGroupsFromProject.graphql":                 _lgraphqlUsergroupsRemovegroupsfromprojectGraphql,
	"_lgraphql/usergroups/removeUserFromGroup.graphql":                     _lgraphqlUsergroupsRemoveuserfromgroupGraphql,
	"_lgraphql/usergroups/userByEmail.graphql":                             _lgraphqlUsergroupsUserbyemailGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"_lgraphql": &bintree{nil, map[string]*bintree{
		"addEnvVariable.graphql":                &bintree{_lgraphqlAddenvvariableGraphql, map[string]*bintree{}},
		"addNotificationEmail.graphql":          &bintree{_lgraphqlAddnotificationemailGraphql, map[string]*bintree{}},
		"addNotificationMicrosoftTeams.graphql": &bintree{_lgraphqlAddnotificationmicrosoftteamsGraphql, map[string]*bintree{}},
		"addNotificationRocketChat.graphql":     &bintree{_lgraphqlAddnotificationrocketchatGraphql, map[string]*bintree{}},
		"addNotificationSlack.graphql":          &bintree{_lgraphqlAddnotificationslackGraphql, map[string]*bintree{}},
		"addNotificationToProject.graphql":      &bintree{_lgraphqlAddnotificationtoprojectGraphql, map[string]*bintree{}},
		"deployments": &bintree{nil, map[string]*bintree{
			"deployEnvironmentBranch.graphql":      &bintree{_lgraphqlDeploymentsDeployenvironmentbranchGraphql, map[string]*bintree{}},
			"deployEnvironmentLatest.graphql":      &bintree{_lgraphqlDeploymentsDeployenvironmentlatestGraphql, map[string]*bintree{}},
			"deployEnvironmentPromote.graphql":     &bintree{_lgraphqlDeploymentsDeployenvironmentpromoteGraphql, map[string]*bintree{}},
			"deployEnvironmentPullrequest.graphql": &bintree{_lgraphqlDeploymentsDeployenvironmentpullrequestGraphql, map[string]*bintree{}},
			"deploymentByName.graphql":             &bintree{_lgraphqlDeploymentsDeploymentbynameGraphql, map[string]*bintree{}},
			"deploymentByRemoteID.graphql":         &bintree{_lgraphqlDeploymentsDeploymentbyremoteidGraphql, map[string]*bintree{}},
			"getDeploymentsByBulkID.graphql":       &bintree{_lgraphqlDeploymentsGetdeploymentsbybulkidGraphql, map[string]*bintree{}},
			"getDeploymentsForEnvironment.graphql": &bintree{_lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql, map[string]*bintree{}},
			"updateDeployment.graphql":             &bintree{_lgraphqlDeploymentsUpdatedeploymentGraphql, map[string]*bintree{}},
		}},
		"deploytargetconfigs": &bintree{nil, map[string]*bintree{
			"addDeployTargetConfig.graphql":          &bintree{_lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql, map[string]*bintree{}},
			"deleteDeployTargetConfig.graphql":       &bintree{_lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql, map[string]*bintree{}},
			"deployTargetConfigsByProjectId.graphql": &bintree{_lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql, map[string]*bintree{}},
			"updateDeployTargetConfig.graphql":       &bintree{_lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql, map[string]*bintree{}},
		}},
		"deploytargets": &bintree{nil, map[string]*bintree{
			"addDeployTarget.graphql":    &bintree{_lgraphqlDeploytargetsAdddeploytargetGraphql, map[string]*bintree{}},
			"deleteDeployTarget.graphql": &bintree{_lgraphqlDeploytargetsDeletedeploytargetGraphql, map[string]*bintree{}},
			"listDeployTargets.graphql":  &bintree{_lgraphqlDeploytargetsListdeploytargetsGraphql, map[string]*bintree{}},
			"updateDeployTarget.graphql": &bintree{_lgraphqlDeploytargetsUpdatedeploytargetGraphql, map[string]*bintree{}},
		}},
		"environments": &bintree{nil, map[string]*bintree{
			"addOrUpdateEnvironment.graphql":        &bintree{_lgraphqlEnvironmentsAddorupdateenvironmentGraphql, map[string]*bintree{}},
			"addOrUpdateEnvironmentStorage.graphql": &bintree{_lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql, map[string]*bintree{}},
			"addRestore.graphql":                    &bintree{_lgraphqlEnvironmentsAddrestoreGraphql, map[string]*bintree{}},
			"backupsForEnvironmentByName.graphql":   &bintree{_lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql, map[string]*bintree{}},
			"deleteEnvironment.graphql":             &bintree{_lgraphqlEnvironmentsDeleteenvironmentGraphql, map[string]*bintree{}},
			"environmentByName.graphql":             &bintree{_lgraphqlEnvironmentsEnvironmentbynameGraphql, map[string]*bintree{}},
			"environmentByNamespace.graphql":        &bintree{_lgraphqlEnvironmentsEnvironmentbynamespaceGraphql, map[string]*bintree{}},
			"setEnvironmentServices.graphql":        &bintree{_lgraphqlEnvironmentsSetenvironmentservicesGraphql, map[string]*bintree{}},
			"updateEnvironment.graphql":             &bintree{_lgraphqlEnvironmentsUpdateenvironmentGraphql, map[string]*bintree{}},
		}},
		"lagoonSchema.graphql":  &bintree{_lgraphqlLagoonschemaGraphql, map[string]*bintree{}},
		"lagoonVersion.graphql": &bintree{_lgraphqlLagoonversionGraphql, map[string]*bintree{}},
		"projects": &bintree{nil, map[string]*bintree{
			"addProject.graphql":                 &bintree{_lgraphqlProjectsAddprojectGraphql, map[string]*bintree{}},
			"minimalProjectByName.graphql":       &bintree{_lgraphqlProjectsMinimalprojectbynameGraphql, map[string]*bintree{}},
			"projectByName.graphql":              &bintree{_lgraphqlProjectsProjectbynameGraphql, map[string]*bintree{}},
			"projectByNameMetadata.graphql":      &bintree{_lgraphqlProjectsProjectbynamemetadataGraphql, map[string]*bintree{}},
			"projectsByMetadata.graphql":         &bintree{_lgraphqlProjectsProjectsbymetadataGraphql, map[string]*bintree{}},
			"removeProjectMetadataByKey.graphql": &bintree{_lgraphqlProjectsRemoveprojectmetadatabykeyGraphql, map[string]*bintree{}},
			"updateProject.graphql":              &bintree{_lgraphqlProjectsUpdateprojectGraphql, map[string]*bintree{}},
			"updateProjectMetadata.graphql":      &bintree{_lgraphqlProjectsUpdateprojectmetadataGraphql, map[string]*bintree{}},
		}},
		"sshEndpointsByProject.graphql": &bintree{_lgraphqlSshendpointsbyprojectGraphql, map[string]*bintree{}},
		"taskByID.graphql":              &bintree{_lgraphqlTaskbyidGraphql, map[string]*bintree{}},
		"tasks": &bintree{nil, map[string]*bintree{
			"switchActiveStandby.graphql": &bintree{_lgraphqlTasksSwitchactivestandbyGraphql, map[string]*bintree{}},
			"updateTask.graphql":          &bintree{_lgraphqlTasksUpdatetaskGraphql, map[string]*bintree{}},
		}},
		"usergroups": &bintree{nil, map[string]*bintree{
			"addGroup.graphql":                &bintree{_lgraphqlUsergroupsAddgroupGraphql, map[string]*bintree{}},
			"addGroupsToProject.graphql":      &bintree{_lgraphqlUsergroupsAddgroupstoprojectGraphql, map[string]*bintree{}},
			"addSshKey.graphql":               &bintree{_lgraphqlUsergroupsAddsshkeyGraphql, map[string]*bintree{}},
			"addUser.graphql":                 &bintree{_lgraphqlUsergroupsAdduserGraphql, map[string]*bintree{}},
			"addUserToGroup.graphql":          &bintree{_lgraphqlUsergroupsAddusertogroupGraphql, map[string]*bintree{}},
			"allGroupMembers.graphql":         &bintree{_lgraphqlUsergroupsAllgroupmembersGraphql, map[string]*bintree{}},
			"allUsers.graphql":                &bintree{_lgraphqlUsergroupsAllusersGraphql, map[string]*bintree{}},
			"me.graphql":                      &bintree{_lgraphqlUsergroupsMeGraphql, map[string]*bintree{}},
			"removeGroupsFromProject.graphql": &bintree{_lgraphqlUsergroupsRemovegroupsfromprojectGraphql, map[string]*bintree{}},
			"removeUserFromGroup.graphql":     &bintree{_lgraphqlUsergroupsRemoveuserfromgroupGraphql, map[string]*bintree{}},
			"userByEmail.graphql":             &bintree{_lgraphqlUsergroupsUserbyemailGraphql, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

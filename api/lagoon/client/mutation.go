package client

import (
	"context"
	"fmt"
	"regexp"

	"github.com/uselagoon/machinery/api/lagoon"
	"github.com/uselagoon/machinery/api/schema"
)

var duplicate = regexp.MustCompile("^graphql: Duplicate entry ")

// wrapErr wraps a response error with a lagoon.ErrExist type if the response
// is due to an object already existing
func wrapErr(err error) error {
	if err != nil && duplicate.MatchString(err.Error()) {
		return fmt.Errorf("couldn't create object: %w: %v", lagoon.ErrExist, err)
	}
	return err
}

// AddGroup adds a group.
func (c *Client) AddGroup(
	ctx context.Context, in *schema.AddGroupInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/addGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addGroup"`
	}{
		Response: out,
	})
}

// AddUser adds a user.
func (c *Client) AddUser(
	ctx context.Context, in *schema.AddUserInput, out *schema.User) error {
	req, err := c.newRequest("_lgraphql/addUser.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"addUser"`
	}{
		Response: out,
	})
}

// AddSSHKey adds an SSH key to a user.
func (c *Client) AddSSHKey(
	ctx context.Context, in *schema.AddSSHKeyInput, out *schema.SSHKey) error {
	req, err := c.newRequest("_lgraphql/addSshKey.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.SSHKey `json:"addSshKey"`
	}{
		Response: out,
	})
}

// AddUserToGroup adds a user to a group.
func (c *Client) AddUserToGroup(
	ctx context.Context, in *schema.UserGroupRoleInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/addUserToGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addUserToGroup"`
	}{
		Response: out,
	})
}

// AddNotificationSlack defines a Slack notification.
func (c *Client) AddNotificationSlack(ctx context.Context,
	in *schema.AddNotificationSlackInput, out *schema.NotificationSlack) error {
	req, err := c.newRequest("_lgraphql/addNotificationSlack.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationSlack `json:"addNotificationSlack"`
	}{
		Response: out,
	})
}

// AddNotificationRocketChat defines a RocketChat notification.
func (c *Client) AddNotificationRocketChat(ctx context.Context,
	in *schema.AddNotificationRocketChatInput,
	out *schema.NotificationRocketChat) error {
	req, err := c.newRequest("_lgraphql/addNotificationRocketChat.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationRocketChat `json:"addNotificationRocketChat"`
	}{
		Response: out,
	})
}

// AddNotificationEmail defines an Email notification.
func (c *Client) AddNotificationEmail(ctx context.Context,
	in *schema.AddNotificationEmailInput,
	out *schema.NotificationEmail) error {
	req, err := c.newRequest("_lgraphql/addNotificationEmail.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationEmail `json:"addNotificationEmail"`
	}{
		Response: out,
	})
}

// AddNotificationMicrosoftTeams defines a MicrosoftTeams notification.
func (c *Client) AddNotificationMicrosoftTeams(ctx context.Context,
	in *schema.AddNotificationMicrosoftTeamsInput,
	out *schema.NotificationMicrosoftTeams) error {
	req, err := c.newRequest("_lgraphql/addNotificationMicrosoftTeams.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationMicrosoftTeams `json:"addNotificationMicrosoftTeams"`
	}{
		Response: out,
	})
}

// AddProject adds a project.
func (c *Client) AddProject(
	ctx context.Context, in *schema.AddProjectInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/addProject.graphql", in)
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addProject"`
	}{
		Response: out,
	}))
}

// AddEnvVariable adds an EnvVariable to an Environment or Project.
func (c *Client) AddEnvVariable(ctx context.Context,
	in *schema.EnvVariableInput, out *schema.EnvKeyValue) error {
	req, err := c.newRequest("_lgraphql/addEnvVariable.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.EnvKeyValue `json:"addEnvVariable"`
	}{
		Response: out,
	})
}

// AddOrUpdateEnvironment adds or updates a Project Environment.
func (c *Client) AddOrUpdateEnvironment(ctx context.Context,
	in *schema.AddEnvironmentInput, out *schema.Environment) error {
	req, err := c.newRequest("_lgraphql/addOrUpdateEnvironment.graphql", in)
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"addOrUpdateEnvironment"`
	}{
		Response: out,
	}))
}

// AddGroupsToProject adds Groups to a Project.
func (c *Client) AddGroupsToProject(ctx context.Context,
	in *schema.ProjectGroupsInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/addGroupsToProject.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addGroupsToProject"`
	}{
		Response: out,
	})
}

// AddNotificationToProject adds a Notification to a Project.
func (c *Client) AddNotificationToProject(ctx context.Context,
	in *schema.AddNotificationToProjectInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/addNotificationToProject.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addNotificationToProject"`
	}{
		Response: out,
	})
}

// RunActiveStandbySwitch deploys a branch.
func (c *Client) RunActiveStandbySwitch(ctx context.Context,
	project string, out *schema.Task) error {
	req, err := c.newRequest("_lgraphql/switchActiveStandby.graphql", map[string]interface{}{
		"project": project,
	})
	if err != nil {
		return err
	}

	// return c.client.Run(ctx, req, &out)
	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"switchActiveStandby"`
	}{
		Response: out,
	})
}

// UpdateProjectMetadata updates a projects metadata.
func (c *Client) UpdateProjectMetadata(
	ctx context.Context, id int, key string, value string, projects *schema.ProjectMetadata) error {

	req, err := c.newRequest("_lgraphql/updateProjectMetadata.graphql",
		map[string]interface{}{
			"id":    id,
			"key":   key,
			"value": value,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.ProjectMetadata `json:"updateProjectMetadata"`
	}{
		Response: projects,
	})
}

// RemoveProjectMetadataByKey removes metadata from a project for given key.
func (c *Client) RemoveProjectMetadataByKey(
	ctx context.Context, id int, key string, projects *schema.ProjectMetadata) error {

	req, err := c.newRequest("_lgraphql/removeProjectMetadataByKey.graphql",
		map[string]interface{}{
			"id":  id,
			"key": key,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.ProjectMetadata `json:"removeProjectMetadataByKey"`
	}{
		Response: projects,
	})
}

// AddRestore adds a restore.
func (c *Client) AddRestore(
	ctx context.Context, backupID string, out *schema.Restore) error {
	req, err := c.newRequest("_lgraphql/addRestore.graphql",
		map[string]interface{}{
			"backupid": backupID,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Restore `json:"addRestore"`
	}{
		Response: out,
	})
}

// UpdateEnvironmentStorage updates an environments storage values.
func (c *Client) UpdateEnvironmentStorage(ctx context.Context,
	in *schema.UpdateEnvironmentStorageInput, out *schema.UpdateEnvironmentStorage) error {
	req, err := c.newRequest("_lgraphql/addOrUpdateEnvironmentStorage.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.UpdateEnvironmentStorage `json:"addOrUpdateEnvironmentStorage"`
	}{
		Response: out,
	})
}

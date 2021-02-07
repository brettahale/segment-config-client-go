package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
// List Roles
// GET
// List Policies
// POST
// Create Policy
// DEL
// Delete Policy
// GET
// List Invites
// POST
// Create Invite
// DEL
// Delete Invite
func (c *Client) GetIamRoles() ([]IamRole, error) {
    // 'https://platform.segmentapis.com/v1beta/workspaces/myworkspace/roles'
    iamList:= []IamRole{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/roles", c.HostURL, c.Workspace))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/roles", c.HostURL, c.Workspace), nil)
	if err != nil {
		return iamList, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return iamList, err
	}

	err = json.Unmarshal(body, &iamList)
	if err != nil {
		return iamList, err
	}

	return iamList, nil
}

func (c *Client) GetIamPolicies() ([]IamPolicy, error) {
    //'https://platform.segmentapis.com/v1beta/workspaces/myworkspace/roles/-/policies'
    iamList:= []IamPolicy{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/roles/-/policies", c.HostURL, c.Workspace))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/roles/-/policies", c.HostURL, c.Workspace), nil)
	if err != nil {
		return iamList, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return iamList, err
	}

	err = json.Unmarshal(body, &iamList)
	if err != nil {
		return iamList, err
	}

	return iamList, nil
}

func (c *Client) GetIamInvites() ([]IamInvite, error) {
    //'https://platform.segmentapis.com/v1beta/workspaces/myworkspace/roles/-/policies'
    iamList:= []IamInvite{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/invites", c.HostURL, c.Workspace))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/invites", c.HostURL, c.Workspace), nil)
	if err != nil {
		return iamList, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return iamList, err
	}

	err = json.Unmarshal(body, &iamList)
	if err != nil {
		return iamList, err
	}

	return iamList, nil
}

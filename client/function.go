package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetFunction(id string) (Function, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    function := Function{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/%s", c.HostURL, c.Workspace, id))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/%s", c.HostURL, c.Workspace, id), nil)
	if err != nil {
		return function, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return function, err
	}


	err = json.Unmarshal(body, &function)
	if err != nil {
		return function, err
	}

	return function, nil
}

func (c *Client) DeployFunction(id string) (error) {
    //curl --location --request GET 'https://platform.segmentapis.com/v1beta/workspaces/workspace_id/functions/sfn_{{source_id}}/deploy' \
    _, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/%s/deploy", c.HostURL, c.Workspace, id), nil)
    return err
}

func (c *Client) IsLatestVersionFunction(id string) (error) {
    //curl --location --request GET 'https://platform.segmentapis.com/v1beta/workspaces/workspace_id/functions/sfn_{{source_id}}/deploy' \
    _, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/%s/deploy", c.HostURL, c.Workspace, id), nil)
    return err
}

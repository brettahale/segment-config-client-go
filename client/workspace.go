package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetWorkspace() (Workspace, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    workspace := Workspace{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s", c.HostURL, c.Workspace))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s", c.HostURL, c.Workspace), nil)
	if err != nil {
		return workspace, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return workspace, err
	}


	err = json.Unmarshal(body, &workspace)
	if err != nil {
		return workspace, err
	}

	return workspace, nil
}

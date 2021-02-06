package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

POST
Create Function
GET
Get Function
GET
List Function
POST
Preview Function
PATCH
Update Function
DEL
Delete Function

func (c *Client) GetFunction(id string) (Function, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    function := Function{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/%s", c.HostURL, c.Workspace, id))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/%s", c.HostURL, c.Workspace, id), nil)
	if err != nil {
		return function, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return function, err
	}


	err = json.Unmarshal(body, &function)
	if err != nil {
		return function, err
	}

	return function, nil
}


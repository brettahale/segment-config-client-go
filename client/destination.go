package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetDestination(name string, source_name string) (Destination, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    destination := Destination{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/destinations/%s", c.HostURL, c.Workspace, source_name, name), nil)
	if err != nil {
		return destination, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return destination, err
	}


	err = json.Unmarshal(body, &destination)
	if err != nil {
		return destination, err
	}

	return destination, nil
}

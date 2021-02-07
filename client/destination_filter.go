// POST
// Create Filter
// GET
// Get Filter
// GET
// List Filters
// PATCH
// Update Filter
// DEL
// Delete Filter
// POST
// Preview Filter

package client

import (
    "bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListDestinationFilter(source_name string) ([]DestinationFilter, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    var destinations []DestinationFilter
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/destinations/filters", c.HostURL, c.Workspace, source_name), nil)
	if err != nil {
		return destinations, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return destinations, err
	}


	err = json.Unmarshal(body, &destinations)
	if err != nil {
		return destinations, err
	}

	return destinations, nil
}

func (c *Client) GetDestinationFilter(name string, source_name string, destination_name string) (DestinationFilter, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    destination := DestinationFilter{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/destinations/%s/filters/%s", c.HostURL, c.Workspace, source_name, destination_name, name), nil)
	if err != nil {
		return destination, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return destination, err
	}


	err = json.Unmarshal(body, &destination)
	if err != nil {
		return destination, err
	}

	return destination, nil
}

func (c *Client) CreateDestinationFilter(destination_name string, d DestinationFilter) (*DestinationFilter, error) {
    reqBody := DestinationFilterCreate{}
    reqBody.Filter = d
    payloadBuf := new(bytes.Buffer)
    err := json.NewEncoder(payloadBuf).Encode(reqBody)
    if err != nil {
    	return nil, err
	}
    destination := DestinationFilter{}
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1beta/%s/filters", c.HostURL, c.Workspace, destination_name), payloadBuf)
    if err != nil {
    	return nil, err
    }

    body, err := c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
    if err != nil {
    	return nil, err
    }

    err = json.Unmarshal(body, &destination)
    if err != nil {
    	return nil, err
    }

    return &destination, nil
}

func (c *Client) UpdateDestinationFilter(destination_name string, d DestinationFilter) (*DestinationFilter, error) {
    reqBody := DestinationFilterCreate{}
    reqBody.Filter = d
    payloadBuf := new(bytes.Buffer)
    err := json.NewEncoder(payloadBuf).Encode(reqBody)
    if err != nil {
    	return nil, err
	}
    destination := DestinationFilter{}
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1beta/%s/filters/%s", c.HostURL, c.Workspace, destination_name, d.Name), payloadBuf)
    if err != nil {
    	return nil, err
    }

    body, err := c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
    if err != nil {
    	return nil, err
    }

    err = json.Unmarshal(body, &destination)
    if err != nil {
    	return nil, err
    }

    return &destination, nil
}

func (c *Client) DeleteDestinationFilter(d DestinationFilter) (error) {
    _, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1beta/%s", c.HostURL, c.Workspace, d.Name), nil)
    return err
}

// func (d *DestinationFilter) Delete() error {
//     return DeleteDestinationFilter(d)
// }

package client

import (
    "bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListDestination(source_name string) ([]Destination, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    var destinations []Destination
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/destinations", c.HostURL, c.Workspace, source_name), nil)
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

func (c *Client) GetDestination(name string, source_name string) (Destination, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    destination := Destination{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/destinations/%s", c.HostURL, c.Workspace, source_name, name), nil)
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

func (c *Client) CreateDestination(d Destination, source_name string) (*Destination, error) {
    reqBody := DestinationUpsert{}
    reqBody.Destination = d
    payloadBuf := new(bytes.Buffer)
    err := json.NewEncoder(payloadBuf).Encode(reqBody)
    if err != nil {
    	return nil, err
	}
    destination := Destination{}
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/destinations", c.HostURL, c.Workspace, source_name), payloadBuf)
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

func (c *Client) UpdateDestination(d Destination, name string, source_name string, paths []string) (*Destination, error) {
    reqBody := DestinationUpsert{}
    reqBody.Destination = d
    reqBody.UpdateMask = UpdateMask{Paths:paths}
    payloadBuf := new(bytes.Buffer)
    err := json.NewEncoder(payloadBuf).Encode(reqBody)
    if err != nil {
    	return nil, err
	}
    destination := Destination{}
    req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/destinations/%s", c.HostURL, c.Workspace, source_name, name), payloadBuf)
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

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetTrackingPlan(id string) (TrackingPlan, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    trackingPlan := TrackingPlan{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans/%s", c.HostURL, c.Workspace, id), nil)
	if err != nil {
		return trackingPlan, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return trackingPlan, err
	}

	err = json.Unmarshal(body, &trackingPlan)

	if err != nil {
		return trackingPlan, err
	}

	return trackingPlan, nil
}


func (c *Client) CreateTrackingPlan(p TrackingPlan) (TrackingPlan, error) {

    reqBody := TrackingPlanCreate{}
    reqBody.Source = s
    payloadBuf := new(bytes.Buffer)
    json.NewEncoder(payloadBuf).Encode(reqBody)
    source := Source{}
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans/%s", c.HostURL, c.Workspace, ), payloadBuf)
    if err != nil {
    	return source, err
    }

    body, err := c.doRequest(req)
    if err != nil {
    	return source, err
    }

    err = json.Unmarshal(body, &source)
    if err != nil {
    	return source, err
    }

    return source, nil
}

func (c *Client) DeleteSource(name string) (error) {
    _, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s", c.HostURL, c.Workspace, name), nil)
    return err
}

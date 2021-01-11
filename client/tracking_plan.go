package client

import (
   "bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetTrackingPlan(id string) (TrackingPlan, *http.Request, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    trackingPlan := TrackingPlan{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans/%s", c.HostURL, c.Workspace, id), nil)
	if err != nil {
		return trackingPlan, req, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return trackingPlan, req, err
	}

	err = json.Unmarshal(body, &trackingPlan)

	if err != nil {
		return trackingPlan, req, err
	}

	return trackingPlan, req, nil
}


func (c *Client) CreateTrackingPlan(p TrackingPlan) (TrackingPlan, *http.Request, error) {

    reqBody := TrackingPlanCreate{}
    reqBody.TrackingPlan = p
    payloadBuf := new(bytes.Buffer)
    json.NewEncoder(payloadBuf).Encode(reqBody)
    trackingPlan := TrackingPlan{}
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans", c.HostURL, c.Workspace), payloadBuf)
    if err != nil {
    	return trackingPlan, req, err
    }

    body, err := c.doRequest(req)
    if err != nil {
    	return trackingPlan, req, err
    }

    err = json.Unmarshal(body, &trackingPlan)
    if err != nil {
    	return trackingPlan, req, err
    }

    return trackingPlan, req, nil
}

func (c *Client) DeleteTrackingPlan(name string) (string, *http.Request, error) {
    req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans/%s", c.HostURL, c.Workspace, name), nil)
    return name, req, err
}

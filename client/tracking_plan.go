package client

import (
   "bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetTrackingPlan(id string) (*TrackingPlan, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    trackingPlan := TrackingPlan{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans/%s", c.HostURL, c.Workspace, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &trackingPlan)

	if err != nil {
		return nil, err
	}

	return &trackingPlan, nil
}


func (c *Client) CreateTrackingPlan(p TrackingPlan) (*TrackingPlan, error) {

    reqBody := TrackingPlanUpsert{}
    reqBody.TrackingPlan = p
    payloadBuf := new(bytes.Buffer)
    err := json.NewEncoder(payloadBuf).Encode(reqBody)
    if err != nil {
    	return nil, err
	}
    trackingPlan := TrackingPlan{}
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans", c.HostURL, c.Workspace), payloadBuf)
    if err != nil {
    	return nil, err
    }

    body, err := c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
    if err != nil {
    	return nil, err
    }
    err = nil
    err = json.Unmarshal(body, &trackingPlan)
    if err != nil {
    	return nil, err
    }

    return &trackingPlan, nil
}

func (c *Client) UpdateTrackingPlan(p TrackingPlan, name string, paths []string) (*TrackingPlan, error) {

	reqBody := TrackingPlanUpsert{}
	reqBody.TrackingPlan = p
	reqBody.UpdateMask = UpdateMask{Paths: paths}
	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans/%s", c.HostURL, c.Workspace, name), payloadBuf)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
	if err != nil {
		return nil, err
	}

	trackingPlan := TrackingPlan{}
	err = json.Unmarshal(body, &trackingPlan)
	if err != nil {
		return nil, err
	}

	return &trackingPlan, nil
}

func (c *Client) DeleteTrackingPlan(name string) (error) {
    req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans/%s", c.HostURL, c.Workspace, name), nil)
    if err != nil {
      return err
    }
    _, err = c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
    return err
}

func (c *Client) CreateTrackingPlanSourceConnection(name string, connection TrackingPlanSourceConnection) (*TrackingPlanSourceConnection, error) {

    payloadBuf := new(bytes.Buffer)
    err := json.NewEncoder(payloadBuf).Encode(connection)
    if err != nil {
    	return nil, err
	}
    trackingPlanSourceConnection := TrackingPlanSourceConnection{}
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans/%s/source-connections", c.HostURL, c.Workspace, name), payloadBuf)
    if err != nil {
    	return nil, err
    }

    body, err := c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
    if err != nil {
    	return nil, err
    }
    err = nil
    err = json.Unmarshal(body, &trackingPlanSourceConnection)
    if err != nil {
    	return nil, err
    }

    return &trackingPlanSourceConnection, nil
}

func (c *Client) DeleteTrackingPlanSourceConnection(trackingPlanName, name string) (error) {
    req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1beta/workspaces/%s/tracking-plans/%s/source-connections/%s", c.HostURL, c.Workspace, trackingPlanName, name), nil)
    if err != nil {
      return err
    }
    _, err = c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
    return err
}

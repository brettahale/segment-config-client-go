package client

import (
    "bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetSource(name string) (Source, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js
    source := Source{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s", c.HostURL, c.Workspace, name))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s", c.HostURL, c.Workspace, name), nil)
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

func (c *Client) GetSources() ([]Source, error) {
    //https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources?page_size=&page_token=

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources", c.HostURL, c.Workspace), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	sources := []Source{}
	err = json.Unmarshal(body, &sources)
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func (c *Client) CreateSource(s Source) (Source, error) {

    reqBody := SourceCreate{}
    reqBody.Source = s
    payloadBuf := new(bytes.Buffer)
    json.NewEncoder(payloadBuf).Encode(reqBody)
    source := Source{}
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1beta/workspaces/%s/sources", c.HostURL, c.Workspace), payloadBuf)
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

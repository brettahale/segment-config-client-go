package client

import (
    "bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetSource(name string) (*Source, error) {
    source := Source{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s", c.HostURL, c.Workspace, name))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s", c.HostURL, c.Workspace, name), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return nil, err
	}


	err = json.Unmarshal(body, &source)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (c *Client) ListSource() ([]*Source, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources", c.HostURL, c.Workspace), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return nil, err
	}

	sources := []*Source{}
	err = json.Unmarshal(body, &sources)
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func (c *Client) CreateSource(s Source) (*Source, error) {

    reqBody := SourceCreate{}
    reqBody.Source = s
    payloadBuf := new(bytes.Buffer)
    err := json.NewEncoder(payloadBuf).Encode(reqBody)
    if err != nil {
    	return nil, err
	}
    source := Source{}
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1beta/workspaces/%s/sources", c.HostURL, c.Workspace), payloadBuf)
    if err != nil {
    	return nil, err
    }

    body, err := c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
    if err != nil {
    	return nil, err
    }

    err = json.Unmarshal(body, &source)
    if err != nil {
    	return nil, err
    }

    return &source, nil
}

func (c *Client) DeleteSource(name string) (error) {
    _, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s", c.HostURL, c.Workspace, name), nil)
    return err
}

func (c *Client) GetSourceSchemaConfiguration(name string) (*SourceSchemaConfiguration, error){
    source := SourceSchemaConfiguration{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/schema-config", c.HostURL, c.Workspace, name))
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/schema-config", c.HostURL, c.Workspace, name), nil)
    if err != nil {
    	return nil, err
    }

    body, err := c.doRequest(req, []int{http.StatusOK})
    if err != nil {
    	return nil, err
    }


    err = json.Unmarshal(body, &source)
    if err != nil {
    	return nil, err
    }

    return &source, nil
}

func (c *Client) UpdateSourceSchemaConfiguration(d SourceSchemaConfiguration, name string, paths []string) (*SourceSchemaConfiguration, error){
    //PATCH https://platform.segmentapis.com/v1beta/workspaces/myworkspace/sources/js/schema-config

    if len(paths) < 1 {
        return nil, errors.New("UpdateMask must be set to update entities, otherwise, you could null some fields")
    }

    reqBody := SourceSchemaConfigurationUpsert{}
    reqBody.SourceSchemaConfiguration = d
    reqBody.UpdateMask = UpdateMask{Paths:paths}
    payloadBuf := new(bytes.Buffer)
    err := json.NewEncoder(payloadBuf).Encode(reqBody)
    if err != nil {
    	return nil, err
	}
    sourceSchemaConfiguration := SourceSchemaConfiguration{}
    req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/schema-config", c.HostURL, c.Workspace, name), payloadBuf)
    if err != nil {
    	return nil, err
    }

    body, err := c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
    if err != nil {
    	return nil, err
    }

    err = json.Unmarshal(body, &sourceSchemaConfiguration)
    if err != nil {
    	return nil, err
    }

    return &sourceSchemaConfiguration, nil
}

func (c *Client) UpdateSourceConnectedWarehouses(u SourceConnectedWarehousesUpdate, name string) (error){

   payloadBuf := new(bytes.Buffer)
   err := json.NewEncoder(payloadBuf).Encode(u)
   if err != nil {
   	return err
   }
   req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v1beta/workspaces/%s/sources/%s/connected-warehouses", c.HostURL, c.Workspace, name), payloadBuf)
   if err != nil {
   	return err
   }

   _, err = c.doRequest(req, []int{http.StatusOK, http.StatusCreated})
   if err != nil {
   	return err
   }

   return nil
}

func (c *Client) IsLatestVersionSourceFunction (source_id string) (bool, error) {
    // GET https://platform.segmentapis.com/v1beta/workspaces/workspace_id/functions/sfn_{{source_id}}
    source := SourceLatestFunction{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/sfn_%s", c.HostURL, c.Workspace, source_id))
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/sfn_%s", c.HostURL, c.Workspace, source_id), nil)
    if err != nil {
    	return false, err
    }

    body, err := c.doRequest(req, []int{http.StatusOK})
    if err != nil {
    	return false, err
    }

    err = json.Unmarshal(body, &source)
    if err != nil {
    	return false, err
    }

    return source.IsLatest, nil
}

func (c *Client) DeploySourceFunction (source_id string) error {
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/sfn_%s/deploy", c.HostURL, c.Workspace, source_id))
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/workspaces/%s/functions/sfn_%s/deploy", c.HostURL, c.Workspace, source_id), nil)
    if err != nil {
    	return err
    }

    _, err = c.doRequest(req, []int{http.StatusOK})
    return err
}

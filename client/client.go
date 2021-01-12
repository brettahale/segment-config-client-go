package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "https://platform.segmentapis.com"

// Client -
type Client struct {
	HostURL     string
	HTTPClient  *http.Client
	Workspace   string
	Token       string
}

// NewClient -
func NewClient(host, workspace, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if workspace != nil {
    	c.Workspace = *workspace
    }

	if token != nil {
    	c.Token = *token
    }

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, accepts []int) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
    req.Header.Set("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	for _, statusCode := range accepts {
	    if statusCode == res.StatusCode {
	        return body, nil
	    }
	}

	return nil, fmt.Errorf("StatusCode: %d, StatusText: %s,  body: %s", res.StatusCode, http.StatusText(res.StatusCode), body)

}

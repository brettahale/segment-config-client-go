package segment

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "https://platform.segmentapis.com"

// Client -
type SegmentClient struct {
	HostURL     string
	HTTPClient  *http.Client
	Workspace   string
	Token       string
}

// NewClient -
func NewClient(host, workspace, token *string) (*SegmentClient, error) {
	c := SegmentClient{
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

func (c *SegmentClient) doRequest(req *http.Request) ([]byte, error) {
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

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
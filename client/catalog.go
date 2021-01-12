package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetCatalogSource(name string) (CatalogSource, error) {
    source := CatalogSource{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/catalog/sources/%s", c.HostURL, name))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/catalog/sources/%s", c.HostURL, name), nil)
	if err != nil {
		return source, err
	}

	body, err := c.doRequest(req, []int{http.StatusOK})
	if err != nil {
		return source, err
	}


	err = json.Unmarshal(body, &source)
	if err != nil {
		return source, err
	}

	return source, nil
}

func (c *Client) GetCatalogDestination(name string) (CatalogDestination, error) {
    destination := CatalogDestination{}
    log.Printf("[DEBUG] HTTP request to API %s ", fmt.Sprintf("%s/v1beta/catalog/destinations/%s", c.HostURL, name))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1beta/catalog/destinations/%s", c.HostURL, name), nil)
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

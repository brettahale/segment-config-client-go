package segment

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *SegmentClient) GetTrackingPlan(id string) (TrackingPlan, error) {
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


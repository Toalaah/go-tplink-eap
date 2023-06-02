package tplink

import (
	"net/http"
	"net/url"
)

type LedInfo struct {
	Enable string `json:"enable"`
}

type LedStatus string

const (
	LedStatusOff LedStatus = "off"
	LedStatusOn            = "on"
)

// GetLedStatus returns the EAP's LED status from the `/data/ledctrl.json` endpoint
func (c *TPLinkClient) GetLedStatus() (LedInfo, error) {

	var res LedInfo

	resp, err := c.makeRequest(http.MethodGet, "/data/ledctrl.json", &url.Values{"operation": {"read"}}, nil)
	if err != nil {
		return res, err
	}

	parseFromBodyNested(resp, &res)

	return res, err
}

// GetLedStatus sets the EAP's LED status either on or off
func (c *TPLinkClient) SetLedStatus(status LedStatus) (LedInfo, error) {

	var res LedInfo

	resp, err := c.makeRequest(http.MethodPost, "/data/ledctrl.json", &url.Values{"operation": {"write"}, "enable": {string(status)}}, nil)
	if err != nil {
		return res, err
	}

	parseFromBodyNested(resp, &res)

	return res, err
}

package tplink

import (
	"net/http"
	"net/url"
	"strconv"
)

type RadioInfo struct {
	Status       string `json:"status,omitempty"`
	Wirelessmode int    `json:"wirelessmode,omitempty"`
	Chwidth      int    `json:"chwidth,omitempty"`
	Channel      int    `json:"channel,omitempty"`
	Txpower      int    `json:"txpower,omitempty"`
	RegionCode   int    `json:"regionCode,omitempty"`
	RadioID      int    `json:"radioID,omitempty"`
	MinPower     int    `json:"minPower,omitempty"`
}

// GetRadioInfo returns the information about the passed radio from the EAP. It calls the `/data/wireless.basic.json` endpoint under the hood.
func (c *TPLinkClient) GetRadioInfo(radioId int) (RadioInfo, error) {

	var res RadioInfo

	params := &url.Values{
		"radioID": {strconv.Itoa(radioId)},
	}

	resp, err := c.makeRequest(http.MethodGet, "/data/wireless.basic.json", nil, params)
	if err != nil {
		return res, err
	}

	if err = parseFromBodyNested(resp, &res); err != nil {
		return res, err
	}

	return res, err
}
package tplink

import (
	"net/http"
	"net/url"
)

type DeviceInfo struct {
	DeviceName      string `json:"deviceName,omitempty"`
	DeviceModel     string `json:"deviceModel,omitempty"`
	FirmwareVersion string `json:"firmwareVersion,omitempty"`
	HardwareVersion string `json:"hardwareVersion,omitempty"`
	Mac             string `json:"mac,omitempty"`
	IP              string `json:"ip,omitempty"`
	SubnetMask      string `json:"subnetMask,omitempty"`
	LanPortList     []struct {
		Status string `json:"status,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"lan_port_list,omitempty"`
	Time   string `json:"time,omitempty"`
	Uptime string `json:"uptime,omitempty"`
	CPU    int    `json:"cpu,omitempty"`
	Memory int    `json:"memory,omitempty"`
}

// GetDeviceInfo returns the client's device info from the `/data/status.device.json` endpoint.
func (c *TPLinkClient) GetDeviceInfo() (DeviceInfo, error) {

	var res DeviceInfo

	params := &url.Values{
		"operation": {"read"},
	}

	resp, err := c.makeRequest(http.MethodGet, "/data/status.device.json", nil, params)
	if err != nil {
		return res, err
	}

	if err = parseFromBodyNested(resp, &res); err != nil {
		return res, err
	}

	return res, err
}

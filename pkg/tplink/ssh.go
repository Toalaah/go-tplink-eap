package tplink

import (
	"net/http"
	"net/url"
	"strconv"
)

type SSHInfo struct {
	// Whether the AP's SSH server is enabled
	SSHServerEnable bool `json:"sshServerEnable"`
	// Whether layer-3 accessibility is enabled
	RemoteEnable bool `json:"remoteEnable"`
	// What port the AP's SSH server is running on
	ServerPort int `json:"serverPort"`
}

// GetLedStatus returns information of the EAP's SSH-Server.
func (c *TPLinkClient) GetSSHStatus() (SSHInfo, error) {

	var res SSHInfo
	form := &url.Values{
		"operation": {"read"},
	}

	resp, err := c.makeRequest(http.MethodPost, "/data/sshServer.json", form, nil)
	if err != nil {
		return res, err
	}

	err = parseFromBody(resp, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

// GetLedStatus configures the EAP's SSH-Server.
func (c *TPLinkClient) SetSSHStatus(info SSHInfo) (responseStatus, error) {

	var res responseStatus

	form := &url.Values{
		"operation":       {"write"},
		"remoteEnable":    {strconv.FormatBool(info.RemoteEnable)},
		"serverPort":      {strconv.Itoa(info.ServerPort)},
		"sshServerEnable": {strconv.FormatBool(info.SSHServerEnable)},
	}
	resp, err := c.makeRequest(http.MethodPost, "/data/sshServer.json", form, nil)
	if err != nil {
		return res, err
	}

	err = parseFromBody(resp, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

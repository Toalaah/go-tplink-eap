package tplink

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetSSIDInfo returns general diagnostics about the EAP's configured SSIDs.
func (c *TPLinkClient) GetSSIDDiagnostics() ([]SSIDDiagnostic, error) {
	var res []SSIDDiagnostic

	params := &url.Values{
		"operation": {"load"},
	}

	resp, err := c.makeRequest(http.MethodGet, "/data/status.wireless.ssid.json", nil, params)
	if err != nil {
		return res, err
	}

	if err = parseFromBody(resp, &res); err != nil {
		return res, err
	}

	return res, err
}

// GetSSIDInfo returns more detailed diagnostics of the EAP's configured SSIDs for a given radio.
func (c *TPLinkClient) GetSSIDDs(radioID int) ([]SSID, error) {

	var res []SSID

	form := &url.Values{
		"operation": {"load"},
		"radioID":   {strconv.Itoa(radioID)},
	}

	resp, err := c.makeRequest(http.MethodPost, "/data/wireless.ssids.json", form, nil)
	if err != nil {
		return res, err
	}

	if err = parseFromBody(resp, &res); err != nil {
		return res, err
	}

	return res, err
}

// GetSSIDInfo createss a new SSID on the given band. On successful execution, it returns a list of all the currently configured SSIDs (including the newly created one).
func (c *TPLinkClient) CreateSSIDD(s SSID, radioID int) ([]SSID, error) {

	var res []SSID

	b, err := json.Marshal(s)
	if err != nil {
		return res, err
	}

	form := &url.Values{
		"operation": {"insert"},
		"key":       {"add"},
		"index":     {"0"},
		"old":       {"add"},
		"new":       {string(b)},
		"radioID":   {strconv.Itoa(radioID)},
	}

	resp, err := c.makeRequest(http.MethodPost, "/data/wireless.ssids.json", form, nil)
	if err != nil {
		return res, err
	}

	if err = parseFromBody(resp, &res); err != nil {
		return res, err
	}

	return res, err
}

// GetSSIDInfo deletes an existing SSID and returns the remaining SSIDs.
func (c *TPLinkClient) DeleteSSIDD(s ssidDeleter, radioID int) ([]SSID, error) {

	var res []SSID

	form := &url.Values{
		"operation": {"remove"},
		"radioID":   {strconv.Itoa(radioID)},
		"key":       {strconv.Itoa(s.GetKey())},
	}

	resp, err := c.makeRequest(http.MethodPost, "/data/wireless.ssids.json", form, nil)
	if err != nil {
		return res, err
	}

	if err = parseFromBody(resp, &res); err != nil {
		return res, err
	}

	return res, err
}

type SSID struct {
	Ssidname           string `json:"ssidname"`
	Vlanid             int    `json:"vlanid"`
	Ssidbcast          int    `json:"ssidbcast"`
	Guest              int    `json:"guest"`
	Portal             int    `json:"portal"`
	Key                int    `json:"key"`
	Limit              bool   `json:"limit"`
	LimitDownload      int    `json:"limit_download"`
	LimitDownloadUnit  int    `json:"limit_download_unit"`
	LimitUpload        int    `json:"limit_upload"`
	LimitUploadUnit    int    `json:"limit_upload_unit"`
	SecurityMode       int    `json:"securityMode"`
	PskVersion         int    `json:"psk_version"`
	PskCipher          int    `json:"psk_cipher"`
	PskKey             string `json:"psk_key"`
	PskKeyUpdate       int    `json:"psk_key_update"`
	WpaVersion         int    `json:"wpa_version"`
	WpaCipher          int    `json:"wpa_cipher"`
	Server             string `json:"server"`
	Port               int    `json:"port"`
	WpaKey             string `json:"wpa_key"`
	WpaKeyUpdate       int    `json:"wpa_key_update"`
	AcctEnable         int    `json:"acct_enable"`
	AcctServer         string `json:"acct_server"`
	AcctPort           int    `json:"acct_port"`
	AcctKey            string `json:"acct_key"`
	AcctUpdateEnable   int    `json:"acct_update_enable"`
	AcctUpdateInterval int    `json:"acct_update_interval"`
	WepMode            int    `json:"wep_mode"`
	WepSelect          int    `json:"wep_select"`
	WepFormat1         int    `json:"wep_format1"`
	WepKey1            string `json:"wep_key1"`
	WepType1           int    `json:"wep_type1"`
	WepFormat2         int    `json:"wep_format2"`
	WepKey2            string `json:"wep_key2"`
	WepType2           int    `json:"wep_type2"`
	WepFormat3         int    `json:"wep_format3"`
	WepKey3            string `json:"wep_key3"`
	WepType3           int    `json:"wep_type3"`
	WepFormat4         int    `json:"wep_format4"`
	WepKey4            string `json:"wep_key4"`
	WepType4           int    `json:"wep_type4"`
}

type SSIDDiagnostic struct {
	Ssid     string `json:"SSID"`
	Key      int    `json:"key"`
	Vlan     int    `json:"vlan"`
	Clients  int    `json:"clients"`
	Guest    bool   `json:"guest"`
	Radio    int    `json:"Radio"`
	Portal   bool   `json:"portal"`
	Security int    `json:"security"`
	DownTh   int    `json:"downTh"`
	UpTh     int    `json:"upTh"`
}

type ssidDeleter interface {
	GetKey() int
}

func (s SSID) GetKey() int           { return s.Key }
func (s SSIDDiagnostic) GetKey() int { return s.Key }

func (s SSID) MarshalJSON() ([]byte, error) {
	// https://stackoverflow.com/questions/43176625/
	type testerImpl SSID
	marshaledTester, err := json.Marshal(testerImpl(s))
	if err != nil {
		return nil, err
	}

	var tmp map[string]interface{}
	err = json.Unmarshal(marshaledTester, &tmp)
	if err != nil {
		return nil, err
	}

	// convert each field in the converted struct-map to a string, then re-marshal
	n := make(map[string]string)
	for k, v := range tmp {
		switch v := v.(type) {
		case bool:
			if v {
				n[k] = "true"
			} else {
				n[k] = ""
			}
		default:
			n[k] = fmt.Sprintf("%v", v)
		}
	}

	return json.Marshal(n)
}

package tplink

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *TPLinkClient) makeRequest(method, path string, formData *url.Values, params *url.Values) (*http.Response, error) {
	if !c.IsAuthenticated() && path != "/" {
		if err := c.Authenticate(); err != nil {
			return nil, err
		}
	}

	endpoint := c.BaseAddr.JoinPath(path)

	var r *http.Request
	var err error
	if formData != nil {
		r, err = http.NewRequest(method, endpoint.String(), strings.NewReader(formData.Encode()))
	} else {
		r, err = http.NewRequest(method, endpoint.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Referer", c.BaseAddr.String())

	if params != nil {
		r.URL.RawQuery = params.Encode()
	}

	return c.httpClient.Do(r)
}

func parseFromBody(resp *http.Response, out interface{}) error {
	// parse the response status from the to check for error codes returned despite a 200 OK
	var status responseStatus
	var tmp interface{}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &status)
	if err != nil {
		return err
	}

	if !status.Ok() {
		return fmt.Errorf("received error: %s", status.String())
	}

	err = json.Unmarshal(body, &tmp)
	if err != nil {
		return err
	}

	b, err := json.Marshal(tmp.(map[string]interface{})["data"])
	if err != nil {
		return err
	}

	return json.Unmarshal(b, out)
}

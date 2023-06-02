package tplink

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func (c *TPLinkClient) IsAuthenticated() bool {
	if c.httpClient.Jar == nil {
		return false
	}
	cookies := c.httpClient.Jar.Cookies(c.BaseAddr)
	return len(cookies) > 0 && cookies[0].Expires.Before(time.Now())
}

func (c *TPLinkClient) Authenticate() error {
	// get the cookie to authenticate
	resp, err := c.makeRequest(http.MethodGet, "/", nil, nil)
	if err != nil {
		return err
	}
	c.httpClient.Jar.SetCookies(c.BaseAddr, resp.Cookies())

	// now authenticate the cookie
	form := &url.Values{
		"username": {c.Username},
		"password": {c.Password},
	}
	resp, err = c.makeRequest(http.MethodPost, "/", form, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not authenticate, received: %s", resp.Status)
	}

	return nil
}

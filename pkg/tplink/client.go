package tplink

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type TPLinkClient struct {
	BaseAddr *url.URL
	Username string
	Password string

	httpClient http.Client
}

func (c *TPLinkClient) Cookie() string {
	if c.httpClient.Jar == nil {
		return ""
	}
	cookies := c.httpClient.Jar.Cookies(c.BaseAddr)
	if len(cookies) == 0 {
		return ""
	}
	return cookies[0].Value
}

func NewClient(baseAddr, username, password string) *TPLinkClient {
	// md5-hash password, is required for basic-auth on tplink side
	pw := md5.Sum([]byte(password))
	url, err := url.Parse(baseAddr)
	if err != nil {
		panic(err)
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	return &TPLinkClient{
		BaseAddr: url,
		Username: username,
		Password: strings.ToUpper(hex.EncodeToString(pw[:])),
		httpClient: http.Client{
			Jar: jar,
		},
	}
}

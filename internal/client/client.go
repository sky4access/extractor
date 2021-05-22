package client

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL *url.URL
	*http.Client
	auth string
}

func NewClient(address string, username, password string) (c *Client, err error) {
	c = new(Client)
	c.BaseURL, _ = url.Parse(address)
	c.auth = basicAuth(username, password)

	timeout := time.Duration(5 * time.Second)
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
	}

	c.Client = &http.Client{
		Transport: transport,
		Timeout:   time.Second * 2,
	}

	return c, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

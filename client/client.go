package client

import (
	"fmt"
	"net/http"
	"net/url"
)

/*
type Getter interface {
	Get(string) (*http.Response, error)
}
*/

type Client struct {
	c http.Client
	// c Getter
	baseURL string
}

// TODO: Context
func (c *Client) Health() error {
	u, err := url.JoinPath(c.baseURL, "/health")
	if err != nil {
		return err
	}

	resp, err := c.c.Get(u)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%q - %s", u, resp.Status)
	}

	return nil
}

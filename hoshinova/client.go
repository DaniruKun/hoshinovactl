package hoshinova

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Host  string // Where hoshinova is hosted at, together with port (e.g. localhost:1104)
	HTTPc http.Client
}

// Returns a client ready to call Hoshinova API
func NewClient(host, port string) (*Client, error) {
	client := &Client{
		HTTPc: http.Client{Timeout: time.Duration(10) * time.Second},
		Host:  host + ":" + port,
	}

	return client, nil
}

// Returns the Hoshinova server version (e.g. `hoshinova v0.2.2 (108a7aa)`)
func (client *Client) GetVersion() (string, error) {
	url := url.URL{Scheme: "http", Host: client.Host, Path: "/api/version"}

	res, err := client.HTTPc.Get(url.String())
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return string(body), nil
}

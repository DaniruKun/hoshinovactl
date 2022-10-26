package hoshinova

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Host  string // Where hoshinova is hosted at, together with port (e.g. localhost:1104)
	HTTPc http.Client
}

type Task struct {
	Title           string `json:"title"`
	VideoId         string `json:"video_id"`
	VideoPicture    string `json:"video_picture"`
	ChannelName     string `json:"channel_name"`
	ChannelId       string `json:"channel_id"`
	ChannelPicture  string `json:"channel_picture"`
	OutputDirectory string `json:"output_directory"`
}

type Status struct {
	Version        string      `json:"version"`
	State          interface{} `json:"state"` // can be either string or map[string]string
	LastOutput     string      `json:"last_output"`
	LastUpdate     string      `json:"last_update"`
	VideoFragments int         `json:"video_fragments"`
	AudioFragments int         `json:"audio_fragments"`
	TotalSize      string      `json:"total_size"`
	VideoQuality   string      `json:"video_quality"`
	OutputFile     string      `json:"output_file"`
}

type TaskItem struct {
	Task   `json:"task"`
	Status `json:"status"`
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

// Returns the currently loaded configuration on the server
func (client *Client) GetConfig() (string, error) {
	url := url.URL{Scheme: "http", Host: client.Host, Path: "/api/config"}

	res, err := client.HTTPc.Get(url.String())
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return string(body), nil
}

// Returns slice of currently scheduled tasks
func (client *Client) GetTasks() ([]TaskItem, error) {
	var taskItems []TaskItem
	url := url.URL{Scheme: "http", Host: client.Host, Path: "/api/tasks"}

	res, err := client.HTTPc.Get(url.String())
	if err != nil {
		return taskItems, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &taskItems)

	return taskItems, nil
}

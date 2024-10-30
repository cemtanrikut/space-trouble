package spacex

import (
	"encoding/json"
	"net/http"
	"time"
)

type Launchpad struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location struct {
		Name      string  `json:"name"`
		Region    string  `json:"region"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
}

type Launch struct {
	ID          string `json:"id"`
	MissionName string `json:"mission_name"`
	LaunchDate  string `json:"launch_date_utc"`
	LaunchpadID string `json:"launchpad"`
}

// Client for send request SpaceX API
type Client struct {
	httpClient *http.Client
	BaseURL    string
}

// NewClient creates Client for SpaceX API
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    "https://api.spacexdata.com/v4",
	}
}

// GetLaunchpads fetches all available launchpads from SpaceX API
func (c *Client) GetLaunchpads() ([]Launchpad, error) {
	resp, err := c.httpClient.Get(c.BaseURL + "/launchpads")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var launchpads []Launchpad
	if err := json.NewDecoder(resp.Body).Decode(&launchpads); err != nil {
		return nil, err
	}

	return launchpads, nil
}

// GetUpcomingLaunches fetches upcoming launches from SpaceX API
func (c *Client) GetUpcomingLaunches() ([]Launch, error) {
	resp, err := c.httpClient.Get(c.BaseURL + "/launches/upcoming")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var launches []Launch
	if err := json.NewDecoder(resp.Body).Decode(&launches); err != nil {
		return nil, err
	}

	return launches, nil
}

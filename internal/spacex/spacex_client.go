package spacex

import (
	"encoding/json"
	"net/http"
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

// GetLaunchpads fetches all available launchpads from SpaceX API
func GetLaunchpads() ([]Launchpad, error) {
	url := "https://api.spacexdata.com/v4/launchpads"
	resp, err := http.Get(url)
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
func GetUpcomingLaunches() ([]Launch, error) {
	url := "https://api.spacexdata.com/v4/launches/upcoming"
	resp, err := http.Get(url)
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

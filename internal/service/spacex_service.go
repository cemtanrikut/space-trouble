package service

import (
	"context"
	"space-trouble/internal/spacex"
)

type SpaceXService struct {
	SpaceXClient *spacex.Client
}

func NewSpaceXService(client *spacex.Client) *SpaceXService {
	return &SpaceXService{
		SpaceXClient: client,
	}
}

func (s *SpaceXService) GetUpcomingLaunches(ctx context.Context) ([]spacex.Launch, error) {
	return s.SpaceXClient.GetUpcomingLaunches()
}

func (s *SpaceXService) GetLaunchpads(ctx context.Context) ([]spacex.Launchpad, error) {
	return s.SpaceXClient.GetLaunchpads()
}

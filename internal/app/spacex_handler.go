package app

import (
	"encoding/json"
	"net/http"
	"space-trouble/internal/service"
)

type SpaceXHandler struct {
	SpaceXService *service.SpaceXService
}

func NewSpaceXHandler(spaceXService *service.SpaceXService) *SpaceXHandler {
	return &SpaceXHandler{
		SpaceXService: spaceXService,
	}
}

func (h *SpaceXHandler) GetUpcomingLaunches(w http.ResponseWriter, r *http.Request) {
	launches, err := h.SpaceXService.GetUpcomingLaunches(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(launches)
}

func (h *SpaceXHandler) GetLaunchpads(w http.ResponseWriter, r *http.Request) {
	launchpads, err := h.SpaceXService.GetLaunchpads(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(launchpads)
}

package api

import (
	"encoding/json"
	"net/http"

	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/rata"

	"github.com/concourse/baggageclaim"
	"github.com/concourse/baggageclaim/volume"
)

func NewHandler(logger lager.Logger, volumeRepo volume.Repository) (http.Handler, error) {
	volumeServer := NewVolumeServer(
		logger.Session("volume-server"),
		volumeRepo,
	)

	handlers := rata.Handlers{
		baggageclaim.CreateVolume: http.HandlerFunc(volumeServer.CreateVolume),
		baggageclaim.GetVolumes:   http.HandlerFunc(volumeServer.GetVolumes),
		baggageclaim.SetProperty:  http.HandlerFunc(volumeServer.SetProperty),
	}

	return rata.NewRouter(baggageclaim.Routes, handlers)
}

type errorResponse struct {
	Message string `json:"error"`
}

func respondWithError(w http.ResponseWriter, err error, statusCode ...int) {
	var code int

	if len(statusCode) > 0 {
		code = statusCode[0]
	} else {
		code = http.StatusInternalServerError
	}

	w.WriteHeader(code)
	errResponse := errorResponse{Message: err.Error()}
	json.NewEncoder(w).Encode(errResponse)
}

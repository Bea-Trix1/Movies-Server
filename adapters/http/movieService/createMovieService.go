package movieservice

import (
	"encoding/json"
	"net/http"

	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (service service) CreateMovie(response http.ResponseWriter, request *http.Request) {
	movieRequest, err := dto.JSONMovieRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	movies, err := service.usecase.CreateMovie(movieRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(movies)
}

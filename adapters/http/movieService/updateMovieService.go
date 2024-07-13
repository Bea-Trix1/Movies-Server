package movieservice

import (
	"encoding/json"
	"net/http"

	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (service service) UpdateMovie(response http.ResponseWriter, request *http.Request) {
	movieRequest, err := dto.JSONMovieRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	movie, err := service.usecase.UpdateMovie(movieRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(movie)
}

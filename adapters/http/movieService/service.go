package movieservice

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (service service) DeleteMovie(response http.ResponseWriter, request *http.Request) {

	movieID := request.URL.Query().Get("id")
	if movieID == "" {
		http.Error(response, "Esperado ID do filme", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(movieID, 10, 32)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	err = service.usecase.DeleteMovie(uint32(id))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Filme deletado com sucesso!"))
}

func (service service) GetAllMovies(response http.ResponseWriter, request *http.Request) {
	_, err := dto.JSONMovieRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	movie, err := service.usecase.GetAllMovies()

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(movie)
}

func (service service) GetMovieById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	movieID := request.URL.Query().Get("id")
	if movieID == "" {
		http.Error(response, "Esperado ID do filme", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(movieID, 10, 32)
	if err != nil {
		http.Error(response, "ID do filme inválido", http.StatusBadRequest)
		return
	}

	movie, err := service.usecase.GetMovieById(uint32(id))
	if err != nil {
		http.Error(response, "Erro ao obter o filme: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(response).Encode(movie); err != nil {
		http.Error(response, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
	}
}

func (service service) UpdateMovie(response http.ResponseWriter, request *http.Request) {
	movieRequest, err := dto.JSONMovieRequest(request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedMovie, err := service.usecase.UpdateMovie(movieRequest)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(updatedMovie)
}

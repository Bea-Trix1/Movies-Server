package domain

import (
	"net/http"

	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Gender   string `json:"gender"`
	Year     string `json:"year"`
	Director string `json:"director"`
}

// O service irá atender as requisições externas que batem na API.
type MovieService interface {
	CreateMovie(resp http.ResponseWriter, req *http.Request)
	UpdateMovie(resp http.ResponseWriter, req *http.Request)
	DeleteMovie(resp http.ResponseWriter, req *http.Request)
	GetAllMovies(resp http.ResponseWriter, req *http.Request)
	GetMovieById(resp http.ResponseWriter, req *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

// usecase é a regra de negócio.
type MovieUseCase interface {
	CreateMovie(req *dto.MovieRequest) (*Movie, error)
	DeleteMovie(req *dto.MovieRequest) (*Movie, error)
	GetAllMovies(req *dto.MovieRequest) (*Movie, error)
	GetMovieById(req *dto.MovieRequest) (*Movie, error)
	UpdateMovie(req *dto.MovieRequest) (*Movie, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination[[]Movie], error)
}

// repository é nosso adapter do banco de dados.
type MovieRepository interface {
	CreateMovie(req *dto.MovieRequest) (*Movie, error)
	DeleteMovie(req *dto.MovieRequest) (*Movie, error)
	GetAllMovies(req *dto.MovieRequest) (*Movie, error)
	GetMovieById(req *dto.MovieRequest) (*Movie, error)
	UpdateMovie(req *dto.MovieRequest) (*Movie, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination[[]Movie], error)
}

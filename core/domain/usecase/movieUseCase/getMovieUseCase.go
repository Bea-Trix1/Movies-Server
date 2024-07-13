package movieusecase

import (
	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (usecase usecase) GetAllMovies(movieRequest *dto.MovieRequest) (*domain.Movie, error) {
	movie, err := usecase.repository.GetAllMovies(movieRequest)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (usecase usecase) GetMovieById(movieRequest *dto.MovieRequest) (*domain.Movie, error) {
	movie, err := usecase.repository.GetMovieById(movieRequest)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

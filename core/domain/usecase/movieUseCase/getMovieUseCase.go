package movieusecase

import (
	"github.com/Bea-Trix1/Movies-Server/core/domain"
)

func (usecase usecase) GetAllMovies() ([]domain.Movie, error) {
	movie, err := usecase.repository.GetAllMovies()

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (usecase usecase) GetMovieById(id uint32) (*domain.Movie, error) {
	movie, err := usecase.repository.GetMovieById(id)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

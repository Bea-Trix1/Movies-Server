package movieusecase

import (
	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (usecase usecase) CreateMovie(movieRequest *dto.MovieRequest) (*domain.Movie, error) {
	movie, err := usecase.repository.CreateMovie(movieRequest)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

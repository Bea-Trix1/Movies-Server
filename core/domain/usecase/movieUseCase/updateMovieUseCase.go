package movieusecase

import (
	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (usecase usecase) UpdateMovie(movieRequest *dto.MovieRequest) (*domain.Movie, error) {
	movie, err := usecase.repository.UpdateMovie(movieRequest)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

package movieusecase

import "github.com/Bea-Trix1/Movies-Server/core/domain"

type usecase struct {
	repository domain.MovieRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.MovieRepository) domain.MovieUseCase {
	return &usecase{
		repository: repository,
	}
}

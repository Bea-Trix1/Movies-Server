package movieusecase

import "github.com/Bea-Trix1/Movies-Server/core/domain"

type usecase struct {
	repository domain.MovieRepository
}

// Retorna o contrato da implementação do use case.
func New(repository domain.MovieRepository) domain.MovieUseCase {
	return &usecase{
		repository: repository,
	}
}

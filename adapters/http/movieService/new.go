package movieservice

import "github.com/Bea-Trix1/Movies-Server/core/domain"

type service struct {
	usecase domain.MovieUseCase
}

func New(usecase domain.MovieUseCase) domain.MovieService {
	return &service{
		usecase: usecase,
	}
}

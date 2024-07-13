package di

import (
	movieservice "github.com/Bea-Trix1/Movies-Server/adapters/http/movieService"
	"github.com/Bea-Trix1/Movies-Server/adapters/postgres"
	movierepository "github.com/Bea-Trix1/Movies-Server/adapters/postgres/movieRepository"
	"github.com/Bea-Trix1/Movies-Server/core/domain"
	movieusecase "github.com/Bea-Trix1/Movies-Server/core/domain/usecase/movieUseCase"
)

// Retorna o movieService com a configuração de injeção de dependencia.
func ConfigProductDI(conn postgres.PoolInterface) domain.MovieService {
	movieRepository := movierepository.New(conn)
	movieUseCase := movieusecase.New(movieRepository)
	movieService := movieservice.New(movieUseCase)

	return movieService
}

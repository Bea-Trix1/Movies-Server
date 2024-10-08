package movierepository

import (
	"github.com/Bea-Trix1/Movies-Server/adapters/postgres"
	"github.com/Bea-Trix1/Movies-Server/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// retorna o contrato das implementações do MovieRepository
func New(db postgres.PoolInterface) domain.MovieRepository {
	return &repository{
		db: db,
	}
}

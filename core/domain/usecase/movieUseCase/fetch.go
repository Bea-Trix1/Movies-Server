package movieusecase

import (
	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination[[]domain.Movie], error) {
	movie, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

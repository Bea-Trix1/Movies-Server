package movierepository

import (
	"context"

	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
	"github.com/booscaaa/go-paginate/paginate"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParms) (*domain.Pagination[[]domain.Movie], error) {
	ctx := context.Background()
	movies := []domain.Movie{}
	total := int32(0)

	pagin := paginate.Instance(movies)
	query, queryCount := pagin.
		Query("SELECT * FROM movie").
		Sort([]string{"title", "gender"}).
		Desc([]string{"true", "false"}).
		Page(3).
		RowsPerPage(3).
		SearchBy("title", "gender").
		Select()

	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			movie := domain.Movie{}

			rows.Scan(
				&movie.ID,
				&movie.Title,
				&movie.Gender,
				&movie.Year,
				&movie.Director,
			)

			movies = append(movies, movie)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination[[]domain.Movie]{
		Items: movies,
		Total: total,
	}, nil
}

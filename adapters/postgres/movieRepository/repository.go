package movierepository

import (
	"context"

	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (repository repository) CreateMovie(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO movie (title, gender, year, director)  VALUES ($1, $2, $3, $4) returning *",
		movie.Title,
		movie.Gender,
		movie.Year,
		movie.Director,
	).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Gender,
		&movie.Year,
		&movie.Director,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repository repository) DeleteMovie(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{}

	err := repository.db.QueryRow(
		ctx,
		"DELETE FROM movie WHERE id = ?",
		movie.ID,
	).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Gender,
		&movie.Year,
		&movie.Director,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repository repository) GetAllMovies(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{}

	err := repository.db.QueryRow(
		ctx,
		"SELECT * FROM movie",
		movie.Title,
		movie.Gender,
		movie.Year,
		movie.Director,
	).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Gender,
		&movie.Year,
		&movie.Director,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repository repository) GetMovieById(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{}

	err := repository.db.QueryRow(
		ctx,
		"SELECT * FROM movie WHERE id = ?",
		movie.ID,
	).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Gender,
		&movie.Year,
		&movie.Director,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repository repository) UpdateMovie(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{}

	err := repository.db.QueryRow(
		ctx,
		"UPDATE movie SET title = ?, gender = ?, year = ?, director = ?",
		movie.Title,
		movie.Gender,
		movie.Year,
		movie.Director,
	).Scan(
		&movie.Title,
		&movie.Gender,
		&movie.Year,
		&movie.Director,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

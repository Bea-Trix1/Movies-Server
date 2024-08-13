package movierepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (repository repository) CreateMovie(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{
		Title:    req.Title,
		Gender:   req.Gender,
		Year:     req.Year,
		Director: req.Director,
	}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO movie (title, gender, year, director) VALUES ($1, $2, $3, $4) RETURNING id",
		movie.Title,
		movie.Gender,
		movie.Year,
		movie.Director,
	).Scan(&movie.ID)

	if err != nil {
		return nil, fmt.Errorf("falha ao criar novo filme: %w", err)
	}

	return &movie, nil
}

func (repository repository) DeleteMovie(id uint32) error {
	ctx := context.Background()

	_, err := repository.db.Exec(
		ctx,
		"DELETE FROM movie WHERE id = $1",
		id,
	)

	if err != nil {
		return fmt.Errorf("falha ao deletar filme com id %d: %w", id, err)
	}

	return nil
}

func (repository repository) GetAllMovies() ([]domain.Movie, error) {
	ctx := context.Background()
	rows, err := repository.db.Query(ctx, "SELECT id, title, gender, year, director FROM movie")
	if err != nil {
		return nil, fmt.Errorf("falha ao buscar todos os filmes: %w", err)
	}
	defer rows.Close()

	var movies []domain.Movie
	for rows.Next() {
		var movie domain.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Gender, &movie.Year, &movie.Director); err != nil {
			return nil, fmt.Errorf("falha ao buscar filmes: %w", err)
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao filtrar linhas: %w", err)
	}

	return movies, nil
}

func (repository repository) GetMovieById(id uint32) (*domain.Movie, error) {
	ctx := context.Background()

	var movie domain.Movie

	err := repository.db.QueryRow(
		ctx,
		"SELECT id, title, gender, year, director FROM movie WHERE id = $1",
		id,
	).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Gender,
		&movie.Year,
		&movie.Director,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("filme não encontrado")
		}
		return nil, err
	}
	return &movie, nil
}

func (repository repository) UpdateMovie(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{
		Title:    req.Title,
		Gender:   req.Gender,
		Year:     req.Year,
		Director: req.Director,
	}

	err := repository.db.QueryRow(
		ctx,
		"UPDATE movie SET title = $1, gender = $2, year = $3, director = $4 WHERE id = $5 RETURNING id, title, gender, year, director",
		movie.ID,
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
		log.Fatal("Falha ao atualizar filme.", err)
	}

	return &movie, nil
}

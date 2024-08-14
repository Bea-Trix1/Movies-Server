package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Bea-Trix1/Movies-Server/adapters/postgres"
	"github.com/Bea-Trix1/Movies-Server/di"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	movieService := di.ConfigProductDI(conn)

	router := mux.NewRouter()
	router.Handle("/createMovie", http.HandlerFunc(movieService.CreateMovie)).Methods("POST")
	router.Handle("/deleteMovieById", http.HandlerFunc(movieService.DeleteMovie)).Methods("DELETE")
	router.Handle("/getAllMovies", http.HandlerFunc(movieService.GetAllMovies)).Methods("GET")
	router.Handle("/getMovieById", http.HandlerFunc(movieService.GetMovieById)).Methods("GET")
	router.Handle("/updateMovieById", http.HandlerFunc(movieService.UpdateMovie)).Methods("PUT")
	router.Handle("/movie", http.HandlerFunc(movieService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}

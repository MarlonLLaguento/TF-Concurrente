package main

import (
	"fmt"
	"log"
)

func main() {
	// Cargar el dataset completo
	movieMap, err := loadDataset("ml_10m/movies.dat", "ml_10m/ratings_copia.dat")
	if err != nil {
		log.Fatalf("Error cargando el dataset: %v\n", err)
	}

	// Convertir el mapa de películas en un slice de películas
	movies := make([]Movie, 0, len(movieMap))
	for _, movie := range movieMap {
		movies = append(movies, movie)
	}

	// Probar el sistema con diferentes movieIDs
	testMovieIDs := []int{1193, 260, 589, 2571, 480, 527}
	for _, movieID := range testMovieIDs {
		fmt.Printf("\nProbando el sistema de recomendación para la película ID %d:\n", movieID)
		startClient(movieID, movies)
	}
}

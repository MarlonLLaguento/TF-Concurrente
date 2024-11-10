package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Cargar los datos de películas y calificaciones
	movies, err := loadDataset("ml_10m/movies.dat", "ml_10m/ratings.dat")
	if err != nil {
		log.Fatalf("Error cargando el dataset: %v\n", err)
	}

	// Solicitar al usuario un ID de película
	fmt.Print("Por favor, ingresa el ID de la película: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Convertir el input a un número entero
	movieID, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID de película inválido.")
		return
	}

	// Buscar el nombre de la película y generar recomendaciones
	var targetMovie *Movie
	for _, movie := range movies {
		if movie.ID == movieID {
			targetMovie = &movie
			break
		}
	}

	if targetMovie == nil {
		fmt.Printf("Película con ID %d no encontrada.\n", movieID)
		return
	}

	fmt.Printf("Película seleccionada: %s\n", targetMovie.Title)
	recommendations := generateItemBasedRecommendations(*targetMovie, movies)
	fmt.Printf("Recomendaciones para la película %d - %s: %v\n", targetMovie.ID, targetMovie.Title, recommendations)
}


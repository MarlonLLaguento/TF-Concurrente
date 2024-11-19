package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Definición de la estructura Movie
type Movie struct {
	ID          int
	Title       string
	Genres      []string
	AvgRating   float64
	NumRatings  int
	UserRatings map[int]float64
}

// Función para cargar el dataset de películas y calificaciones
func loadDataset(moviePath, ratingPath string) ([]Movie, error) {
	movieFile, err := os.Open(moviePath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo el archivo de películas: %v", err)
	}
	defer movieFile.Close()

	ratingFile, err := os.Open(ratingPath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo el archivo de calificaciones: %v", err)
	}
	defer ratingFile.Close()

	movieMap := make(map[int]Movie)
	scanner := bufio.NewScanner(movieFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "::")
		if len(parts) < 3 {
			continue
		}
		id, _ := strconv.Atoi(parts[0])
		title := parts[1]
		genres := strings.Split(parts[2], "|")
		movieMap[id] = Movie{
			ID:          id,
			Title:       title,
			Genres:      genres,
			UserRatings: make(map[int]float64),
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error leyendo el archivo de películas: %v", err)
	}

	scanner = bufio.NewScanner(ratingFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "::")
		if len(parts) < 4 {
			continue
		}
		userID, _ := strconv.Atoi(parts[0])
		movieID, _ := strconv.Atoi(parts[1])
		rating, _ := strconv.ParseFloat(parts[2], 64)

		if movie, exists := movieMap[movieID]; exists {
			movie.UserRatings[userID] = rating
			movie.NumRatings++
			movie.AvgRating += rating
			movieMap[movieID] = movie
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error leyendo el archivo de calificaciones: %v", err)
	}

	// Calcular promedio de calificación
	movies := make([]Movie, 0, len(movieMap))
	for _, movie := range movieMap {
		if movie.NumRatings > 0 {
			movie.AvgRating /= float64(movie.NumRatings)
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

// Función de similitud de coseno
func cosineSimilarity(ratings1, ratings2 map[int]float64) float64 {
	var dotProduct, normA, normB float64
	for userID, ratingA := range ratings1 {
		if ratingB, exists := ratings2[userID]; exists {
			dotProduct += ratingA * ratingB
			normA += ratingA * ratingA
			normB += ratingB * ratingB
		}
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

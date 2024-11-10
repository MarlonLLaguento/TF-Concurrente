package main

import (
	"fmt"
	"sort"
)

// Función que genera recomendaciones basadas en ítems
func generateItemBasedRecommendations(targetMovie Movie, allMovies []Movie) []int {
	type recommendation struct {
		id         int
		similarity float64
	}
	var recommendations []recommendation

	for _, movie := range allMovies {
		if movie.ID != targetMovie.ID {
			similarity := cosineSimilarity(targetMovie.UserRatings, movie.UserRatings)
			if similarity > 0.3 { // Umbral de similitud
				recommendations = append(recommendations, recommendation{movie.ID, similarity})
			}
		}
	}

	// Ordenar recomendaciones por similitud
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].similarity > recommendations[j].similarity
	})

	// Limitar a las 10 recomendaciones principales
	topRecommendations := []int{}
	for i, rec := range recommendations {
		if i >= 10 {
			break
		}
		topRecommendations = append(topRecommendations, rec.id)
	}

	return topRecommendations
}

// Función que inicia el cliente y devuelve recomendaciones
func startClient(movieID int, allMovies []Movie) []int {
	var targetMovie *Movie
	for _, movie := range allMovies {
		if movie.ID == movieID {
			targetMovie = &movie
			break
		}
	}

	if targetMovie == nil {
		fmt.Printf("Película con ID %d no encontrada.\n", movieID)
		return []int{}
	}

	recommendations := generateItemBasedRecommendations(*targetMovie, allMovies)
	fmt.Printf("Recomendaciones para la película %d: %v\n", targetMovie.ID, recommendations)
	return recommendations
}

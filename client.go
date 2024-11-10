package main

import (
	"fmt"
	"math"
	"sort"
)

const similarityThreshold = 0.5
const maxRecommendations = 10

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

func generateItemBasedRecommendations(targetMovie Movie, allMovies []Movie) []int {
	type recommendation struct {
		id         int
		similarity float64
	}
	var recommendations []recommendation

	for _, movie := range allMovies {
		if movie.ID != targetMovie.ID {
			similarity := cosineSimilarity(targetMovie.UserRatings, movie.UserRatings)
			if similarity > similarityThreshold {
				recommendations = append(recommendations, recommendation{movie.ID, similarity})
			}
		}
	}

	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].similarity > recommendations[j].similarity
	})

	topRecommendations := []int{}
	for i, rec := range recommendations {
		if i >= maxRecommendations {
			break
		}
		topRecommendations = append(topRecommendations, rec.id)
	}

	return topRecommendations
}

func startClient(movieID int, allMovies []Movie) {
	var targetMovie *Movie
	for _, movie := range allMovies {
		if movie.ID == movieID {
			targetMovie = &movie
			break
		}
	}

	if targetMovie == nil {
		fmt.Printf("Película con ID %d no encontrada.\n", movieID)
		return
	}

	recommendations := generateItemBasedRecommendations(*targetMovie, allMovies)
	fmt.Printf("Recomendaciones para la película %d: %v\n", targetMovie.ID, recommendations)
}

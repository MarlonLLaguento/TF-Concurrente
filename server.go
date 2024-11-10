package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

func startServer(port string) error {
	// Inicia el servidor en el puerto especificado
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("Servidor iniciado en el puerto", port)

	// Cargar dataset completo (una sola vez en el servidor)
	movieMap, err := loadDataset("ml_10m/movies.dat", "ml_10m/ratings.dat")
	if err != nil {
		return fmt.Errorf("error cargando el dataset: %v", err)
	}

	// Convertir el mapa de películas en un slice de películas
	movies := make([]Movie, 0, len(movieMap))
	for _, movie := range movieMap {
		movies = append(movies, movie)
	}

	// Dividir el dataset entre los clientes
	clientDatasets := splitDataset(movies, 3) // Dividir en 3 partes (por ejemplo)

	clientIndex := 0
	for {
		// Acepta conexiones de clientes
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}
		fmt.Printf("Cliente %d conectado: %v\n", clientIndex, conn.RemoteAddr())

		// Enviar la parte correspondiente del dataset al cliente
		go handleClient(conn, clientDatasets[clientIndex])
		clientIndex++
	}
}

// Función para dividir el dataset en partes iguales
func splitDataset(movies []Movie, numParts int) [][]Movie {
	var divided [][]Movie
	chunkSize := (len(movies) + numParts - 1) / numParts

	for i := 0; i < len(movies); i += chunkSize {
		end := i + chunkSize
		if end > len(movies) {
			end = len(movies)
		}
		divided = append(divided, movies[i:end])
	}

	return divided
}

func handleClient(conn net.Conn, clientData []Movie) {
	defer conn.Close()

	// Serializar los datos antes de enviarlos
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(clientData)
	if err != nil {
		fmt.Println("Error al codificar datos para el cliente:", err)
		return
	}

	// Enviar los datos al cliente
	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		fmt.Println("Error al enviar datos al cliente:", err)
		return
	}

	fmt.Println("Datos enviados al cliente")
}

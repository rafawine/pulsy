package main

import (
	"fmt"
	"log"
	"os"

	"pulsy/internal/firebase"
	"pulsy/internal/routes"
)

func main() {
	// Obtener puerto de ejecución
	port := os.Getenv("PULSY_PORT")

	// Inicializar Firebase
	firebase.InitializeFirebase()
	firebase.InitializeStorage()
	firebase.InitializeFirestore()

	// Cerrar Firestore al finalizar
	defer firebase.CloseFirestore()

	// Configurar servidor Gin
	router := routes.SetupRouter()

	// Ejecutar servidor
	err := router.Run(fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

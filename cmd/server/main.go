package main

import (
	"fmt"
	"log"
	"os"

	"pulsy/internal/firebase"
	"pulsy/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Obtener puerto de ejecución
	port := os.Getenv("PORT")

	if port == "" {
		port = "4321"
	}

	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Inicializar Firebase
	firebase.InitializeFirebase()
	firebase.InitializeStorage()
	firebase.InitializeFirestore()

	// Cerrar Firestore al finalizar
	defer firebase.CloseFirestore()

	// Configurar servidor Gin
	router := routes.SetupRouter()

	// Ejecutar servidor
	err = router.Run(fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

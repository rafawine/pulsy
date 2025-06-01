package main

import (
	"log"

	"pulsy/internal/firebase"
	"pulsy/internal/routes"
)

func main() {
	// Inicializar Firebase
	firebase.InitializeFirebase()
	firebase.InitializeStorage()
	firebase.InitializeFirestore()
	firebase.InitializeAuthentication()

	// Cerrar Firestore al finalizar
	defer firebase.CloseFirestore()

	// Configurar servidor Gin
	router := routes.SetupRouter()

	// Ejecutar servidor
	err := router.Run(":4321")

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

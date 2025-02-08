package firebase

import (
	"log"

	"firebase.google.com/go/v4/auth"
)

var AuthenticationClient *auth.Client

func InitializeAuthentication() {
	if FirebaseApp == nil {
		log.Fatal("firebaseApp is not initialized")
	}

	client, err := FirebaseApp.Auth(GetNewContext())
	if err != nil {
		log.Fatalf("error getting firebase authentication client: %v", err)
	}

	AuthenticationClient = client
}

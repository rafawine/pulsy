package firebase

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitializeFirebase() {
	// Crear el objeto de credenciales desde las variables de entorno
	credentials := map[string]string{
		"type":                        os.Getenv("FIREBASE_TYPE"),
		"project_id":                  os.Getenv("FIREBASE_PROJECT_ID"),
		"private_key_id":              os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		"private_key":                 os.Getenv("FIREBASE_PRIVATE_KEY"),
		"client_email":                os.Getenv("FIREBASE_CLIENT_EMAIL"),
		"client_id":                   os.Getenv("FIREBASE_CLIENT_ID"),
		"auth_uri":                    os.Getenv("FIREBASE_AUTH_URI"),
		"token_uri":                   os.Getenv("FIREBASE_TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("FIREBASE_AUTH_PROVIDER_X509"),
		"client_x509_cert_url":        os.Getenv("FIREBASE_CLIENT_X509"),
		"universe_domain":             os.Getenv("FIREBASE_UNIVERSE_DOMAIN"),
	}

	// Convertir las credenciales a JSON
	credentialsJSON, err := json.Marshal(credentials)
	if err != nil {
		log.Fatalf("Error creating Firebase credentials JSON: %v", err)
	}

	// Inicializar Firebase
	opt := option.WithCredentialsJSON(credentialsJSON)

	app, err := firebase.NewApp(GetNewContext(), nil, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	FirebaseApp = app
}

func GetNewContext() context.Context {
	return context.Background()
}

func GetNewContextWithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

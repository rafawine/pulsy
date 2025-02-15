package firebase

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitializeFirebase() {
	// Crear el objeto de credenciales desde las variables de entorno
	credentials := map[string]string{
		"type":                        os.Getenv("TYPE"),
		"project_id":                  os.Getenv("PROJECT_ID"),
		"private_key_id":              os.Getenv("PRIVATE_KEY_ID"),
		"private_key":                 strings.ReplaceAll(os.Getenv("PRIVATE_KEY"), "\\n", "\n"),
		"client_email":                os.Getenv("CLIENT_EMAIL"),
		"client_id":                   os.Getenv("CLIENT_ID"),
		"auth_uri":                    os.Getenv("AUTH_URI"),
		"token_uri":                   os.Getenv("TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("AUTH_PROVIDER_CERT_URL"),
		"client_x509_cert_url":        os.Getenv("CLIENT_CERT_URL"),
		"universe_domain":             os.Getenv("UNIVERSE_DOMAIN"),
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

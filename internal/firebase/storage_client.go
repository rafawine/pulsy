package firebase

import (
	"log"

	"firebase.google.com/go/v4/storage"
)

var StorageClient *storage.Client

type QueryCondition struct {
	Field    string
	Operator string
	Value    interface{}
}

func InitializeStorage() {
	if FirebaseApp == nil {
		log.Fatalf("firebaseApp is not initialized")
	}

	client, err := FirebaseApp.Storage(GetNewContext())
	if err != nil {
		log.Fatalf("error getting firebase storage client: %v", err)
	}

	StorageClient = client
}

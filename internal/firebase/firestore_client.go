package firebase

import (
	"log"

	"cloud.google.com/go/firestore"
)

var FirestoreClient *firestore.Client

func InitializeFirestore() {
	if FirebaseApp == nil {
		log.Fatalf("firebaseApp is not initialized")
	}

	client, err := FirebaseApp.Firestore(GetNewContext())
	if err != nil {
		log.Fatalf("firebaseApp is not initialized")
	}

	FirestoreClient = client
}

func CloseFirestore() {
	if FirestoreClient != nil {
		FirestoreClient.Close()
	}
}

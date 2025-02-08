package services

import (
	"pulsy/internal/firebase"

	"firebase.google.com/go/v4/auth"
)

func GetUser(uid string) (*auth.UserRecord, error) {
	return firebase.AuthenticationClient.GetUser(firebase.GetNewContext(), uid)
}

func GerUserByEmail(email string) (*auth.UserRecord, error) {
	return firebase.AuthenticationClient.GetUserByEmail(firebase.GetNewContext(), email)
}

func GetToken(uid string) (string, error) {
	return firebase.AuthenticationClient.CustomToken(firebase.GetNewContext(), uid)
}

func VerifyToken(idToken string) (*auth.Token, error) {
	return firebase.AuthenticationClient.VerifyIDToken(firebase.GetNewContext(), idToken)
}

package graph

import (
	"cloud.google.com/go/firestore"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	FirestoreClient *firestore.Client
}

func NewResolver(client *firestore.Client) *Resolver {
	return &Resolver{
		FirestoreClient: client,
	}
}

package datastore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func NewFirestoreClient(projId string) (*firestore.Client, error) {
	ctx := context.Background()
	//config := &firebase.Config{ProjectID: projId}

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

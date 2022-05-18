package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"apptrack/graph/generated"
	"apptrack/graph/model"
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"cloud.google.com/go/firestore"
)

func (r *mutationResolver) CreateApplication(ctx context.Context, input model.ApplicationInput, userID string) (*model.Application, error) {
	if input.Company == nil || input.Role == nil || input.Status == nil { // Necessities to create a new application.
		return nil, errors.New("Company, Role, and Status must be provided")
	}

	users := r.FirestoreClient.Collection("users")
	_, err := users.Doc(userID).Get(ctx) // Check to make sure user exists

	if err != nil {
		return nil, err
	}

	apps := users.Doc(userID).Collection("applications")

	timeNow := fmt.Sprintf("%v", time.Now().UnixMilli())

	m := model.Application{
		Company:     *input.Company,
		Role:        *input.Role,
		Status:      *input.Status,
		Location:    input.Location,
		DateApplied: input.DateApplied,
		DateUpdated: &timeNow,
		Notes:       input.Notes,
	}
	newDoc := apps.NewDoc()

	_, err = newDoc.Create(ctx, m)

	if err != nil {
		return nil, err
	}

	m.ID = newDoc.ID

	return &m, nil
}

func (r *mutationResolver) UpdateApplication(ctx context.Context, userID string, appID string, input model.ApplicationInput) (*model.Application, error) {
	users := r.FirestoreClient.Collection("users")
	_, err := users.Doc(userID).Get(ctx) // Check to make sure user exists

	if err != nil {
		return nil, err
	}

	apps := users.Doc(userID).Collection("applications")
	_, err = apps.Doc(appID).Get(ctx) // Check to make sure application exists

	if err != nil {
		return nil, err
	}

	appDoc := apps.Doc(appID) // Get document reference

	batch := r.FirestoreClient.Batch() // Get a new write batch

	// Iterate over each value we have in the given struct
	v := reflect.ValueOf(input)

	// For each field in the struct, batch a update on the doc for that corresponding field if the field is populated.
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldVal := v.Field(i).Interface()

		if fieldVal != nil {
			batch.Update(appDoc, []firestore.Update{
				{Path: fieldName, Value: fieldVal},
			})
		}
	}

	// Update time updated...
	batch.Update(appDoc, []firestore.Update{
		{Path: "DateUpdated", Value: fmt.Sprintf("%v", time.Now().UnixMilli())},
	})

	_, err = batch.Commit(ctx)

	if err != nil {
		return nil, err
	}

	var m model.Application

	doc, err := apps.Doc(appID).Get(ctx) // Get the updated document TODO: Do we really need to do it like this??
	err = doc.DataTo(&m)

	m.ID = appID

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *mutationResolver) DeleteApplication(ctx context.Context, userID string, appID string) (*model.Application, error) {
	users := r.FirestoreClient.Collection("users")
	_, err := users.Doc(userID).Get(ctx) // Check to make sure user exists

	if err != nil {
		return nil, err
	}

	apps := users.Doc(userID).Collection("applications")
	appDocRef := apps.Doc(appID)
	doc, err := appDocRef.Get(ctx) // Check to make sure application exists

	if err != nil {
		return nil, err
	}

	var m model.Application
	err = doc.DataTo(&m)
	m.ID = appID

	_, err = appDocRef.Delete(ctx) // Delete the document

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, userID string, input model.UserInput) (*model.User, error) {
	users := r.FirestoreClient.Collection("users")
	usrDoc := users.Doc(userID)
	_, err := usrDoc.Get(ctx) // Check if user already exists in DB...
	if err == nil {           // If err is not set, that means the document already exists.
		return nil, errors.New("User already exists in database.")
	}

	u := model.User{
		ID:    userID,
		Email: input.Email,
	}
	_, err = usrDoc.Set(ctx, &u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *queryResolver) Applications(ctx context.Context, userID string) ([]*model.Application, error) {
	users := r.FirestoreClient.Collection("users")
	_, err := users.Doc(userID).Get(ctx) // Check to make sure user exists

	if err != nil {
		return nil, err
	}

	apps := users.Doc(userID).Collection("applications")
	appDocs, err := apps.Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	appList := make([]*model.Application, len(appDocs))

	for i, v := range appDocs {
		err = v.DataTo(&appList[i])

		if err != nil {
			return nil, err
		}

		appList[i].ID = v.Ref.ID
	}

	return appList, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

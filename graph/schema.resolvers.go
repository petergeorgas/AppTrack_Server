package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"apptrack/graph/generated"
	"apptrack/graph/model"
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
)

func (r *mutationResolver) CreateApplic
 ation(ctx context.Context, input model.ApplicationInput, userID string) (*model.Application, error) {
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

func (r *mutationResolver) UpdateApplication(ctx context.Context, userID string, appID string, status model.Status) (*model.Application, error) {
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

	appDoc := apps.Doc(appID)
	_, err = appDoc.Update(ctx, []firestore.Update{
		{Path: "Status", Value: status},
		{Path: "DateUpdated", Value: fmt.Sprintf("%v", time.Now().UnixMilli())},
	})

	if err != nil {
		return nil, err
	}

	var m model.Application

	doc, err := apps.Doc(appID).Get(ctx)
	err = doc.DataTo(&m)

	m.ID = appID

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	// TODO: Check if account already exists for email
	users := r.FirestoreClient.Collection("users")

	u := model.User{
		Email: input.Email,
	}
	newDoc := users.NewDoc()

	_, err := newDoc.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	u.ID = newDoc.ID

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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetApplications(ctx context.Context, id string) ([]*model.Application, error) {
	panic(fmt.Errorf("not implemented"))
}

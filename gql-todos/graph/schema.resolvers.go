package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"gql-tools/graph/generated"
	"gql-tools/graph/model"
	"gql-tools/graph/utils"
	"strconv"
)

func (r *mutationResolver) UpsertUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	id := input.ID
	var user model.User
	user.IsAdmin = input.IsAdmin
	n := len(r.Resolver.UserStore)
	if n == 0 {
		r.Resolver.UserStore = make(map[string]model.User)
	}
	if id != nil {
		new_user, ok := r.Resolver.UserStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		r.Resolver.UserStore[*id] = new_user
	} else {
		nid := strconv.Itoa(n + 1)
		user.ID = nid
		r.Resolver.UserStore[nid] = user
	}
	return &user, nil
}

func (r *mutationResolver) GenerateBlock(ctx context.Context, input model.BlockInput) (*model.Block, error) {
	id := input.ID
	var block model.Block
	block.Contents = input.Contents
}

func (r *mutationResolver) ValidateKey(ctx context.Context, apiKey string) (*model.Token, error) {
	if !utils.DecodeAPIKeys(apiKey) {
		return nil, errors.New("incorrrect key")
	}

	token := &model.Token{
		Token: utils.CreateJWT(),
	}

	return token, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, ok := r.Resolver.UserStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &user, nil
}

func (r *queryResolver) Blocks(ctx context.Context) ([]*model.Block, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

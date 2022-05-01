package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"gql-tools/graph/generated"
	"gql-tools/graph/middlewares"
	"gql-tools/graph/model"
	"gql-tools/graph/utils"
	"strconv"
)

func (r *mutationResolver) UpsertUser(ctx context.Context, input model.UserInput) (*model.Token, error) {
	var user model.User

	var token model.Token
	user.IsAdmin = input.IsAdmin
	user.Name = input.Name
	created_token, err := utils.CreateJWT(user.Name)
	token.Token = created_token
	if err != nil {
		return &token, fmt.Errorf(err.Error())
	}
	n := len(r.Resolver.UserStore)
	if n == 0 {
		r.Resolver.UserStore = make(map[string]model.User)
	} else {
		nid := strconv.Itoa(n + 1)
		user.ID = nid
		r.Resolver.UserStore[nid] = user
	}
	return_token := model.Token{
		Token: token.Token,
		User:  &user,
	}
	return &return_token, nil
}

func (r *mutationResolver) GenerateBlock(ctx context.Context, input model.BlockInput) (*model.Block, error) {
	user := middlewares.GetAuthFromContext(ctx)
	fmt.Println(user)
	if user == nil {
		return &model.Block{}, fmt.Errorf("access denied")
	}
	id := input.ID
	var block model.Block
	block.Contents = input.Contents
	n := len(r.Resolver.BlockStore)
	if n == 0 {
		r.Resolver.BlockStore = make(map[string]model.Block)
	}
	if id != nil {
		updated_block, ok := r.Resolver.BlockStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		r.Resolver.BlockStore[*id] = updated_block
	} else {
		nid := strconv.Itoa(n + 1)
		block.ID = nid
		r.Resolver.BlockStore[nid] = block
	}
	return &block, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	var user model.User
	user.Name = *input.Name
	user.Password = *input.Password
	auth := utils.DecodeAPIKeys(user.Password)
	if !auth {
		return "", errors.New("not auth")
	}
	token, err := utils.CreateJWT(user.Name)
	if err != nil {
		return "", err
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
	blocks := make([]*model.Block, 0)
	for idx := range r.Resolver.BlockStore {
		block := r.Resolver.BlockStore[idx]
		blocks = append(blocks, &block)
	}

	return blocks, nil
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
func (r *mutationResolver) ValidateKey(ctx context.Context, apiKey string) (*model.Token, error) {
	panic("blah")
}

package graph

import "github.com/kimj99/gql-tools/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CharacterStore map[string]model.Character
	UserStore      map[string]model.User
}

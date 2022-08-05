package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"math/rand"

	"codecentric.de/demo/graphql-schema-first-fridge/graph/generated"
	"codecentric.de/demo/graphql-schema-first-fridge/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "u" + input.UserID,
		},
	}

	r.todos = append(r.todos, todo)
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, done *bool) ([]*model.Todo, error) {
	return r.todos, nil
}

// Beers is the resolver for the beers field.
func (r *queryResolver) Beers(ctx context.Context, minPercercentage *float64) ([]*model.Beer, error) {
	beersFiltered := lo.Filter[*model.Beer](r.BeerResolver(), func(it *model.Beer, _ int) bool {
		return it.Percentage >= *minPercercentage
	})
	// return r.BeerResolver(), nil
	return beersFiltered, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"codecentric.de/demo/graphql-schema-first-fridge/graph/generated"
	"codecentric.de/demo/graphql-schema-first-fridge/graph/model"
	"github.com/samber/lo"
)

// Beers is the resolver for the beers field.
func (r *queryResolver) Beers(_ context.Context, minPercentage *float64) ([]*model.Beer, error) {
	if *minPercentage < 0.0 {
		return nil, errors.New("percentage must be bigger or equal to 0")
	}

	beersFiltered := lo.Filter[*model.Beer](r.BeerResolver(), func(it *model.Beer, _ int) bool {
		return it.Percentage >= *minPercentage
	})

	return beersFiltered, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

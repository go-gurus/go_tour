## GraphQL

Lets be a hipster for now.

----

### GraphQL in Go

Approaches

- Code first
- Schema first

----

### GraphQLGen

- Schema-First
    - Generates Go types and resolvers
    - Integrates into present solutions

----

### Beer Fridge Example

Objectives:

- same functionality as gin/chi variant
- implement a percentage filter

----

#### Generate Skeleton Project

```bash
go run github.com/99designs/gqlgen init
```

----

#### Write a simple beer schema

```graphql
type Beer {
    id: ID!
    manufacturer: String!
    name: String!
    origin: String!
    type: String!
    percentage : Float!
    ibu: Int
}
```

<small>graph/beer.graphqls</small>

----

#### Provide a query specification

```graphql
type Query {
    beers: [Beer!]!
}
```

<small>graph/schema.graphqls</small>

----

#### Generate Go code from schema

```bash
go run github.com/99designs/gqlgen generate
```

----

#### Lets stick with a fake data source for now

```golang
type Resolver struct {
BeerResolver func () []*model.Beer
}
```

<small>graph/resolver.go</small>

```golang
func GetBeers() []*model.Beer {
return funnyFakeBeerList
}
```

<small>service/BeerService.go</small>


----

#### Implement resolver

```go
// Beers is the resolver for the beers field.
func (r *queryResolver) Beers(ctx context.Context) ([]*model.Beer, error) {
return r.BeerResolver(), nil
}
```

<small>graph/schema.resolvers.go</small>

----

#### Our first GraphQL Query

```graphql
query{
    beers {
        id
        name
        manufacturer
        ibu
        percentage
    }
}
```

```json
{
  "data": {
    "beers": [
      {
        "id": "OMTR",
        "name": "Oostalle Trappist",
        "manufacturer": "Oostmalle",
        "ibu": 48,
        "percentage": 14.1
      } // more...
```

----

#### We care for certain properties only..

```graphql
query{
    beers {
        origin
        manufacturer
        percentage
    }
}
```

```json
"data": {
"beers": [
{
"origin": "BE",
"manufacturer": "Oostmalle",
"percentage": 14.1
}, ///more
```

----
<big>Wait..</big>
----
<big>Haven't we mentioned filtering before?</big>
----

#### Provide a query with a filter

1. Update Query schema

```graphql
type Query {
    beers(minPercercentage: Float = 0.0) : [Beer!]!
}
```

<small>graph/schema.graphqls</small>

2. Re-Run code generator

```bash
go run github.com/99designs/gqlgen generate
```

----

#### Provide a query with a filter

3. Implement resolver

```golang
func (r *queryResolver) Beers(ctx context.Context, 
	minPercercentage *float64) ([]*model.Beer, error) {
	// In practice, this would happen in the filter. Who wants to implement it?
    beersFiltered := lo.Filter[*model.Beer](r.BeerResolver(), 
		func(it *model.Beer, _ int) bool {
            return it.Percentage >= *minPercercentage
    })
}
```
<small>graph/schema.resolvers.go</small>

----

#### Query with filters

Return all beers with >10% and where they are from.
```graphql
query{
  beers(minPercercentage : 10) {
    name
    origin
    manufacturer
    percentage
  }
}
```

```json
{
  "data": {
    "beers": [
      {
        "name": "Rituel Quatorze",
        "origin": "BE",
        "manufacturer": "Grim Fandango",
        "percentage": 14.9
      }, // more
    ]
  }
}
```

----


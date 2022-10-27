## GraphQL

Our beer fridge provides HTTP/REST style APIs now. With Go, it is also possible to provide a GraphQL interface. Lets dive into
it.

----

### GraphQL in Go

Approaches

* Code first
* Schema first

----

### Code first

- Generate GraphQL scheme from Go types
- Define relations by type relations and annotations

----

### Schema first

- Generate Go classes and Queries from GraphQL scheme

----

### Schema-first with GraphQLGen

----

### Micro Project: Beer Fridge

As a thirsty employee, I want:

- the same functionality as the REST API variant so I don't need two APIs
- a built in filter for beer IBU and origin to discover new beers

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

#### Write a simple beer schema

##### GraphQL type documentation

Of course the GraphQL type documentation is passed through to the generated Go code

```graphql
"""
Beer defines key criteria of a beer
"""
type Beer {
    id: ID!
    manufacturer: String!
    name: String!
    """
    Origin of the beer as ISO country code
    """
    origin: String!
    type: String!
    percentage : Float!
    ibu: Int
}
```

<small>graph/beer.graphqls</small>

----

#### Provide a query specification

To expose data to clients, we need a Query specification.

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
#### How to issue a GraphQL Query

#### Web Interface
GraphiQL Web Interface: http://localhost:8080

#### Direct HTTP request 
```bash
curl 'http://localhost:8080/query' \  
  --data-raw '{"query":"..."'} \
```
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
        "name": "Oostmalle Trappist",
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
<big>Job done.</big>
----
<big>Wait..</big>
----
<big>Haven't we mentioned filtering before?</big>
----

#### Provide a query with a filter

1. Update Query schema

```graphql
type Query {
    beers(minPercentage: Float = 0.0) : [Beer!]!
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
func (r *queryResolver) Beers(_ context.Context, 
	minPercentage *float64) ([]*model.Beer, error) {
    if *minPercentage < 0.0 {
        return nil, errors.New("percentage must be bigger or equal to 0")
    }
    
    beersFiltered := lo.Filter[*model.Beer](r.BeerResolver(), 
		func(it *model.Beer, _ int) bool {
        return it.Percentage >= *minPercentage
    })
    
    return beersFiltered, nil
}
```

<small>graph/schema.resolvers.go</small>

----

#### Query with filters

Return all beers with >10% and where they are from.

```graphql
query{
    beers(minPercentage : 10) {
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
      }
      // more
    ]
  }
}
```

----

### Further readings

* [https://gqlgen.com/getting-started/(https://gqlgen.com/getting-started/)
* [https://graphql.org/](https://graphql.org/)
---

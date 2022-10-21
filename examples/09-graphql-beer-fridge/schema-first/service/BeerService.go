package service

import "codecentric.de/demo/graphql-schema-first-fridge/graph/model"

/*
type Beer struct {
	id           string
	manufacturer string
	name         string
	origin       string
	beerType     string
	percentage   float32
	ibu          float32
}*/

var beerList = []*model.Beer{
	{
		ID:           "OMTR",
		Manufacturer: "Oostmalle",
		Name:         "Oostalle Trappist",
		Origin:       "BE",
		Type:         "TRAPPIST",
		Percentage:   14.1,
	}, {
		ID:           "SCHE",
		Manufacturer: "Schängche",
		Name:         "Aix Escalation",
		Origin:       "DE",
		Type:         "IPA",
		Percentage:   7.0,
	}, {
		ID:           "BRWJ",
		Manufacturer: "Brewcat",
		Name:         "John Lennon",
		Origin:       "UK",
		Type:         "IPA",
		Percentage:   6.9,
	}, {
		ID:           "NALX",
		Manufacturer: "Napoleon",
		Name:         "Sainte Helene",
		Origin:       "FR",
		Type:         "Stout",
		Percentage:   8.1,
	}, {
		ID:           "NALX",
		Manufacturer: "Grim Fandango",
		Name:         "Rituel Quatorze",
		Origin:       "BE",
		Type:         "Quadrupel",
		Percentage:   14.9,
	},
}

func GetBeers() []*model.Beer {
	return beerList
}

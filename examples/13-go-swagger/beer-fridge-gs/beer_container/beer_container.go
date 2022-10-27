package beer_container

import (
	"codecentric.de/beer-fridge-go-swagger/models"
	"github.com/go-openapi/errors"
	"sync"
	"sync/atomic"
)

var beerList = make(map[int64]*models.Beer)
var lastID int64
var beerListLock = &sync.Mutex{}

func newBeerID() int64 {
	return atomic.AddInt64(&lastID, 1)
}

func AddBeer(beer *models.Beer) error {
	if beer == nil {
		return errors.New(500, "beer must be present")
	}

	beerListLock.Lock()
	defer beerListLock.Unlock()

	newID := newBeerID()
	beer.ID = newID
	beerList[newID] = beer

	return nil
}

func DeleteBeer(id int64) error {
	beerListLock.Lock()
	defer beerListLock.Unlock()

	_, exists := beerList[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(beerList, id)
	return nil
}

func AllBeers(limit int32) (result []*models.Beer) {
	result = make([]*models.Beer, 0)
	for _, beer := range beerList {
		if len(result) >= int(limit) {
			return
		}
		result = append(result, beer)
	}
	return
}

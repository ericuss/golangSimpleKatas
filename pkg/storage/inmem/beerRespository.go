package inmem

import (
	"fmt"
	"strconv"
	"sync"

	beer "github.com/ericuss/golangSimpleKatas/pkg"
)

type beerRepository struct {
	mtx   sync.RWMutex
	beers map[int]*beer.Beer
}

func NewBeerRepository(beers map[int]*beer.Beer) beer.BeerRepository {
	if beers == nil {
		beers = make(map[int]*beer.Beer)
	}

	return &beerRepository{
		beers: beers,
	}
}

func (r *beerRepository) Fetch() ([]*beer.Beer, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]*beer.Beer, len(r.beers))
	for _, value := range r.beers {
		values = append(values, value)
	}
	return values, nil
}

func (r *beerRepository) FetchById(id int) (*beer.Beer, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	for _, value := range r.beers {
		if value.Id == id {
			return value, nil
		}
	}
	return nil, fmt.Errorf("The Id %s doesn't exist", strconv.Itoa(id))
}

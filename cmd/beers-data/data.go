package beers

import (
	beer "github.com/ericuss/golangSimpleKatas/pkg"
)

var Beers = map[int]*beer.Beer{
	0: &beer.Beer{
		ID:    "0",
		Name:  "Tricerahops",
		Price: 9.50,
	},
	1: &beer.Beer{
		ID:    "1",
		Name:  "All Together",
		Price: 4.90,
	},
	2: &beer.Beer{
		ID:    "2",
		Name:  "Imparable",
		Price: 2.85,
	},
	3: &beer.Beer{
		ID:    "3",
		Name:  "Mucho Caliente",
		Price: 5.95,
	},
}

package beer

type Beer struct {
	ID    string  `bson:"-"`
	Name  string  `bson:"Name"`
	Price float64 `bson:"Price"`
	// ID    string  `json:"id",bson:"-"`
	// Name  string  `json:"name,omitempty",bson:"Name"`
	// Price float32 `json:"price,omitempty",bson:"Price"`
}

type BeerRepository interface {
	Fetch() ([]*Beer, error)
	FetchById(Id int) (*Beer, error)
}

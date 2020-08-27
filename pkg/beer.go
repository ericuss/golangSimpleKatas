package beer

type Beer struct {
	Id    int     `json:"id"`
	Name  string  `json:"name,omitempty"`
	Price float32 `json:"price,omitempty"`
}

type BeerRepository interface {
	Fetch() ([]*Beer, error)
	FetchById(Id int) (*Beer, error)
}

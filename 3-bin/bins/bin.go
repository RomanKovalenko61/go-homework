package bins

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBin(id string, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   false,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

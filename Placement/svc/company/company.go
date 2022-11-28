package company

import (
	"Project3/Placement/store"
)

type Company struct {
	store store.Company
}

func New(store store.Company) Company {
	return Company{store: store}
}

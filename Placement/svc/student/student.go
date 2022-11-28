package student

import (
	"Project3/Placement/store"
)

type Service struct {
	store store.Student
}

func New(store store.Student) Service {
	return Service{store: store}
}

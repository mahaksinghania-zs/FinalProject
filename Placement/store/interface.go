package store

type Company interface {
	GetById()
	Create()
	Update()
	Delete()
}

type Student interface {
	Get()
	GetById()
	Create()
	Update()
	Delete()
}

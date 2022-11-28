package svc

import (
	"Project3/Placement/entities"
)

type Company interface {
	GetById(Id string) entities.Company
	Create(company entities.Company) string
	Update(company entities.Company) string
	Delete(Id string) string
}

type Student interface {
	Get(Id string) entities.Student
	GetById(Id string) entities.Student
	Create(student entities.Student)
	Update(student entities.Student)
	Delete(student entities.Student)
}

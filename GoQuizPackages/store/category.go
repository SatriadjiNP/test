package store

import "github.com/google/uuid"

type Category struct {
	cateId   uuid.UUID
	cateName string
}

// constructor
func NewCategory(cateName string) *Category {
	return &Category{
		cateId:   uuid.New(),
		cateName: cateName,
	}
}

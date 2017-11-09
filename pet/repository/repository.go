package repository

import models "petstore/pet"

type PetRepository interface {
	Fetch(cursor string, num int64) ([]*models.Pet, error)
	GetByID(id int64) (*models.Pet, error)
	Update(pet *models.Pet) (*models.Pet, error)
	Store(a *models.Pet) (int64, error)
	Delete(id int64) (bool, error)
}

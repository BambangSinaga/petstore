package usecase

import (
	"strconv"
	"time"

	"petstore/pet"
	"petstore/pet/repository"
)

type PetUsecase interface {
	Fetch(cursor string, num int64) ([]*pet.Pet, string, error)
	GetByID(id int64) (*pet.Pet, error)
	Update(ar *pet.Pet) (*pet.Pet, error)
	Store(*pet.Pet) (*pet.Pet, error)
	Delete(id int64) (bool, error)
}

type petUsecase struct {
	petRepos repository.PetRepository
}

func (a *petUsecase) Fetch(cursor string, num int64) ([]*pet.Pet, string, error) {
	if num == 0 {
		num = 10
	}

	listPet, err := a.petRepos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}
	nextCursor := ""

	if size := len(listPet); size == int(num) {
		lastId := listPet[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listPet, nextCursor, nil
}

func (a *petUsecase) GetByID(id int64) (*pet.Pet, error) {

	return a.petRepos.GetByID(id)
}

func (a *petUsecase) Update(ar *pet.Pet) (*pet.Pet, error) {

	oldPet, _ := a.GetByID(ar.ID)
	if oldPet == nil {
		return nil, pet.NOT_FOUND_ERROR
	}

	if ar.Name == "" {
		ar.Name = oldPet.Name
	}
	if ar.Age == 0 {
		ar.Age = oldPet.Age
	}
	if ar.Photo == "" {
		ar.Photo = oldPet.Photo
	}
	
	ar.UpdatedAt = time.Now()


	return a.petRepos.Update(ar)
}

func (a *petUsecase) Store(m *pet.Pet) (*pet.Pet, error) {

	/*existedPet, _ := a.GetByTitle(m.Title)
	if existedPet != nil {
		return nil, pet.CONFLIT_ERROR
	}*/

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	id, err := a.petRepos.Store(m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (a *petUsecase) Delete(id int64) (bool, error) {
	existedPet, _ := a.GetByID(id)

	if existedPet == nil {
		return false, pet.NOT_FOUND_ERROR
	}

	return a.petRepos.Delete(id)
}

func NewPetUsecase(a repository.PetRepository) PetUsecase {
	return &petUsecase{a}
}

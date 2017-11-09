package mysql

import (
	"database/sql"

	models "petstore/pet"
	"petstore/pet/repository"
)

type mysqlPetRepository struct {
	Conn *sql.DB
}

func (m *mysqlPetRepository) fetch(query string, args ...interface{}) ([]*models.Pet, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {

		return nil, models.INTERNAL_SERVER_ERROR
	}
	defer rows.Close()
	result := make([]*models.Pet, 0)
	for rows.Next() {
		t := new(models.Pet)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Age,
			&t.Photo,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {

			return nil, models.INTERNAL_SERVER_ERROR
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlPetRepository) Fetch(cursor string, num int64) ([]*models.Pet, error) {

	query := `SELECT id,name,age,photo,updated_at, created_at
  						FROM pet WHERE ID > ? LIMIT ?`

	return m.fetch(query, cursor, num)

}
func (m *mysqlPetRepository) GetByID(id int64) (*models.Pet, error) {
	query := `SELECT id,name,age,photo,updated_at, created_at
  						FROM pet WHERE ID = ?`

	list, err := m.fetch(query, id)
	if err != nil {
		return nil, err
	}

	a := &models.Pet{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}

	return a, nil
}

func (m *mysqlPetRepository) Store(a *models.Pet) (int64, error) {

	query := `INSERT  pet SET name=? , age=? , photo=? , updated_at=? , created_at=?`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {

		return 0, models.INTERNAL_SERVER_ERROR
	}
	res, err := stmt.Exec(a.Name, a.Age, a.Photo, a.CreatedAt, a.UpdatedAt)
	if err != nil {

		return 0, models.INTERNAL_SERVER_ERROR
	}
	return res.LastInsertId()
}

func (m *mysqlPetRepository) Delete(id int64) (bool, error) {
	query := "DELETE FROM pet WHERE id = ?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return false, models.INTERNAL_SERVER_ERROR
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return false, models.INTERNAL_SERVER_ERROR
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, models.INTERNAL_SERVER_ERROR
	}
	if rowsAfected <= 0 {
		return false, models.INTERNAL_SERVER_ERROR
	}

	return true, nil
}
func (m *mysqlPetRepository) Update(ar *models.Pet) (*models.Pet, error) {
	query := `UPDATE pet set name=?, age=?, photo=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil, nil
	}
	res, err := stmt.Exec(ar.Name, ar.Age, ar.Photo, ar.UpdatedAt, ar.ID)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect < 1 {
		return nil, models.INTERNAL_SERVER_ERROR
	}

	return ar, nil
}

func NewMysqlPetRepository(Conn *sql.DB) repository.PetRepository {

	return &mysqlPetRepository{Conn}
}

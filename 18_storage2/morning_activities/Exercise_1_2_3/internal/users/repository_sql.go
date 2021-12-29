package internal

import (
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/18_storage2/morning_activities/Exercise_1_2_3/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/18_storage2/morning_activities/Exercise_1_2_3/pkg/db"
)

type Repository_sql interface {
	Store(users models.Users) (models.Users, error)
	Get_one_user(id int) models.Users
	Get_by_name(name string) ([]models.Users, error)
	Get_all_users() ([]models.Users, error)
	Update(users models.Users) (models.Users, error)
}

type repository_sql struct{}

func New_repository_sql() Repository_sql {
	return &repository_sql{}
}

func (r *repository_sql) Store(users models.Users) (models.Users, error) {
	db := db.Storage_DB

	insert_user_store := `INSERT INTO users_sql(
		first_name, 
		last_name, 
		email, 
		age, 
		height, 
		active, 
		date) 
		VALUES ( ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := db.Prepare(insert_user_store)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(users.FirstName, users.LastName, users.Email, users.Age, users.Height, users.Active, users.Date)
	if err != nil {
		return models.Users{}, err
	}
	id_created, _ := result.LastInsertId()
	users.ID = int(id_created)
	return users, nil
}

func (r *repository_sql) Get_one_user(id int) models.Users {
	db := db.Storage_DB
	var user_readed models.Users

	select_one_user := `SELECT 
	id,
	first_name, 
	last_name, 
	email, 
	age
	FROM users_sql WHERE id = ?`
	rows, err := db.Query(select_one_user, id)

	if err != nil {
		log.Fatal(err)
		return user_readed
	}
	for rows.Next() {
		err := rows.Scan(
			&user_readed.ID,
			&user_readed.FirstName,
			&user_readed.LastName,
			&user_readed.Email,
			&user_readed.Age,
		)
		if err != nil {
			log.Fatal(err)
			return user_readed
		}
	}
	return user_readed
}

func (r *repository_sql) Get_by_name(name string) ([]models.Users, error) {
	db := db.Storage_DB
	var users_by_name []models.Users
	var user_readed models.Users

	select_by_name := `SELECT
	first_name,
	last_name,
	email,
	age
	FROM users_sql WHERE first_name = ?`
	rows, err := db.Query(select_by_name, name)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&user_readed.FirstName,
			&user_readed.LastName,
			&user_readed.Email,
			&user_readed.Age,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	users_by_name = append(users_by_name, user_readed)
	return users_by_name, nil
}

func (r *repository_sql) Get_all_users() ([]models.Users, error) {
	db := db.Storage_DB
	var all_users []models.Users
	var user_readed models.Users

	query_get_all_users := `SELECT
	first_name,
	last_name,
	email,
	age
	FROM users_sql`

	rows, err := db.Query(query_get_all_users)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(
			&user_readed.FirstName,
			&user_readed.LastName,
			&user_readed.Email,
			&user_readed.Age,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		all_users = append(all_users, user_readed)
	}
	return all_users, nil
}

func (r *repository_sql) Update(users models.Users) (models.Users, error) {
	db := db.Storage_DB

	update_user := `UPDATE users_sql SET first_name = ?, last_name = ?, email = ? WHERE id = ?`
	stmt, err := db.Prepare(update_user)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(users.FirstName, users.LastName, users.Email, users.ID)
	if err != nil {
		return models.Users{}, err
	}
	rows_updated, _ := result.RowsAffected()

	if rows_updated == 0 {
		return models.Users{}, errors.New("user not found")
	}
	return users, nil
}

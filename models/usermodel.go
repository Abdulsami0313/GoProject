package models

import (
	"database/sql"

	"github.com/project/config"
	"github.com/project/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {

	conn, err := config.DBConnection()

	if err != nil {

		panic(err)
	}

	return &UserModel{

		db: conn,
	}

}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("select * from students where "+fieldName+"= ? limit 1", fieldValue)
	if err != nil {

		return err
	}

	defer row.Close()

	for row.Next() {

		row.Scan(&user.Id, &user.Name, &user.Email, &user.Address, &user.Phone, &user.Password)
	}

	return nil
}

func UserNew() *UserModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &UserModel{db: db}
}

func (u UserModel) Create(user entities.User) (int64, error) {

	result, err := u.db.Exec("insert into students (name,email,address,phone,password) values (?,?,?,?,?)", user.Name, user.Email, user.Address, user.Phone, user.Password)
	if err != nil {

		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId, nil

}

func (m *UserModel) FindAll(user *[]entities.User) error {

	rows, err := m.db.Query("select * from students")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.User
		rows.Scan(&data.Id, &data.Name, &data.Email, &data.Address, &data.Phone, &data.Password)
		*user = append(*user, data)
	}
	return nil
}

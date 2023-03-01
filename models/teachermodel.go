package models

import (
	"database/sql"

	"github.com/project/config"
	"github.com/project/entities"
)

type TeacherModel struct {
	db *sql.DB
}

func TeacherNew() *TeacherModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &TeacherModel{db: db}
}

func (m *TeacherModel) FindAll(teacher *[]entities.Teacher) error {

	rows, err := m.db.Query("select * from teachers")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Teacher
		rows.Scan(&data.Id, &data.TeacherName, &data.Position, &data.Description, &data.Image)
		*teacher = append(*teacher, data)
	}
	return nil
}

func (m *TeacherModel) Create(teacher *entities.Teacher) error {

	result, err := m.db.Exec("insert into teachers (teachername,position,description,image)values(?,?,?,?)", teacher.TeacherName, teacher.Position, teacher.Description, teacher.Image)
	if err != nil {

		return err
	}

	lastInsertId, _ := result.LastInsertId()
	teacher.Id = lastInsertId
	return nil

}

func (m *TeacherModel) Find(id int64, teacher *entities.Teacher) error {

	return m.db.QueryRow("select * from teachers where id = ?", id).Scan(

		&teacher.Id,
		&teacher.TeacherName,
		&teacher.Position,
		&teacher.Description,
		&teacher.Image)
}

func (m *TeacherModel) Update(teacher entities.Teacher) error {

	if teacher.Image != "" {

		_, err := m.db.Exec("update teachers set teachername = ?, position = ?, description = ?, image = ? where id = ?", teacher.TeacherName, teacher.Position, teacher.Description, teacher.Image, teacher.Id)

		if err != nil {
			return err
		}

		return nil
	}

	if teacher.Image == "" {

		_, err := m.db.Exec("update teachers set teachername = ?, position = ?, description = ?, image = ? where id = ?", teacher.TeacherName, teacher.Position, teacher.Description, teacher.Image, teacher.Id)

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (m *TeacherModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from teachers where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

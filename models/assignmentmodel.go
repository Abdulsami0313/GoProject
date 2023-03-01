package models

import (
	"database/sql"
	"fmt"

	"github.com/project/config"
	"github.com/project/entities"
)

type AssignmentModel struct {
	db *sql.DB
}

func AssignmentNew() *AssignmentModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &AssignmentModel{db: db}
}

func (m *AssignmentModel) FindAll(assignment *[]entities.Assignment) error {

	rows, err := m.db.Query("select * from assignments")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Assignment
		rows.Scan(&data.Id, &data.Title, &data.Assignment, &data.DueDate, &data.TotalMarks, &data.Submit, &data.Result)
		*assignment = append(*assignment, data)
	}
	return nil
}

func (m *AssignmentModel) Create(assignment *entities.Assignment) error {

	result, err := m.db.Exec("insert into assignments (title,assignment,duedate,totalmarks,submit,result)values(?,?,?,?,?,?)", assignment.Title, assignment.Assignment, assignment.DueDate, assignment.TotalMarks, assignment.Submit, assignment.Result)
	if err != nil {

		return err
	}

	fmt.Println(result)

	lastInsertId, _ := result.LastInsertId()
	assignment.Id = lastInsertId
	return nil

}

func (m *AssignmentModel) Find(id int64, assignment *entities.Assignment) error {

	return m.db.QueryRow("select * from assignments where id = ?", id).Scan(

		&assignment.Id,
		&assignment.Title,
		&assignment.Assignment,
		&assignment.DueDate,
		&assignment.TotalMarks,
		&assignment.Submit,
		&assignment.Result)
}

func (m *AssignmentModel) Update(assignment entities.Assignment) error {

	_, err := m.db.Exec("update assignments set title = ?, assignment = ?, duedate = ?, totalmarks = ?, submit = ? , result = ? where id = ?", assignment.Title, assignment.Assignment, assignment.DueDate, assignment.TotalMarks, assignment.Submit, assignment.Result, assignment.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *AssignmentModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from assignments where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

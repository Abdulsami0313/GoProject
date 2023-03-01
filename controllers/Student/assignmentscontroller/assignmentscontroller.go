package assignments

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/project/entities"
	"github.com/project/models"
)

var assignmentModel = models.AssignmentNew()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{

		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/studentassignment.html", "views/studentdashboardlayout.html")
	temp.Execute(w, data)
}

func GetData() string {

	buffer := &bytes.Buffer{}
	temp, _ := template.New("studentassignmentdata.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {

			return a + b
		},
	}).ParseFiles("views/studentassignmentdata.html")

	var assignment []entities.Assignment
	err := assignmentModel.FindAll(&assignment)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"assignment": assignment,
	}

	temp.ExecuteTemplate(buffer, "studentassignmentdata.html", data)
	return buffer.String()
}

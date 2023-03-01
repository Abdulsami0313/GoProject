package manageassignmentscontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/project/entities"
	"github.com/project/models"
)

var assignmentModel = models.AssignmentNew()

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Assignments")
	data := map[string]interface{}{

		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/assignment.html", "views/dashboardlayout.html")
	temp.Execute(w, data)
}

func GetData() string {

	buffer := &bytes.Buffer{}
	temp, _ := template.New("assignmentdata.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {

			return a + b
		},
	}).ParseFiles("views/assignmentdata.html")

	var assignment []entities.Assignment
	err := assignmentModel.FindAll(&assignment)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"assignment": assignment,
	}

	temp.ExecuteTemplate(buffer, "assignmentdata.html", data)
	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}

	if err != nil {

		data = map[string]interface{}{
			"title": "Add Assignment Data",
		}
	} else {

		var assignment entities.Assignment
		err := assignmentModel.Find(id, &assignment)
		if err != nil {

			panic(err)
		}
		data = map[string]interface{}{
			"title":      "Edit Assignment Data",
			"assignment": assignment,
		}
	}

	temp, _ := template.ParseFiles("views/assignmentform.html")
	temp.Execute(w, data)

}

func Store(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		r.ParseForm()
		var assignment entities.Assignment
		fmt.Println(assignment)
		assignment.Title = r.Form.Get("title")
		assignment.Assignment = r.Form.Get("assignment")
		assignment.DueDate = r.Form.Get("duedate")
		assignment.TotalMarks = r.Form.Get("totalmarks")
		assignment.Submit = r.Form.Get("submit")
		assignment.Result = r.Form.Get("result")

		fmt.Println(assignment.Title)

		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		var data map[string]interface{}

		if err != nil {

			err := assignmentModel.Create(&assignment)
			if err != nil {

				RepsonseError(w, http.StatusInternalServerError, err.Error())
				return
			}

			data = map[string]interface{}{

				"message": "Data Successfully Changed",
				"data":    template.HTML(GetData()),
			}
		} else {

			fmt.Println("Update")

			assignment.Id = id
			err := assignmentModel.Update(assignment)
			if err != nil {

				RepsonseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{

				"message": "Data Successfully Updated",
				"data":    template.HTML(GetData()),
			}
		}

		ResponseJson(w, http.StatusOK, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

	if err != nil {

		panic(err)
	}
	err = assignmentModel.Delete(id)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"message": "Student Delete Successfully",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)
}

func RepsonseError(w http.ResponseWriter, code int, message string) {

	ResponseJson(w, code, map[string]string{"error": message})

}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

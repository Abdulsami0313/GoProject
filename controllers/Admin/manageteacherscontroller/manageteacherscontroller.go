package manageteacherscontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/project/entities"
	"github.com/project/models"
)

var teacherModel = models.TeacherNew()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{

		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/teacher.html", "views/dashboardlayout.html")
	temp.Execute(w, data)
}

func GetData() string {

	buffer := &bytes.Buffer{}
	temp, _ := template.New("adminteacher.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {

			return a + b
		},
	}).ParseFiles("views/adminteacher.html")

	var teacher []entities.Teacher

	err := teacherModel.FindAll(&teacher)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"teacher": teacher,
	}

	temp.ExecuteTemplate(buffer, "adminteacher.html", data)
	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}

	if err != nil {

		data = map[string]interface{}{
			"title": "Add Student Data",
		}
	} else {

		var teacher entities.Teacher
		err := teacherModel.Find(id, &teacher)
		if err != nil {

			panic(err)
		}
		data = map[string]interface{}{
			"title":   "Edit Student Data",
			"teacher": teacher,
		}
	}

	temp, _ := template.ParseFiles("views/teacherform.html")
	temp.Execute(w, data)

}

func Store(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		// r.ParseForm()
		// var teacher entities.Teacher
		// teacher.TeacherName = r.Form.Get("teachername")
		// teacher.Position = r.Form.Get("position")
		// teacher.Description = r.Form.Get("description")

		r.ParseMultipartForm(10 << 20)
		var teacher entities.Teacher
		teacher.TeacherName = r.Form.Get("teachername")
		teacher.Position = r.Form.Get("position")
		teacher.Description = r.Form.Get("description")

		id, errs := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		file, handler, err := r.FormFile("image")

		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			fmt.Println(r.FormFile("image"))
			return
		}

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Create file
		var imagename = time.Now().UnixNano()
		dst, err := os.Create(fmt.Sprintf("./static/uploads/teachers/%d%s", imagename, filepath.Ext(handler.Filename)))
		var dbimagename = fmt.Sprintf("%d%s", imagename, filepath.Ext(handler.Filename))
		teacher.Image = dbimagename
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data map[string]interface{}

		if errs != nil {

			errs := teacherModel.Create(&teacher)
			if err != nil {

				RepsonseError(w, http.StatusInternalServerError, errs.Error())
				return
			}

			data = map[string]interface{}{

				"message": "Data Successfully Changed",
				"data":    template.HTML(GetData()),
			}
		} else {

			teacher.Id = id
			err := teacherModel.Update(teacher)
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
	err = teacherModel.Delete(id)
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

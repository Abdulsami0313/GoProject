package managecoursecontroller

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

var courseModel = models.New()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{

		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/course.html", "views/dashboardlayout.html")
	temp.Execute(w, data)
}

func GetData() string {

	buffer := &bytes.Buffer{}
	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {

			return a + b
		},
	}).ParseFiles("views/data.html")

	var course []entities.Course
	err := courseModel.FindAll(&course)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"course": course,
	}

	temp.ExecuteTemplate(buffer, "data.html", data)
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

		var course entities.Course
		err := courseModel.Find(id, &course)
		if err != nil {

			panic(err)
		}
		data = map[string]interface{}{
			"title":  "Edit Student Data",
			"course": course,
		}
	}

	temp, _ := template.ParseFiles("views/form.html")
	temp.Execute(w, data)

}

func Store(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		// r.ParseForm()
		// var course entities.Course
		// course.CourseName = r.Form.Get("coursename")
		// course.Lesson = r.Form.Get("lesson")
		// course.Week = r.Form.Get("week")
		// course.Price = r.Form.Get("price")
		// course.Description = r.Form.Get("description")

		// id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		// Maximum upload of 10 MB files
		r.ParseMultipartForm(10 << 20)
		var course entities.Course
		course.CourseName = r.Form.Get("coursename")
		fmt.Println(course.Description)
		course.Lesson = r.Form.Get("lesson")
		course.Week = r.Form.Get("week")
		course.Price = r.Form.Get("price")
		course.Description = r.Form.Get("description")

		id, errs := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		// Get handler for filename, size and headers
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
		dst, err := os.Create(fmt.Sprintf("./static/uploads/course/%d%s", imagename, filepath.Ext(handler.Filename)))
		var dbimagename = fmt.Sprintf("%d%s", imagename, filepath.Ext(handler.Filename))
		course.Image = dbimagename
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

		fmt.Println(err)
		if errs != nil {

			errs := courseModel.Create(&course)
			if err != nil {

				RepsonseError(w, http.StatusInternalServerError, errs.Error())
				return
			}

			data = map[string]interface{}{

				"message": "Data Successfully Changed",
				"data":    template.HTML(GetData()),
			}
		} else {

			course.Id = id
			err := courseModel.Update(course)
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
	err = courseModel.Delete(id)
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

package studentcontroller

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"

	"github.com/project/config"
	"github.com/project/entities"
	"github.com/project/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Email    string
	Password string
}

var userModel = models.NewUserModel()
var courseModel = models.New()
var teacherModel = models.TeacherNew()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{

		"data":         template.HTML(GetData()),
		"teacher_data": template.HTML(TeacherGetData()),
	}

	temp, _ := template.ParseFiles("views/index.html")
	temp.Execute(w, data)

}

func TeacherGetData() string {

	buffer := &bytes.Buffer{}
	teacherTemplate, _ := template.New("teachers.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {

			return a + b
		},
	}).ParseFiles("views/teachers.html")

	var teacher []entities.Teacher
	teacher_err := teacherModel.FindAll(&teacher)
	if teacher_err != nil {

		panic(teacher_err)
	}

	teacher_data := map[string]interface{}{

		"teacher": teacher,
	}

	teacherTemplate.ExecuteTemplate(buffer, "teachers.html", teacher_data)
	return buffer.String()
}

func GetData() string {

	buffer := &bytes.Buffer{}
	temp, _ := template.New("courses.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {

			return a + b
		},
	}).ParseFiles("views/courses.html")

	var course []entities.Course
	err := courseModel.FindAll(&course)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"course": course,
	}

	temp.ExecuteTemplate(buffer, "courses.html", data)

	return buffer.String()
}

func StudentRegistration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/index.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		user := entities.User{

			Name:     r.Form.Get("name"),
			Email:    r.Form.Get("email"),
			Address:  r.Form.Get("address"),
			Phone:    r.Form.Get("phone"),
			Password: r.Form.Get("password"),
		}

		// errorMessages := make(map[string]interface{})

		// if user.Name == "" {

		// 	errorMessages["Name"] = "Name Required"
		// }

		// if user.Email == "" {

		// 	errorMessages["Email"] = "Email Required"
		// }

		// if user.Address == "" {

		// 	errorMessages["Address"] = "Address Required"
		// }

		// if user.Phone == "" {

		// 	errorMessages["Phone"] = "Contact number Required"
		// }

		// if user.Password == "" {

		// 	errorMessages["Password"] = "Password Required"
		// }

		// if len(errorMessages) > 0 {

		// 	data := map[string]interface{}{

		// 		"validation": errorMessages,
		// 	}

		// 	temp, _ := template.ParseFiles("views/index.html")
		// 	temp.Execute(w, data)

		// } else {

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		user.Password = string(hashPassword)

		_, err := userModel.Create(user)

		var message string
		if err != nil {

			message = "Register Not successfully" + message
		} else {

			message = "Register Success"
		}

		// data := map[string]interface{}{

		// 	"success": message,
		// }

		// temp, _ := template.ParseFiles("views/index.html")
		// temp.Execute(w, data)
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
}

func StudentLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/studentlogin.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		UserInput := &UserInput{

			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		var user entities.User
		userModel.Where(&user, "email", UserInput.Email)

		var message error
		if UserInput.Email == "" {
			message = errors.New("email and password required")
		} else {

			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))

			if errPassword != nil {

				message = errors.New("email and password not matched")
			}

		}

		if message != nil {

			data := map[string]interface{}{

				"error": message,
			}

			temp, _ := template.ParseFiles("views/studentlogin.html")
			temp.Execute(w, data)
		} else {

			session, _ := config.Store.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["email"] = user.Email
			session.Values["name"] = user.Name

			session.Save(r, w)

			http.Redirect(w, r, "/studentdashboard", http.StatusSeeOther)

		}

	}

}

func StudentDashboard(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {

		http.Redirect(w, r, "/student", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {

			http.Redirect(w, r, "/student", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{

				"name": session.Values["name"],
			}

			temp, _ := template.ParseFiles("views/studentdashboard.html", "views/studentdashboardlayout.html")
			temp.Execute(w, data)

		}

	}

}

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/student", http.StatusSeeOther)
}

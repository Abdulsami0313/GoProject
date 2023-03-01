package admicontroller

import (
	"errors"
	"net/http"
	"text/template"

	"github.com/project/config"
	"github.com/project/entities"
	"github.com/project/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
}

var adminModel = models.NewAdminModel()

func AdminLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/adminlogin.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		UserInput := &UserInput{

			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		var user entities.Admin
		adminModel.Where(&user, "username", UserInput.Username)

		var message error
		if user.Username == "" {

			message = errors.New("username and password required")
		} else {

			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))

			if errPassword != nil {

				message = errors.New("username and password not matched")
			}

		}

		if message != nil {

			data := map[string]interface{}{

				"error": message,
			}
			temp, _ := template.ParseFiles("views/adminlogin.html")
			temp.Execute(w, data)
		} else {

			session, _ := config.Store.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["email"] = user.Email
			session.Values["username"] = user.Name
			session.Values["name"] = user.Name

			session.Save(r, w)

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

		}

	}

}

func Dashboard(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {

			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{

				"name": session.Values["name"],
			}

			temp, _ := template.ParseFiles("views/dashboard.html", "views/dashboardlayout.html")
			temp.Execute(w, data)

		}

	}

}

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

package contactcontroller

import (
	"fmt"
	"net/http"

	"github.com/project/entities"
	"gopkg.in/gomail.v2"
)

func Contact(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	email := gomail.NewMessage()

	user := entities.Contact{

		FirstName: r.Form.Get("firstname"),
		LastName:  r.Form.Get("lastname"),
		Subject:   r.Form.Get("subject"),
		Email:     r.Form.Get("email"),
		Message:   r.Form.Get("message"),
	}

	email.SetHeader("From", user.Email)
	email.SetHeader("To", "Receiver-Email")
	email.SetHeader("Subject", user.Subject)
	email.SetBody("text/plain", user.Message)

	emailConfig := gomail.NewDialer("smtp.gmail.com", 587, "Receiver-Email", "Receiver-AppPassword")

	if err := emailConfig.DialAndSend(email); err != nil {

		fmt.Println(err)
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

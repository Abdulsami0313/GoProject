package main

import (
	"net/http"

	"github.com/project/controllers/Admin/admicontroller"
	"github.com/project/controllers/Admin/manageassignmentscontroller"
	"github.com/project/controllers/Admin/managecoursecontroller"
	"github.com/project/controllers/Admin/manageteacherscontroller"
	"github.com/project/controllers/Admin/manageuserscontroller"
	"github.com/project/controllers/Frontend/contactcontroller"
	"github.com/project/controllers/Frontend/studentcontroller"
	assignments "github.com/project/controllers/Student/assignmentscontroller"
)

// func sendMailSimple(){

// 	auth := smtp.PlainAuth(

// 		"",
// 		"hafizabdulsami97@gmail.com"
// 	)
// }

func main() {

	http.HandleFunc("/", studentcontroller.Index)
	http.HandleFunc("/admin", admicontroller.AdminLogin)
	http.HandleFunc("/student", studentcontroller.StudentLogin)
	http.HandleFunc("/contact-us", contactcontroller.Contact)
	http.HandleFunc("/register", studentcontroller.StudentRegistration)
	http.HandleFunc("/logout", admicontroller.Logout)
	http.HandleFunc("/studentlogout", studentcontroller.Logout)
	http.HandleFunc("/dashboard", admicontroller.Dashboard)
	http.HandleFunc("/studentdashboard", studentcontroller.StudentDashboard)
	http.HandleFunc("/manage-users", manageuserscontroller.Index)
	http.HandleFunc("/manage-course", managecoursecontroller.Index)
	http.HandleFunc("/manage-assignment", manageassignmentscontroller.Index)
	http.HandleFunc("/assignment/get_form", manageassignmentscontroller.GetForm)
	http.HandleFunc("/assignment/store", manageassignmentscontroller.Store)
	http.HandleFunc("/assignment/delete", manageassignmentscontroller.Delete)
	http.HandleFunc("/course/get_form", managecoursecontroller.GetForm)
	http.HandleFunc("/course/store", managecoursecontroller.Store)
	http.HandleFunc("/course/delete", managecoursecontroller.Delete)
	http.HandleFunc("/manage-teachers", manageteacherscontroller.Index)
	http.HandleFunc("/teacher/get_form", manageteacherscontroller.GetForm)
	http.HandleFunc("/teacher/store", manageteacherscontroller.Store)
	http.HandleFunc("/teacher/delete", manageteacherscontroller.Delete)
	http.HandleFunc("/assignments", assignments.Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8000", nil)
}

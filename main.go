package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type user struct {
	Email string
	Password string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index (res http.ResponseWriter, req *http.Request) {
	t := "This is some data that was passed in"
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(res, "index.gohtml", t)
}

func register(res http.ResponseWriter, req *http.Request) {
	//err := req.ParseForm()
	//if err != nil {
	//	panic(err)
	//}

	if req.Method == http.MethodPost {
		fmt.Println("This is a post request!")
		user := req.Form
		email := req.Form.Get("email")
		psw := req.Form.Get("psw")
		pswRepeat := req.Form.Get("psw-repeat")
		fmt.Println(user, email, psw, pswRepeat)
	}


	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(res, "register.gohtml", nil)
}

func plex(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(res, "plex.gohtml", nil)
}

func main() {

	router := http.NewServeMux()
	//router.Use(app.JwtAuthentication) //attach JWT auth middleware later
	//Routes
	router.Handle("/", http.HandlerFunc(index))
	router.Handle("/register", http.HandlerFunc(register))
	router.Handle("/plex", http.HandlerFunc(plex))
	router.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	router.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", router)

}

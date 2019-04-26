package main

import (
	"fmt"
	//"github.com/gorilla/mux" // still deciding on a mux
	//"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
)

func index (res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res,`
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}

func main() {


	router := http.NewServeMux()
	//router := mux.NewRouter()
	//router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)
	router.Handle("/", http.HandlerFunc(index))
	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}

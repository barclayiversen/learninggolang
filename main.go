//package gogame
//

package main

import (
	"github.com/barclayiversen/gogame/app"
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"

)

func main() {

	router := mux.NewRouter()

	router.HandleFunc ( path: "/api/user/new", controllers.CreateAccount).Methods( methods: "POST")
	router.HandleFunc ( path: "/api/user/login", controllers.Authenticate).Methods( methods: "POST")
	router.HandleFunc ( path: "/api/games/new", controllers.CreateGame).Methods( methods: "POST")
	//router.HandleFunc ( path: "/api/user/new", controllers.GetContactsFor).Methods( methods: "GET")

	router.Use(app.JwtAuthentication)

	port := os.Getenv( key: "PORT" )
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":" + port, router)
	if err != nil {

		fmt.Println(err)
	}
}
package models

import (
  "github.com/julienschmidt/httprouter"
  "gopkg.in/mgo.v2"
)

func main() {
  r := httprouter.New()
  http.ListenAndServe("localhost:8000", r)
}

func getSession() *mgo.Session {
  s, err := mgo.Dial("mongodb://localhost")
  if err != nil {
    fmt.Println("there was an err", err)
    panic(err)
  }
  return s
}

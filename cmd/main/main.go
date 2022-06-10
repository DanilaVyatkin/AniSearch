package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("hello world"))
}

func main() {
	log.Println("create router")
	router := httprouter.New()

	//cfg := config.GetConfig()
	//
	//postgresqlClient, err := postgesql.NewClient(context.TODO(), cfg.Storage, 3)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//newDB := db.NewDB(postgresqlClient)

	//newCartoon := cartoon.Cartoon{
	//	ID:          "cb23c467-9d45-4ec6-87cc-45c2e0fcb05c",
	//	Name:        "totoro",
	//	Genre:       "qwe",
	//	Rating:      "10",
	//	Description: "totoro",
	//}

	//err = newDB.Create(context.TODO(), &newCartoon)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = newDB.Update(context.TODO(), &newCartoon)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = newDB.Delete(context.TODO(), "57f86bd2-e7d8-4e44-9aed-9f0cb5c92e4f")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//all, err := newDB.FindAll(context.TODO())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, name := range all {
	//	log.Println(name)
	//}

	router.GET("/hello", Hello)
	log.Fatalln(http.ListenAndServe(":8080", router))
}

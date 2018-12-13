package main

import (
	"log"
	"net/http"

	ctl "./controllers"
	dao "./dao"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dao.Connect()

	fs := http.FileServer(http.Dir("view"))
	http.Handle("/", fs)
	http.HandleFunc("/bar", ctl.TestController)

	log.Println("Listening...")

	log.Fatal(http.ListenAndServe(":3001", nil))

}

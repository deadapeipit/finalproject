package main

import (
	"finalproject/database"
	"finalproject/handler"
	"finalproject/server"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	server.GetConfig("config/config.json")

	connString := server.Config.ConnectionString
	sql := database.NewSqlConnection(server.DecryptConnectionString(connString))
	database.SqlDatabase = sql
	defer sql.CloseConnection()

	r := mux.NewRouter()
	handler.InstallUsersHandler(r)
	handler.InstallPhotosHandler(r)
	handler.InstallCommentHandler(r)
	handler.InstallSocialMediaHandler(r)
	r.Use(server.SecureMiddleware)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Link: http://%s \n", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}

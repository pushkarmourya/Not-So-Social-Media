package main

import (
	"Not-So--Social-Media/database"
	postsAction "Not-So--Social-Media/store/posts"
	userAction "Not-So--Social-Media/store/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Not So Social Media...")

	database.StartDatabase()
	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	userAction.HandleUserRoute(router)
	postsAction.HandlePostsRoute(router)
}

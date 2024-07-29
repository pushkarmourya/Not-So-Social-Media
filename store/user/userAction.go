package userAction

import (
	"Not-So--Social-Media/database"
	"Not-So--Social-Media/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleUserRoute(router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("", getAllUser).Methods("GET")
	userRouter.HandleFunc("/create", createUser).Methods("POST")
	userRouter.HandleFunc("/delete/{id}", deleteUser).Methods("GET")
	userRouter.HandleFunc("/update/{id}", updateUser).Methods("POST")
}

func createUser(res http.ResponseWriter, req *http.Request) {
	var newUser model.User

	json.NewDecoder(req.Body).Decode(&newUser)
	database.CreateUser(newUser)
}

func getAllUser(res http.ResponseWriter, req *http.Request) {

}

func deleteUser(res http.ResponseWriter, req *http.Request) {

}

func updateUser(res http.ResponseWriter, req *http.Request) {

}

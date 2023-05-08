package handler

import (
	"fmt"
	"go-jwt-rest-mongodb/repository"
	"log"
	"net/http"
)

type Handler struct {
	Repo *repository.Repository
}

func (h *Handler) HandleRequest() {
	http.Handle("/", isAuthorized(homePage))
	http.HandleFunc("/login", login)

	http.Handle("/users", isAuthorized(UserIndex))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserIndex(w http.ResponseWriter, r *http.Request) {

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secret Information")
}

func login(w http.ResponseWriter, r *http.Request) {
	token, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	w.Write([]byte(token))

}

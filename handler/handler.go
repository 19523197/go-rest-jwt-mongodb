package handler

import (
	"encoding/json"
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

	http.Handle("/users", UserHandler(h.Repo))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserIndex(w http.ResponseWriter, r *http.Request) {

}

func UserHandler(repo *repository.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			user, err := repo.UserRepo.GetUser()
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res, err := json.Marshal(user)
			w.WriteHeader(200)
			w.Write(res)

			return
		}
	})

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

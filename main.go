package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go-jwt-rest-mongodb/database"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var mySigningKey = []byte("my-secret-key")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secret Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	token, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	w.Write([]byte(token))

}

func handleRequest() {
	http.Handle("/", isAuthorized(homePage))
	http.HandleFunc("/login", login)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "rama"
	claims["exp"] = time.Now().Add(time.Hour * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Failed to load environment: %s", err.Error())
	}

	client := database.ConnectMongoClient()

	fmt.Println("Starting server on http://localhost:8080")
	handleRequest()

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

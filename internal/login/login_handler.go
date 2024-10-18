package login

import (
	"encoding/json"
	"fmt"
	"net/http"
    "io"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func printBody(r *http.Request){
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf(`{"message": "error reading request body"}`)
		return
	}
	fmt.Printf("The request body is %s \n", body)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u User

	decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&u)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, `{"message": "error decoding request body"}`)
		fmt.Fprint(w, err.Error())
		fmt.Printf("error decoding request body %s\n", err)
        return
    }
    
    fmt.Printf("The user is %v\n", u)

	if u.Username == "admin" && u.Password == "admin" {
		tokenString, err := CreateToken(u.Username)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, `{"message": "error creating token"}`)
            return 
        }
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, `{"token": "%s"}`, tokenString)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "success"}`)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"message": "invalid credentials"}`)
        return
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	name, err := VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}
	fmt.Fprintf(w, "Welcome to the the protected area %s", name)
}

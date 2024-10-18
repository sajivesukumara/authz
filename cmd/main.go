package main

import (
	"fmt"
	"net/http"

	"authz/internal/controllers"
	"authz/internal/initializers"
	"authz/internal/login"
	"authz/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func mux_http_router() {
    router := mux.NewRouter()
    router.HandleFunc("/login", login.LoginHandler).Methods("POST")
    router.HandleFunc("/protected", login.ProtectedHandler).Methods("GET")

    fmt.Println("Server is running on port 8000")
    err := http.ListenAndServe(":8000", router)    
    if err != nil {
        fmt.Println("Could not start the server", err)
    } 
    fmt.Println("Server is running on port 8000")
}


func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func gin_http_router() {
    router := gin.Default()

    router.POST("/auth/signup", controllers.CreateUser)
    router.POST("/auth/login", controllers.Login)    
    router.GET("/user/profile", middleware.CheckAuth, controllers.GetUserProfile)
    router.Run(":8000")

}


func main() {
    gin_http_router()
}
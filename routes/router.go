package routes

import (
	"fmt"
	"os"

	"net/http"

	logger "github.com/arierimbaboemi/bank-lib-service/config"
	hand "github.com/arierimbaboemi/go-banking-auth/adapter/handler"
	repo "github.com/arierimbaboemi/go-banking-auth/adapter/repository"
	conf "github.com/arierimbaboemi/go-banking-auth/config"
	"github.com/arierimbaboemi/go-banking-auth/middleware"
	serv "github.com/arierimbaboemi/go-banking-auth/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewRouter(router *mux.Router, db *sqlx.DB) {
	// apply middleware to all routes
	router.Use(middleware.ApiKeyMiddleware)
	accountRepo := repo.NewAccountRepositoryDB(db)
	authService := serv.NewAuthService(accountRepo)
	authHandler := hand.NewAuthHandlerDB(authService)

	router.Handle("/login", http.HandlerFunc(authHandler.Login)).Methods("POST")
	router.Handle("/protected", middleware.AuthMiddleware(authService, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Welcome to the protected area!")
	}))).Methods("GET")
}

func StartServer() {
	logger.InitiateLog()
	defer logger.CloseLog()

	db, _ := conf.NewDBConnectionENV()
	port := os.Getenv("SERVER_PORT")
	defer db.Close()

	router := mux.NewRouter()
	NewRouter(router, db)
	fmt.Println("starting server on port " + port)
	http.ListenAndServe(":"+port, router)
}

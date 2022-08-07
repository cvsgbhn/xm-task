package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"xm-task/packages/config"
	"xm-task/packages/domain"
	"xm-task/packages/handlers"
	"xm-task/packages/postgres"
	"xm-task/packages/storage"
)

func main() {
	l := log.New(os.Stdout, "xm-test-task ", log.LstdFlags)

	cf := config.GetConfig()

	dbConn, err := postgres.MakeDBconn(cf)
	if err != nil {
		l.Printf("Error connecting to db: ", err)
		return
	}

	st := storage.NewRepository(dbConn)

	srv := domain.NewCompanyService(st)

	ch := handlers.NewCompanies(l, srv)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ch.ShowCompanies)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ch.UpdateCompany)
	putRouter.Use(ch.MiddlewareCompanyValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/add", ch.AddCompany)
	postRouter.Use(ch.MiddlewareCompanyValidation)

	port := os.Getenv("PORT")

	// create a new server
	s := http.Server{
		Addr:         ":" + port,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port " + port)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

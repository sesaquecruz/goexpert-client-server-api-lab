package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sesaquecruz/goexpert-client-server-api-lab/internal/database/repository"
	"github.com/sesaquecruz/goexpert-client-server-api-lab/internal/service"
)

const (
	ExternalApiTimeout = 200 * time.Millisecond
	DBTimeout          = 10 * time.Millisecond
)

func main() {
	db, err := sql.Open("sqlite3", "server.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	qr := repository.NewQuoteRepository(db)
	qs := service.NewQuoteService()
	ss := service.NewServerService(qr, qs, ExternalApiTimeout, DBTimeout)

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", ss.UsdBrlHandler)

	log.Println("server is running on port 8080...")
	http.ListenAndServe(":8080", mux)
}

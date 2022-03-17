package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	// DB_URL="postgres://username:password@localhost:5432/mydb" ./main
	dbUrl := os.Getenv("DB_URL")
	dbPool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	textHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	jsonHandler := func(w http.ResponseWriter, req *http.Request) {
		tsObj := map[string]interface{}{
			"message": "Hello world",
			"ts":      time.Now().String(),
		}
		tsBytes, err := json.Marshal(tsObj)
		if err != nil {
			fmt.Println("json error:", err)
		}
		io.WriteString(w, string(tsBytes))
	}

	sqlHandler1 := func(w http.ResponseWriter, req *http.Request) {
		var ts string
		row := dbPool.QueryRow(context.Background(), "SELECT CAST(now() AS VARCHAR)")
		err := row.Scan(&ts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "query rows scan failed: %v\n", err)
		}
		tsObj := map[string]interface{}{
			"message": "Hello world",
			"ts":      ts,
		}
		tsJson, _ := json.Marshal(tsObj)
		io.WriteString(w, string(tsJson))
	}

	http.HandleFunc("/text", textHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/sql1", sqlHandler1)
	http.HandleFunc("/", textHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

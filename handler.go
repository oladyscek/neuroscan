package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"net/http"

	_ "modernc.org/sqlite"
)

var db *sql.DB

type UserRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Clinic   string `json:"clinic"`
	JobTitle string `json:"jobTitle"`
	Email    string `json:"email"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Phone == "" || req.Clinic == "" || req.JobTitle == "" || req.Email == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(`--sql
	insert into clients (name_enc,phone_enc,email_enc,job_enc,clinic_enc) values (?,?,?,?,?)`,
		req.Name, req.Phone, req.Email, req.JobTitle, req.Clinic)

	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

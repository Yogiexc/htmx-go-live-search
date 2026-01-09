package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

// User struct untuk data dummy
type User struct {
	ID    int
	Name  string
	Email string
	City  string
}

// Data dummy users di memory
var users = []User{
	{1, "Budi Santoso", "budi@email.com", "Jakarta"},
	{2, "Siti Nurhaliza", "siti@email.com", "Bandung"},
	{3, "Andi Wijaya", "andi@email.com", "Surabaya"},
	{4, "Rina Kusuma", "rina@email.com", "Jakarta"},
	{5, "Joko Widodo", "joko@email.com", "Solo"},
	{6, "Ani Susanti", "ani@email.com", "Yogyakarta"},
	{7, "Bambang Hermanto", "bambang@email.com", "Semarang"},
	{8, "Dewi Lestari", "dewi@email.com", "Bali"},
	{9, "Eko Prasetyo", "eko@email.com", "Malang"},
	{10, "Fitri Handayani", "fitri@email.com", "Medan"},
}

var templates *template.Template

func init() {
	var err error
	// Parse semua template dengan explicit files
	templates, err = template.ParseFiles(
		"templates/index.html",
		"templates/results.html",
	)
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}
}

func main() {
	// Static files (CSS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/search", handleSearch)

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler untuk halaman utama
func handleIndex(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler untuk search - return HTML fragment
func handleSearch(w http.ResponseWriter, r *http.Request) {
	// Simulasi delay network biar keliatan loading indicator
	time.Sleep(300 * time.Millisecond)

	// Ambil query parameter
	query := r.URL.Query().Get("q")
	query = strings.TrimSpace(strings.ToLower(query))

	// Filter users based on query
	var results []User
	if query != "" {
		for _, user := range users {
			// Search di name, email, atau city
			if strings.Contains(strings.ToLower(user.Name), query) ||
				strings.Contains(strings.ToLower(user.Email), query) ||
				strings.Contains(strings.ToLower(user.City), query) {
				results = append(results, user)
			}
		}
	}

	// Data untuk template
	data := struct {
		Query   string
		Results []User
		Count   int
	}{
		Query:   query,
		Results: results,
		Count:   len(results),
	}

	// Render HTML fragment (bukan full page)
	err := templates.ExecuteTemplate(w, "results.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
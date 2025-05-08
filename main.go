package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// rootHandler serves the root path with an HTML response.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "<h1>Welcome to the Home Page</h1>")
}

// homePage serves the home.html file.
func homePage(w http.ResponseWriter, r *http.Request) {
	const filePath = "himmu/home.html"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

// coursePage serves the courses.html file.
func coursePage(w http.ResponseWriter, r *http.Request) {
	const filePath = "himmu/courses.html"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

// aboutPage serves the about.html file.
func aboutPage(w http.ResponseWriter, r *http.Request) {
	const filePath = "himmu/about.html"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

// contactPage serves the contact.html file.
func contactPage(w http.ResponseWriter, r *http.Request) {
	const filePath = "himmu/contact.html"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}
}

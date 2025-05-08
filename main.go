package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, "<h1>Welcome to the Home Page</h1>")
}

func homePage(w http.ResponseWriter, r *http.Request) {
        filePath := "himmu/home.html"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
                http.Error(w, "File not found", http.StatusNotFound)
                return
        }
        http.ServeFile(w, r, filePath)
}

func coursePage(w http.ResponseWriter, r *http.Request) {
        filePath := "himmu/courses.html"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
                http.Error(w, "File not found", http.StatusNotFound)
                return
        }
        http.ServeFile(w, r, filePath)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
        filePath := "himmu/about.html"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
                http.Error(w, "File not found", http.StatusNotFound)
                return
        }
        http.ServeFile(w, r, filePath)
}

func contactPage(w http.ResponseWriter, r *http.Request) {
        filePath := "himmu/contact.html"
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

        err := http.ListenAndServe("0.0.0.0:8080", nil)
        if err != nil {
                log.Fatal(err)
        }
}

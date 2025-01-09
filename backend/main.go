package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
    var err error
    // Connexion à la base de données SQLite
    db, err = sql.Open("sqlite3", "./db/database.db")  // Assurez-vous que le chemin est correct
    if err != nil {
        log.Fatal(err)
    }
}

// Handler pour la page d'accueil
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Go Backend!")
}

// Handler pour afficher les utilisateurs
func usersHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, username FROM users")
    if err != nil {
        http.Error(w, "Failed to request user", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []string
    for rows.Next() {
        var id int
        var username string
        if err := rows.Scan(&id, &username); err != nil {
            http.Error(w, "Failed to parse users", http.StatusInternalServerError)
            return
        }
        users = append(users, fmt.Sprintf("%d: %s", id, username))
    }

    // Afficher la liste des utilisateurs
    for _, user := range users {
        fmt.Fprintln(w, user)
    }
}

func main() {
    // Définir les routes
    http.HandleFunc("/", homePage)     // Page d'accueil
    http.HandleFunc("/users", usersHandler)  // Liste des utilisateurs

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


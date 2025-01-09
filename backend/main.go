package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "github.com/go-redis/redis/v8"
    _ "github.com/mattn/go-sqlite3"
)

var (
    db      *sql.DB
    rdb     *redis.Client
    ctx     = context.Background()
)

func init() {
    var err error
    // Connexion à la base de données SQLite
    db, err = sql.Open("sqlite3", "./db/database.db")
    if err != nil {
        log.Fatal(err)
    }

    // Connexion à Redis
    rdb = redis.NewClient(&redis.Options{
        Addr: "redis:6379", // Nom du service Redis dans le docker-compose
    })
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    // Vérifier dans Redis avant de consulter la base de données
    cacheKey := "users_list"
    cachedData, err := rdb.Get(ctx, cacheKey).Result()
    if err == redis.Nil {
        // Si les données ne sont pas en cache, on les charge depuis SQLite
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

        // Mettre les données dans le cache Redis
        rdb.Set(ctx, cacheKey, users, 0) // 0 signifie pas d'expiration

        // Afficher les utilisateurs
        for _, user := range users {
            fmt.Fprintln(w, user)
        }
    } else if err != nil {
        http.Error(w, "Failed to retrieve data from cache", http.StatusInternalServerError)
        return
    } else {
        // Si les données sont en cache, on les renvoie directement
        fmt.Fprintln(w, cachedData)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Go Backend with Redis Cache!")
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/users", usersHandler)
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


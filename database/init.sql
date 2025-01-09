DROP TABLE IF EXISTS users;  -- Pour être sûr que la table existe bien avant de la recréer

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL
);

-- Ajoutez quelques utilisateurs pour tester
INSERT INTO users (username, password) VALUES ('testuser', 'password123');
INSERT INTO users (username, password) VALUES ('john_doe', 'mypassword');


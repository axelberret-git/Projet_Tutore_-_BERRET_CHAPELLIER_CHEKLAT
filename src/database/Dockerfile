# Utiliser une image de base légère, Alpine
FROM alpine:latest

# Installer SQLite
RUN apk add --no-cache sqlite sqlite-dev

# Créer le répertoire où la base de données sera stockée
RUN mkdir -p /data

# Copiez le fichier init.sql dans le conteneur pour initialiser la base de données
COPY init.sql /docker-entrypoint-initdb.d/

# Définir le répertoire de travail pour l'application
WORKDIR /data

# Commande par défaut pour démarrer le conteneur
CMD ["sqlite3", "/data/database.db"]


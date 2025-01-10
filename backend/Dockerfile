# Étape 1 : Construction du binaire Go
FROM golang:1.20.5 AS builder

WORKDIR /app

# Copier les fichiers de l'application
COPY . .

# Télécharger les dépendances et construire le binaire
RUN go mod tidy
RUN go build -o main .

# Étape 2 : Image finale pour exécuter l'application
FROM debian:bookworm-slim

WORKDIR /app

# Copier le binaire construit depuis l'étape précédente
COPY --from=builder /app/main .

# Installer redis-tools pour interagir avec Redis
RUN apt-get update \
    && apt-get install -y --no-install-recommends redis-tools \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Exposer le port utilisé par le backend
EXPOSE 8080

# Commande par défaut pour lancer l'application
CMD ["./main"]


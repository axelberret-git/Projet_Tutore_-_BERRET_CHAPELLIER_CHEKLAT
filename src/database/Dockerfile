# Étape 1 : Construction sur Alpine
FROM alpine:latest AS builder

RUN apk add --no-cache sqlite sqlite-dev

# Étape 2 : Image minimale pour exécuter SQLite
FROM alpine:latest

WORKDIR /data

# Copier uniquement SQLite depuis l’étape précédente
COPY --from=builder /usr/bin/sqlite3 /usr/bin/sqlite3
COPY init.sql /data/init.sql

CMD ["sqlite3", "/data/database.db"]


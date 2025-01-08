# Étape 1 : Build de l'application React
FROM node:18-slim AS build

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers du frontend
COPY . .

# Installer les dépendances et construire l'application
RUN npm install && npm run build

# Étape 2 : Serveur Nginx pour les fichiers statiques
FROM nginx:1.22.1

# Copier le build dans le répertoire de Nginx
COPY --from=build /app/build /usr/share/nginx/html

# Copier une configuration Nginx personnalisée si nécessaire
# Vous pouvez inclure un fichier default.conf si besoin
# COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

# Exposer le port 3000
EXPOSE 3000

# Commande par défaut pour lancer Nginx
CMD ["nginx", "-g", "daemon off;"]


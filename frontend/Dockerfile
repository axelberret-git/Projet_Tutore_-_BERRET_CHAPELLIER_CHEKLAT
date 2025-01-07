# Utilisation d'une image de base Node.js 18
FROM node:18-slim

# Mettre à jour les packages et installer Nginx
RUN apt-get update && apt-get install -y nginx

# Créer le répertoire pour l'application et copier les fichiers
WORKDIR /app
COPY . /app

# Installer les dépendances (pour un frontend React par exemple)
RUN npm install

# Copier le fichier de configuration nginx dans le conteneur
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

# Exposer les ports nécessaires
EXPOSE 3000 80

# Définir le script d'entrée comme point d'entrée du conteneur
COPY ./docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]


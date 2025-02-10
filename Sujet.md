# Codespace auto hébergée pour le développement web

## Présentation

Une codespace est un environnement de développement en ligne, sur un serveur, et qui permet de développer des applications dans un environnement isolé, personnalisé en fonction du projet, et sans rien installer sur la machine des développeurs.

L'objectif de ce projet est de mettre en place une solution auto-hébergée de codespaces, compatible avec toutes les technologies, grâce à l'utilisation de Docker.

## Fonctionnalités attendues

Soit un repository Git contenant un fichier `docker-compose.yml` qui décrit l'environnement de développement.
Avec par exemple ces containers : nextjs, supabase, redis, maildev...

On devrait pouvoir lancer un environnement de développement dans le navigateur ou directement dans vscode remote development.

Quelque soit sa forme, l'environnement de développement devra permettre:
- de lancer l'application et la voir dans le navigateur
- de voir les logs
- de debugger l'application (front avec les devtools, back avec des breakpoints)
- de voir / se connecter à la base de données
- de lancer les tests
- d'installer des dépendances
- de lancer des commandes directement sur l'environnement de développement
- d'utiliser des outils de linting, formatting, etc.

## Précisions

La solution devra être auto hébergée et fonctionner sur un serveur dédié, un VPS par exemple.

Quelques technologies envisageables : [Gitpod](https://www.gitpod.io/), [Coder](https://coder.com/)

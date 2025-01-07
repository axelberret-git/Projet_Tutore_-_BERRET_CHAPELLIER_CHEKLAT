# Projet_Tutore_-_BERRET_CHAPELLIER_CHEKLAT

## Technos utilisées

Choix de l'environnement de développement éphémère : 
Gitpod

docker-compose :
Go
Redis
Vue.js
MariaDB
Certification HTTPS
DNS

Il est important de savoir que si l'on souhaite utiliser notre Gitpod en auto-hébergé, d'utiliser un cluster Kubernetes.
Or comme nous décrivons une infrastructure de développement au travers d'un docker-compose, alors deux options s'offrent
à nous :

1. Faire du Docker in Docker (DinD)

Pratique dans le cas où l'on souhaite avoir une solution rapide et temporaire mais n'est pas préférable en terme de
performance, d'évolutivité, d'intégration avec Gitpod et d'interopérabilité avec Kubernetes.

2. Faire des manifests Kubernetes

Est plus complexe à mettre en place, mais il est en revanche préféré pour tous les autres aspects cités dans le
premier cas ainsi que pour sa stabilité.


## Sources

cluster Kubernetes : https://bobcares.com/blog/kubernetes-cluster-deployment-on-proxmox-8/

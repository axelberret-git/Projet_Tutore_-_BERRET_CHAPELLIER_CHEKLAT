# Projet_Tutore_-_BERRET_CHAPELLIER_CHEKLAT

## Utilisation

...


## Configuration

### Technos utilisées

Choix de l'environnement de développement éphémère :
Gitpod

docker-compose :
- Go
- React
- Redis
- Mailhog
- SQLite

Il est important de savoir que si l'on souhaite utiliser notre Gitpod en auto-hébergé, il est recommandé d'utiliser un 
cluster Kubernetes. Or comme nous décrivons une infrastructure de développement au travers d'un docker-compose, alors 
deux options s'offrent à nous mais nous avions décidé de faire des manifests Kubernetes.

C'est plus complexe à mettre en place, en revanche c'est préféré pour d'autres aspects tels que :
- des bonnes performances et une bonne optimisation
- une bonne évolutivité
- une intégration à Gitpod 100% native
- une compatibilité avec Helm, ArgoCD, etc ... (interopérabilité avec Kubernetes)


### Performance cluster Kubernetes, Gitpod et infrastructure Docker-Compose

Cette partie traite uniquement les performances recommandées pour la mise en place de Gitpod, de l'infrastructure Docker-Compose
ainsi que du cluster Kubernetes : (CPU, RAM, stockage)

Gitpod et infrastructure Docker-compose : (VM unique)
- Gitpod :          4 CPU, 10 Go RAM, 70 Go 
- Redis :           2 CPU, 1 Go RAM, 4 Go
- MariaDB :         2 CPU, 4 Go RAM, 30 Go
- Go :              2 CPU, 2 Go RAM, 2 Go
- React :           1 CPU, 2 Go RAM, 2 Go
- Mailhog :         1 CPU, 1 Go RAM, 2 Go
12 CPU, 20 Go RAM, 110 Go

Cluster Kubernetes : (3 VM : 1 master et 2 workers)
2 CPU, 4 Go RAM, 20 Go


## Sources

cluster Kubernetes : https://bobcares.com/blog/kubernetes-cluster-deployment-on-proxmox-8/
Video cluster Kubernetes : https://www.youtube.com/watch?v=PtQ8FOepn94

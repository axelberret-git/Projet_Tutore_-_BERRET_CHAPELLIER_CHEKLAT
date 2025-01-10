# Création du cluster Kubernetes

## Prérequis

Tout d'abord créer les VM master et worker 1 (worker 2 est optionel) dans proxmox et vérifier qu'ils ont une connection internet


## Master Kubernetes 

1. update le serveur

`apt update && apt upgrade -y`

2. Installation du k3s dans le master 

`sudo curl -sfL https://get.k3s.io | sh -`

3. Récupérer le token pour les workers 

`cat /var/lib/rancher/k3s/server/node-token`

Garder le token il est imporant 


## Workers Kubernetes 

1. Installer l'agent sur les VM avec l'ip du master et son token 

`sudo curl -sfL https://get.k3s.io | K3S_URL=https://<MASTER_IP>:6443 K3S_TOKEN=<TOKEN> sh -`

2. Vérfier depuis le master si l'agent et bien ajouter 

`sudo kubectl get nodes`

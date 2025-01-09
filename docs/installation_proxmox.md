# Installation Proxmox

## Mise en place de Proxmox

1. connexion ssh

`ssh -t root@100.64.85.12`


2. renomer l'hostname `/etc/hostname`

`proxmox.local`

vérifiable par `hostname -f`


3. ajouter le nouvel hote `/etc/hosts`

```
127.0.0.1	    localhost.localdomain localhost proxmox.local
100.64.85.12	proxmox.local proxmox dacs-503-1103
```


4. regénérer des certificats auto-signés pour mettre en place la première machine :

`openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/pve/local/pve-ssl.key -out /etc/pve/local/pve-ssl.pem -subj "/CN=proxmox.local"`


5. redémarrer les services de Proxmox

```
systemctl restart pveproxy
systemctl restart pvedaemon
```


6. aller sur la web interface

https://100.64.85.12:8006


## Mise en place du NAT


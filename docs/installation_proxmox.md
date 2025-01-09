# Installation Proxmox

## Mise en place de Proxmox

1. Se connecter avec SSH

`ssh -t root@100.64.85.12`


2. Renomer l'hostname `/etc/hostname`

`proxmox.local`

vérifiable par `hostname -f`


3. Ajouter le nouvel hote `/etc/hosts`

```
127.0.0.1	    localhost.localdomain localhost proxmox.local
100.64.85.12	proxmox.local proxmox dacs-503-1103
```


4. Regénérer des certificats auto-signés pour mettre en place la première machine :

`openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/pve/local/pve-ssl.key -out /etc/pve/local/pve-ssl.pem -subj "/CN=proxmox.local"`


5. Redémarrer les services de Proxmox

```
systemctl restart pveproxy
systemctl restart pvedaemon
```


6. Aller sur la web interface

https://100.64.85.12:8006


## Mise en place du NAT

1. Activer le NAT depuis l'hôte Proxmox

`iptables -t nat -A POSTROUTING -s 192.168.100.0/24 -o vmbr0 -j MASQUERADE`


2. Vérifier les règles en places

`iptables -t nat -L -n -v`


3. Configurer le bridge

`nano /etc/network/interfaces`

avec le contenu suivant

```
auto vmbr1
iface vmbr1 inet static
    address 192.168.100.1/24
    bridge-ports none
    bridge-stp off
    bridge-fd 0
```


4. Redémarrer le service de réseau de l'hôte

`systemctl restart networking`
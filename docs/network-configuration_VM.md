1. Création du script

`nano /usr/local/bin/configure-network.sh`

```
#!/bin/bash

CONFIG_FILE="/etc/netplan/50-cloud-init.yaml"

# Modifier le fichier de configuration Netplan
cat <<EOL > $CONFIG_FILE
network:
  version: 2
  renderer: networkd
  ethernets:
    ens18:
      dhcp4: no
      addresses:
        - 192.168.100.x/24
      routes:
        - to: default
          via: 192.168.100.1
      nameservers:
        addresses:
          - 8.8.8.8
          - 1.1.1.1
EOL

# Attendre 5 secondes avant d'appliquer Netplan
sleep 5

# Appliquer la configuration
/usr/sbin/netplan apply
```


2. Ajout des permissions

`chmod +x /usr/local/bin/configure-network.sh`


3. Ajout du service netplan-apply

`nano /etc/systemd/system/netplan-apply.service`

```
[Unit]
Description=Appliquer la configuration Netplan au démarrage
After=network-online.target
Wants=network-online.target

[Service]
Type=oneshot
ExecStart=/usr/sbin/netplan apply
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target
```


4. Activation du nouveau service

```
systemctl daemon-reload
systemctl enable netplan-apply.service
systemctl start netplan-apply.service
systemctl status netplan-apply.service
```


5. Utilisation du script au redémarrage

`nano /var/spool/cron/crontabs/root`

```
@reboot /usr/local/bin/configure-network.sh
```
## Initial Setup of Instance

This document goes over how the Oracle instance was set up, and provides a reference if the network configurations ever change.

### Connecting to instance

Use this command to connect to the instance

```bash
ssh -i <path to private key> ubuntu@<instance public-ip>
```

### Table of contents

- [Setting Up The Firewalls](#firewalls)
- [Setting Up The User](#ubuntu-user)
- [Building The .env File](#env-file)
- [Systemd Service Config](#systemd-service)
- [Caddy Proxy](#caddy)
- [Cloudflare DNS Setup](#cloudflare-dns)
- [Memory Protection (Swap File)](#swap-file-memory-overflow-protection)
- [Running The App](#running-the-website)

## Firewalls

### Firewall 1, Oracle level:

There are two firewalls in this instance. Below is the path to edit the Oracle level firewall:

Console -> Networking -> Virtual Cloud Networks -> Security Lists -> `terrariadle-server` -> Add Ingress Rules.

It should have the following configuration:

| Source CIDR | Protocol | Dest Port | Purpose |
| ----------- | -------- | --------- | ------- |
| `0.0.0.0/0` | TCP      | 22        | SSH     |
| `0.0.0.0/0` | TCP      | 80        | HTTP    |
| `0.0.0.0/0` | TCP      | 443       | HTTPS   |

### Firewall 2, Shape level:

The second firewall is contained within the shape. The `iptables-persistent` package should've saved current rules after restarts. Verify with:

```bash
sudo iptables -L INPUT -n --line-numbers
```

You should see the `80` and `443` ports on the response.

Here is how the ports were opened up:

```bash
sudo iptables -I INPUT -p tcp --dport 80 -j ACCEPT
sudo iptables -I INPUT -p tcp --dport 443 -j ACCEPT
sudo apt update
sudo apt install iptables-persistent -y
```

## Ubuntu User

These were the commands for setting up the user:

```bash
sudo useradd --system --no-create-home --shell /usr/sbin/nologin terrariadle
sudo mkdir -p /opt/terrariadle
sudo chown terrariadle:terrariadle /opt/terrariadle
```

To check if the user exists, you can run the following command:

```bash
getent passwd
```

## .env File

These are the commands used for the .env file:

```bash
sudo touch /etc/terrariadle.env
sudo chmod 600 /etc/terrariadle.env
sudo chown terrariadle:terrariadle /etc/terrariadle.env
```

To modify this .env file, run:

```bash
sudo vim /etc/terrariadle.env
```

## Systemd Service

This file is responsible for managing a Linux a process as a long-running service: when to start it, what user to run it as, what to do if it crashes, how to stop it cleanly, and what boundaries to enforce around it. I've attached the current configuration I'm using:

- [Systemd File](./systemd-file.txt)

This is stored in `/etc/systemd/system/terrariadle.service`. To modify this file, run:

```bash
sudo vim terrariadle.service
```

The actual binary file goes in `/opt/terrariadle`.

To check the securities of this file, you can also run:

```bash
systemd-analyze security terrariadle
```

## Caddy

Caddy is being used as a reverse proxy to support HTTPS and to map port `8080` to port `443`. Below is how I installed Caddy:

```bash
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy
```

Here is how I enabled it and set up logging:

```bash
sudo systemctl enable --now caddy
sudo systemctl reload caddy
journalctl -u caddy -f
```

And here is the Caddy file (`/etc/caddy/Caddyfile`):

```caddyfile
www.terrariadle.com {
    redir https://terrariadle.com{uri} permanent
}

terrariadle.net, www.terrariadle.net {
    redir https://terrariadle.com{uri} permanent
}

terrariadle.com {
    reverse_proxy /api/* localhost:8080 {
        header_up X-Real-IP {http.request.header.CF-Connecting-IP}
        flush_interval -1
    }

    @immutable {
        path /_app/immutable/*
    }
    header @immutable Cache-Control "public, max-age=31536000, immutable"

    @images {
        path *.png *.jpg *.jpeg *.webp *.avif *.svg *.gif *.ico
    }
    header @images Cache-Control "public, max-age=604800"

    reverse_proxy localhost:8080 {
        header_up X-Real-IP {http.request.header.CF-Connecting-IP}
    }

    encode zstd gzip

    header {
        Strict-Transport-Security "max-age=31536000; includeSubDomains"
        X-Content-Type-Options "nosniff"
        X-Frame-Options "DENY"
        Referrer-Policy "strict-origin-when-cross-origin"
        -Server
    }

    log {
        output file /var/log/caddy/terrariadle.log {
            roll_size 10mb
            roll_keep 5
        }
        format json
    }
}
```

_Note: if Caddy ever needs to be reset, ensure that the DNS proxy in CloudFlare is set to DNS mode for the initial Let's Encrypt Certification. You can turn Proxying on afterwards._

To restart the caddy service, run:

`sudo systemctl reload caddy`

## Cloudflare DNS

This is how the Terrariadle DNS domain was configured:

- **Type**: A
- **Name**: @ (represents the root domain, terrariadle.com)
- **IPv4 address**: Oracle instance public IP
- **TTL**: Auto

This `CNAME` record was made for www:

- **Type**: CNAME
- **Name**: www
- **Target**: terrariadle.com

## Swap File (Memory Overflow Protection)

There is a 1 GB swap file on the instance to handle memory spikes and bigger traffic. Below is the commands to check it:

```bash
`sudo swapon --show`
```

or

```bash
`free -h`
```

Below are the commands used to make the swap file:

```bash
sudo fallocate -l 1G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

## Running The Website

See the [Service Config](#systemd-service) for details on how the service is set up. To run the site, the following commands were run:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now terrariadle
sudo systemctl status terrariadle
journalctl -u terrariadle -f
```

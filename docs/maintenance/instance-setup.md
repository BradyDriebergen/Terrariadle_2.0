## How to configure Oracle Instance

This document goes over how the Oracle instance was set up, and provides a reference if the network configurations ever change.

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

```
sudo iptables -L INPUT -n --line-numbers
```

You should see the `80` and `443` ports on the response.

Here is how the ports were opened up:

```
sudo iptables -I INPUT -p tcp --dport 80 -j ACCEPT
sudo iptables -I INPUT -p tcp --dport 443 -j ACCEPT
sudo apt update
sudo apt install iptables-persistent -y
```

### Cloudflare DNS

This is how the Terrariadle DNS domain was configured:

- **Type**: A
- **Name**: @ (represents the root domain, terrariadle.com)
- **IPv4 address**: Oracle instance public IP
- **TTL**: Auto

This `CNAME` record was made for www:

- **Type**: CNAME
- **Name**: www
- **Target**: terrariadle.com

### Swap File (Memory Overflow Protection)

There is a 1 GB swap file on the instance to handle memory spikes and bigger traffic. Below is the commands to check it:

```
`sudo swapon --show`
```

or

```
`free -h`
```

Below are the commands used to make the swap file:

```
sudo fallocate -l 1G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

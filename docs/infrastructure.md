# Hosting Terrariadle

Another area that I haven't had much experience in is configuring and hosting server instances. The only experience I've had is with the first iteration of this project. I basically had AI (before the good models came out) configure it all for me. It was a mess of random network rules, complex port routing, and an inefficient way of hosting my site. I decided to start fresh on this release, making a new instance from scratch and understanding.

Going through this process, I found that all this complex network configuration isn't really too bad. There is a bit of a learning curve, but each of the rules break down into simple concepts.

## Why Oracle Cloud

Compared to other cloud providers, Oracle's Free tier is definitely the best free instance option. You can get 1 OCPU (equivalent to 2 vCPUs) and 1 GB of ram with no limit on time or compute hours. The trade-off is it's more manual than something like Railway or Heroku. Also, Oracle's Cloud website is super clunky to navigate through.

I originally planed on using the free ARM shape (`VM.Standard.A1.Flex`), but it wasn't available in my region. Apparently, it is super popular and notoriously capacity-constrained.

## Connecting

Connecting to instances is pretty easy. All you need to do is create an SSH connection between your machine and the cloud instance. For example, this is what I use:

```
ssh -i <path to private key> ubuntu@<instance public-ip>
```

`ubuntu` refers to the instance name. This was created because I'm using the Ubuntu image.

## Opening Ports on Both Firewalls

This is the part that tripped me up before. There are two independent firewall layers on these instances, and both need to be opened or outside traffic can't leave/enter.

### 1. OCI Security List (cloud-level)

Within Oracle, I added two rules for my instance (port 22 for SSH was already created). This opens up these ports on the machine so HTTP (port 80) and HTTPS (port 443) can send and receive information.

| Source CIDR | Protocol | Destination Port | Purpose |
| ----------- | -------- | ---------------- | ------- |
| `0.0.0.0/0` | TCP      | 22               | SSH     |
| `0.0.0.0/0` | TCP      | 80               | HTTP    |
| `0.0.0.0/0` | TCP      | 443              | HTTPS   |

The way I like to think about this is each port is a door to your machine. By default, you only have your front door (port 22), and that allows you to enter the house. These new rules added a couple mail slots on the door (ports 80 and 443). This doesn't allow others to enter the home, but it allows some communication with the outside world (HTTP/HTTPS requests).

### 2. Instance-level firewall (iptables)

Oracle's Ubuntu images ship with iptables rules that block everything except SSH by default. This firewall also needed to be updated to open ports 80 and 443:

```
sudo iptables -I INPUT -p tcp --dport 80 -j ACCEPT
sudo iptables -I INPUT -p tcp --dport 443 -j ACCEPT
sudo apt update
sudo apt install iptables-persistent -y
```

The `iptables-persistent` tool allows these rules to remain persistent through reboots.

This is like the analogy above. However, with both of these firewalls, it's more like the Oracle instance is a hotel and the shape is your hotel room. Requests can still come through, but it has to go through the front desk and your hotel room front door.

## DNS (Cloudflare)

I use Cloudflare for Domain registrar/DNS. Once my instance was made, I added DNS rules to point directly to my instance:

- **A record**: `@` -> instance public IP
- **CNAME**: `www` -> `terrariadle.com`

Another feature that Cloudflare offers is a proxy between their DNS servers and my domain. Before using this proxy, it was important for me to keep it DNS only so Caddy could complete its _Let's Encrypt_ challenge. This challenge just tests to see if my domain is really my domain. Terrariadle.com is already connected to my server, so it just needs to test in order to issue a TLS certificate. Once complete, I enabled this proxy once my certificate was issued. I also set Cloudflare's SSL/TLS mode to Full (strict) so the Cloudflare-to-origin leg is also encrypted.

## Swap File

As I stated above, my instance is only limited to 1 GB of memory. While I made my program more efficient, It's still a little too close for comfort. I looked into ways to mediate this, and found that you can make swap files in Linux.

If your familiar with Macs, you'll know when they run out of memory, it will use the storage as a temporary buffer. This means that if you're maxing out your memory, your applications won't crash because it will start using the SSD for the excess. SSDs are slower, but it's a no-brainer when dealing with a small amount of memory. A swap file performs the same way, adding a section of storage that's dedicated to taking care of memory overflow.

After learning about this, I added a 1 GB swap file as a buffer against memory spikes:

```
sudo fallocate -l 1G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

## Takeaway

Once I took the time to learn more about infrastructure and hosting this project, the concepts didn't seem that complicated. I am definitely more confident this time around because of the memory overflow safety and the knowledge over how my server is configured.

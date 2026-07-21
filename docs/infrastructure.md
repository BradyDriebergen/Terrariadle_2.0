# Hosting Terrariadle

Another area that I haven't had much experience in is configuring and hosting server instances. The only experience I've had is with the first iteration of this project. I basically had AI (before the good models came out) configure it all for me. It was a mess of random network rules, complex port routing, and an inefficient way of hosting my site. I decided to start fresh on this release, making a new instance from scratch and understanding.

Going through this process, I found that all this complex network configuration isn't really too bad. There is a bit of a learning curve, but each of the rules break down into simple concepts.

### Why Oracle Cloud?

Compared to other cloud providers, Oracle's Free tier is definitely the best free instance option. You can get 1 OCPU (equivalent to 2 vCPUs) and 1 GB of ram with no limit on time or compute hours. The trade-off is it's more manual than something like Railway or Heroku. Also, Oracle's Cloud website is super clunky to navigate through.

I originally planed on using the free ARM shape (`VM.Standard.A1.Flex`), but it wasn't available in my region. Apparently, it is super popular and notoriously capacity-constrained.

### Specific Commands

If you would like to see a more in-depth document about specific actions taken and commands ran, check out the [Instance Setup](./maintenance/instance-setup.md) document in the `/maintenance` folder.

# Setting up

## Opening Ports on Both Firewalls

This is the part that tripped me up before. There are two independent firewall layers on new Oracle instances, and both need to be opened or outside traffic can't leave/enter.

### 1. OCI Security List (cloud-level)

Within Oracle, I added two rules for my instance. I opened ports 80 and 443 on the machine so HTTP (port 80) and HTTPS (port 443) can send and receive information.

The way I like to think about this is each port is a door to your machine. By default, you only have your front door (port 22), and that allows you to enter the house (SSH). These new rules added a couple mail slots on the door (ports 80 and 443). This doesn't allow others to enter the home, but it allows certain communication with the outside world (HTTP/HTTPS requests).

### 2. Instance-level firewall (iptables)

Oracle's Ubuntu images ship with iptables rules that block everything except SSH by default. This firewall also needed to be updated to open ports 80 and 443.

This is like the analogy above. However, with both of these firewalls, it's more like the Oracle instance is a hotel and the shape is your hotel room. Requests can still come through, but it has to go through the front desk and your hotel room front door.

## Setting Up Instance

### Creating a new Linux user

I run Terrariadle under a dedicated, unprivileged system user rather than root or the instance account. The main reason behind this is security. If a vulnerability in the project or one of its dependencies ever gets exploited, the attacker only has the permissions of that one user. Since this new user has no home directory, no login shell, and access to only its own binary, config, and data directory, there's very little that an attacker can do.

Making a dedicated user also keeps permissions/access clear. The user owns its binary and its data directory. It keeps the instance account separate from the running service too, so if I run a bad command (`rm -rf`), it can't mess up production files it doesn't currently have open.

I've learned that this is a good method of practicing principle of least priviledge. This new user is the only one responsible for everything the Terrariadle service. This user also only has access to anything involving this service, and nothing else.

### Setting up the service

### Reverse Proxy (Caddy)

Project environment:

/etc/terrariadle.env
sudo vim /etc/terrariadle.env

`/etc/systemd/system/terrariadle.service`

Installing Caddy:

```bash
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy
```

```bash
sudo systemctl enable --now caddy
sudo systemctl reload caddy   # after editing the Caddyfile, no downtime
journalctl -u caddy -f
```

caddy file:

`/etc/caddy/Caddyfile`

```caddyfile
terrariadle.com {
	redir https://www.terrariadle.com{uri} permanent
}

www.terrariadle.com {
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

`sudo systemctl reload caddy`

Commands to run:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now terrariadle
sudo systemctl status terrariadle
journalctl -u terrariadle -f
```

## DNS (Cloudflare)

I use Cloudflare for Domain registrar/DNS. Once my instance was made, I added DNS rules to point directly to my instance:

- **A record**: `@` -> instance public IP
- **CNAME**: `www` -> `terrariadle.com`

Another feature that Cloudflare offers is a proxy between their DNS servers and my domain. Before using this proxy, it was important for me to keep it DNS only so Caddy could complete its _Let's Encrypt_ challenge. This challenge just tests to see if my domain is really my domain. Terrariadle.com is already connected to my server, so it just needs to test in order to issue a TLS certificate. Once complete, I enabled this proxy once my certificate was issued. I also set Cloudflare's SSL/TLS mode to Full (strict) so the Cloudflare-to-origin leg is also encrypted.

## Swap File

As I stated above, my instance is only limited to 1 GB of memory. While I made my program more efficient, It's still a little too close for comfort. I looked into ways to mediate this, and found that you can make swap files in Linux.

If your familiar with Macs, you'll know when they run out of memory, it will use the storage as a temporary buffer. This means that if you're maxing out your memory, your applications won't crash because it will start using the SSD for the excess. SSDs are slower, but it's a no-brainer when dealing with a small amount of memory. A swap file performs the same way, adding a section of storage that's dedicated to taking care of memory overflow.

After learning about this, I added a 1 GB swap file as a buffer against memory spikes:

```bash
sudo fallocate -l 1G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

## Building The Project:

```makefile
build:
	cd frontend && npm run build
	rm -rf internal/web/build
	cp -r frontend/build internal/web/build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-trimpath \
		-ldflags "-s -w -X main.version=$$(git describe --tags --always --dirty)" \
		-o bin/terrariadle .
```

- **CGO_ENABLED=0**: produces a fully static binary with no libc dependency. Means you don't have to worry about glibc version mismatches between your build machine and whatever's on the Oracle image, and it plays nicely with a minimal systemd unit (no dynamic linker surprises). Since you're using pure Go MongoDB driver (not cgo-based sqlite or similar), this should be a no-op functionally.
- **GOOS=linux GOARCH=amd64**: explicit cross-compile target for the E2.1.Micro shape. Only matters if you're building on your MacBook; skip it if you build on the instance itself.
- **trimpath**: strips local filesystem paths (like /Users/brady/...) from the compiled binary. Minor security/hygiene win, no runtime cost.
- **ldflags "-s -w"**: strips the symbol table and DWARF debug info. Cuts binary size by roughly a third, which matters more for the free tier's limited boot volume than for runtime RAM (Go binaries don't page in symbol tables at runtime), but smaller is still better for scp transfer time and disk headroom.

## Takeaway

Once I took the time to learn more about infrastructure and hosting this project, the concepts didn't seem that complicated. I am definitely more confident this time around because of the memory overflow safety and the knowledge over how my server is configured.

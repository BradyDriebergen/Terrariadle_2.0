# Hosting Terrariadle

### Table of Contents

- [Opening Ports](#opening-ports-on-both-firewalls)
- [Creating a New Linux User](#creating-a-new-linux-user)
- [Setting up a systemd Service](#setting-up-the-service)
- [Setting up Caddy](#reverse-proxy-caddy)
- [Logging](#logging)
- [Building](#building-the-project)

---

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

Features I never had in the past that I wanted in this iteration include process supervision, automatic restarts, and graceful shutdown handled by the OS. In my previous iteration, I used tmux sessions that had none of these features. In fact, my tmux sessions would constantly crash due to the massive overhead of not building my project. I wanted to use something industry standard and was optimized for hosting.

After researching solutions, I quickly came across Linux `systemd` services. These services were built for this application of work, running as a Linux service in the background of the OS. This system supported extra security, resilience, performance, and built-in logging. It was a no-brainer compared to other options.

Despite its benefits, there is quite a lot of configuration to do to run one of these services. Specifically, some of the things you need are:

1. A unit file at `/etc/systemd/system/<name>.service` defining how to run the process.
2. `ExecStart`: path to the binary/command to run.
3. `User` / `Group`: who the process runs as (ideally a dedicated, unprivileged system user).
4. `WorkingDirectory`: where the process runs from.
5. A restart policy like `Restart=on-failure` plus a `RestartSec` delay.
6. `WantedBy=multi-user.target` under [Install] so systemctl enable actually takes effect on boot.
7. `daemon-reload` after creating/editing the unit, then enable --now to activate it.

There are quite a bit more that goes into hosting a `systemd` service. I recommend reading [Running The Service](./running-the-service.md) to see the specifics to how I run my service.

### Reverse Proxy (Caddy)

With my project running on port 8080, I needed to map my traffic from port 443 (where HTTPS requests come in) to port 8080, so my app can serve the client with the requested data. This is what a reverse proxy is for.

I used to use `nginx` for my reverse proxy in my previous implementation. I am actually quite fond of `nginx`, but after some research, I found a more modern tool called `Caddy`. Caddy is super nice because it's a more hands off reverse proxy tool. It has automatic HTTPS certification, a more simple configuration file, and many features built in rather than having to be configured. This includes things like logging, headers, TLS settings, etc. I felt for a project of this scale, this made more sense than `nginx`.

Here is my current Caddy file configuration:

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

As you can see, Caddy is a lot simpler than the equivalent nginx config. This is great for my project because I don't have to worry about issue like renewing my cert. It's a great low-effort option for proxying, but doesn't come without some tradeoffs. Nginx is typically more efficient and has a lot bigger ecosystem when working with it. While nginx is typically a better option for performance at scale, Caddy provides me with a lightweight, super easy way to proxy my application.

## Logging

Rather than writing to a log file myself, `StandardOutput=journal` and `StandardError=journal` send everything my binary writes to stdout/stderr into **journald**, tagged with `SyslogIdentifier=terrariadle`:

```ini
StandardOutput=journal
StandardError=journal
SyslogIdentifier=terrariadle
```

I read it with `journalctl`:

```bash
journalctl -u terrariadle -f       # live tail
journalctl -u terrariadle -n 100   # last 100 lines
journalctl -u terrariadle -p err   # errors and above only
```

Since journald stores logs in a structured, queryable format, I get filtering by time, priority, and unit for free, without grepping through rotating flat files.

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

Infrastructure used to be one of those areas where I struggled with. There are so many nuances when it comes to building, configuring, hosting, and serving data to people across the internet.

Every step here maps to something I didn't understand when making the first iteration of this project. I didn't understand how firewalls worked, I didn't configure my cert renewal correctly, and I didn't know how to run a service. Looking back on it, I'm surprised my old site worked at all.

Once I took the time to learn more about infrastructure and hosting this project, the concepts didn't seem that complicated. I get where requests come in, and responses leave. I understand the tradeoffs of using specific serving platforms. I am starting to learn more about how to make my application better just from configuring my instance. This was a massive learning journey for me, and while I'm not a pro at it yet, I feel a lot better hosting this new project this time around.

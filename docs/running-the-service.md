# Running Terrariadle as a systemd Service

When I first launched my site, I didn't have a clue on how to run a program on a production environment. For a good few months after I deployed my old site, I was running the site off of a development server. This means that every line of code was sent from my program to clients.

This wasn't even the worst of it. I used to host my frontend and backend on separate `tmux` sessions. I didn't know how to make background processes in Linux, so I used tmux sessions to 'run background services'. It was wildly inefficient.

As an example of how bad my original hosting was, here is the script I used to run my frontend:

```bash
#!/usr/bin/env bash
set -euo pipefail

cd ~/Terrariadle/client

while true; do
  echo "[$(date)] starting Angular frontend..."
  serve -s build > /dev/null
  code=$?
  echo "[$(date)] ng serve exited with code $code; restarting in 2s..."
  sleep 2
done
```

I did this because my frontend would constantly crash, and I needed a way to auto-restart it. This method never worked though, I ripped it off of an old classmates project. I wasn't even using Angular, I was using React. My instance was full of stuff like this, and I wanted to completely upgrade the way I run my services.

## Why I chose a systemd service

When planning on deploying the new iteration of Terrariadle, I wanted features such as process supervision, automatic restarts, and graceful shutdown handled by the OS rather than tmux. `systemd` was a no-brainer for this. Instead of using a huge script, `systemd` services already comes baked in with all the features above, as well as other useful features involving security and fault-tolerance.

I recommend referencing my [Service File](./maintenance/systemd-file.txt) while you read through this. It contains all the rules I use for my service.

## Benefits

- **Starts on boot/restarts on crash**: `Restart=on-failure` means if the binary fails, `systemd` starts it back up after a short delay with no manual intervention.
- **Graceful shutdown for free**: `systemctl stop`/`restart` sends `SIGTERM`, which my backend uses as a signal for stopping my program:

    ```go
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    ```

- **Runs as a dedicated, unprivileged user**: Rather than running as root or the instance user, this program is ran through a dedicated user. This allows for better ownership over the process and better security in case of incident.

- **Secrets stay out of the unit file**: Environment files live in a separate, locked-down env file rather than being readable by anyone who can run `systemctl cat`.

## Resource limits

Since I'm sharing only 1GB of RAM with Caddy and the OS, I cap the service's memory usage directly in the unit:

```ini
MemoryMax=400M
MemoryHigh=350M
TasksMax=100
```

`MemoryHigh` is a soft throttle point. `MemoryMax` is a hard ceiling that gets the process OOM-killed if exceeded. This keeps a leak in my app from taking down the whole instance, Caddy included. `TasksMax` caps total threads/processes as a sanity ceiling.

## Security hardening

The binary is publicly reachable over the internet, so I added standard `systemd` securities to minimize malicious impact if anything ever went wrong inside the process:

```ini
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/opt/terrariadle
ProtectKernelTunables=true
ProtectKernelModules=true
ProtectControlGroups=true
RestrictSUIDSGID=true
RestrictNamespaces=true
LockPersonality=true
```

Some of the highlights:

- `ProtectSystem=strict` + `ReadWritePaths` mounts the whole filesystem read-only except the one directory the app actually needs to write to.
- `ProtectHome` blocks access to `/home`/`/root`.
- The `Protect*`/`Restrict*`/`Lock*` set closes off kernel tampering, namespace escapes, and privilege escalation paths.

I recommend reading more about these. It's fascinating learning about all the potential attack points of a service, and how to better defend against these attack points.

## Logging

Rather than writing to a log file myself, The following rules sends everything my binary writes to stdout/stderr into `journald`:

```ini
StandardOutput=journal
StandardError=journal
SyslogIdentifier=terrariadle
```

I can then read it with `journalctl`:

```bash
journalctl -u terrariadle -f
journalctl -u terrariadle -n 100
journalctl -u terrariadle -p err
```

Since journald stores logs in a structured, queryable format, I get filtering by time, priority, and unit for free, without grepping through flat files.

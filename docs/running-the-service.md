# Running Terrariadle as a systemd Service

## Why I chose a systemd service

I'm deploying Terrariadle to Oracle Cloud's free-tier AMD shape — 1 OCPU, 1GB RAM — so I wanted process supervision, automatic restarts, and graceful shutdown handled by the OS itself rather than a tmux session or a process manager like PM2. `systemd` is already PID 1 on the instance, so wrapping my Go binary in a unit file gets me all of that through one consistent interface (`systemctl` / `journalctl`) instead of a hand-rolled init script.

## Benefits

- **Starts on boot, restarts on crash** — `Restart=on-failure` means if the binary panics or exits non-zero, systemd brings it back up after a short delay, no manual intervention.
- **Graceful shutdown for free** — `systemctl stop`/`restart` send `SIGTERM`, which my app already listens for to drain in-flight requests and SSE connections before exiting:

    ```go
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    ```

    I set `TimeoutStopSec=15s` in the unit — a few seconds longer than my app's own 10-second shutdown timeout — so systemd never SIGKILLs the process before my graceful shutdown logic finishes.

- **Runs as a dedicated, unprivileged user**, not root or my login user:

    ```bash
    sudo useradd --system --no-create-home --shell /usr/sbin/nologin terrariadle
    ```

- **Secrets stay out of the unit file** — Mongo URI and other config live in a separate, locked-down env file (`chmod 600`) referenced via `EnvironmentFile=/etc/terrariadle.env`, rather than being readable by anyone who can run `systemctl cat`.

## Resource limits

Since I'm sharing 1GB of RAM with Caddy and the OS, I cap the service's memory usage directly in the unit:

```ini
MemoryMax=400M
MemoryHigh=350M
TasksMax=100
```

`MemoryHigh` is a soft throttle point; `MemoryMax` is a hard ceiling that gets the process OOM-killed if exceeded. This keeps a leak in my app from taking down the whole instance, Caddy included. `TasksMax` caps total threads/processes as a sanity ceiling.

## Security hardening

The binary is publicly reachable over the internet, so I added standard systemd sandboxing directives to minimize blast radius if anything ever went wrong inside the process:

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

In short: `ProtectSystem=strict` + `ReadWritePaths` mounts the whole filesystem read-only except the one directory the app actually needs to write to; `ProtectHome` blocks access to `/home`/`/root`; the `Protect*`/`Restrict*`/`Lock*` set closes off kernel tampering, namespace escapes, and privilege escalation paths. None of this is Terrariadle-specific — it's the general checklist for "single static Go binary with no legitimate need for broad system access." I can verify the whole set anytime with:

```bash
systemd-analyze security terrariadle
```

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

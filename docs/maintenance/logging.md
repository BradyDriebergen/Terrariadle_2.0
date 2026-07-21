## Logging

### Systemd

Within the [System](./systemd-file.txt) file, there are rules in place for tracking logging:

```ini
StandardOutput=journal
StandardError=journal
SyslogIdentifier=terrariadle
```

This means anything your Go binary writes to stdout/stderr (log.Println, fmt.Println, panics, etc.) gets captured by journald, systemd's centralized logging daemon.

You can Access these logs through `journalctl`:

```bash
journalctl -u terrariadle -f              # live tail, like tail -f
journalctl -u terrariadle -n 100          # last 100 lines
journalctl -u terrariadle --since "1 hour ago"
journalctl -u terrariadle --since today
journalctl -u terrariadle -p err          # only errors and above
journalctl -u terrariadle -o json-pretty  # structured JSON output, useful for scripting
```

These logs should remain persistant. Below are the commands used to save the logs after restarting:

```bash
sudo mkdir -p /var/log/journal
sudo systemd-tmpfiles --create --prefix /var/log/journal
sudo systemctl restart systemd-journald
```

### Caddy

Caddy's configuration file contains a section that tracks logging as well. This writes HTTP access logs: every request that hits Caddy, with method, path, status code, response time, client IP, etc. as newline-delimited JSON directly to that file, not through journald.

You can view these logs through the following commands:

```bash
tail -f /var/log/caddy/terrariadle.log
tail -f /var/log/caddy/terrariadle.log | jq .
jq 'select(.status >= 500)' /var/log/caddy/terrariadle.log     # only server errors
jq 'select(.request.uri | startswith("/api"))' /var/log/caddy/terrariadle.log  # API traffic only
```

If you want to see **process-level logs**, that is tracked through `journalctl`:

```bash
journalctl -u caddy -f
```

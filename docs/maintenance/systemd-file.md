```
[Unit]
Description=Terrariadle - Terraria daily puzzle site
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=terrariadle
Group=terrariadle
WorkingDirectory=/opt/terrariadle
EnvironmentFile=/etc/terrariadle.env
ExecStart=/opt/terrariadle/terrariadle
Restart=on-failure
RestartSec=5s

# Graceful shutdown - give your SSE broker time to drain connections
KillSignal=SIGTERM
TimeoutStopSec=10s

# Resource limits - important on a 1GB box, prevents one bad
# memory leak from taking down the whole instance (Caddy included)
MemoryMax=400M
MemoryHigh=350M
TasksMax=100

# Security hardening
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

# Logging goes to journald
StandardOutput=journal
StandardError=journal
SyslogIdentifier=terrariadle

[Install]
WantedBy=multi-user.target
```

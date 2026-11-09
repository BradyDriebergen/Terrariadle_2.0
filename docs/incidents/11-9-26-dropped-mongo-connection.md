# MongoDB Atlas & Oracle Cloud Connectivity Fixes

_Incident Date: November 3, 2026_

This document serves as a comprehensive reference guide summarizing the root cause analysis, application constraints, and infrastructure-level configurations required to resolve stale database connection drops between a Go application hosted on Oracle Cloud Infrastructure (OCI) Always-Free tiers and MongoDB Atlas.

---

## 1. Environment & Setup Reference

The application environment consists of a modern, low-privilege backend architecture constrained by strict infrastructure network boundaries:

- **Compute Instance:** Oracle Cloud Infrastructure (OCI) Always-Free Shape (1 AMD EPYC fractional/shared vCPU).
- **Operating System Layer:** Linux managed natively via a **systemd service** utilizing a **dedicated system user** (non-root).
- **Application Runtime:** Go 1.25.0
- **Database Client Driver:** `go.mongodb.org/mongo-driver` v1.17.9
- **Database Provider:** MongoDB Atlas (Free/Flex/Shared tier cluster).

---

## 2. Root Cause Analysis: Why This Issue Occurs

The sudden connectivity blackouts that surfaced during the November 3, 2026 incident are caused by a mismatch between infrastructure-level firewall rules and default operating system configurations:

### The OCI Virtual Firewall Drop

Oracle Cloud’s Virtual Cloud Network (VCN) enforces an aggressive **idle TCP connection timeout (typically 360 seconds / 6 minutes)**. If a network socket remains completely quiet without traffic for this duration, the OCI firewall silently discards the connection tracking state.

### The Linux Kernel Blind Spot

By default, the Linux operating system kernel keeps TCP sockets alive by sending probe packets only after **2 hours (7,200 seconds)** of total inactivity. Because the Go driver delegates TCP keep-alive handling directly to the host OS, your Go application believes the socket is completely healthy. It is left with a connection pool full of "half-open" dead sockets that can neither send nor receive data.

### The Triggering Catalysts (Late 2026 Platform Updates)

This issue suddenly became apparent due to concurrent platform events:

1.  **Strict Firewall Hardening:** OCI recently standardized and hardened virtual network control planes across Always-Free hypervisors, strictly terminating idle connections.
2.  **Atlas Fleet Updates:** MongoDB Atlas continuously rolls out background maintenance, security updates, and rolling node restarts. When Atlas shifts cluster topologies, drivers attempt to re-evaluate endpoints. If the local network sockets are already silently dead, the background driver loops crash into stale timeouts.

---

## 3. Step-by-Step Fix Implementation Checklist

Perform these three phases sequentially to apply and lock down the permanent fix.

### Phase 1: Host OS Kernel Tweaks (Root Level)

Force the underlying Oracle Cloud host to actively prove its presence to the network before the OCI firewall cuts off the connection.

1. SSH into the Oracle Cloud instance.
2. Create an isolated system configuration profile:
    ```bash
    sudo nano /etc/sysctl.d/99-mongodb-keepalive.conf
    ```
3. Paste the following optimal network parameters to probe idling sockets every 2 minutes:
    ```ini
    # Aggressively probe connections after 120 seconds of idling
    net.ipv4.tcp_keepalive_time = 120
    # Wait 15 seconds between individual keepalive probes
    net.ipv4.tcp_keepalive_intvl = 15
    # Drop the connection only after 5 consecutive unacknowledged probes
    net.ipv4.tcp_keepalive_probes = 5
    ```
4. Save, close the editor (`Ctrl+O`, `Enter`, `Ctrl+X`), and load the parameters immediately:
    ```bash
    sudo sysctl --system
    ```

### Phase 2: systemd Service File Hardening

Provide the dedicated system user the resources required to handle background driver pools and guarantee automatic lifecycle recovery.

1. Open your systemd unit definition file:
    ```bash
    sudo nano /etc/systemd/system/your-app.service
    ```
2. Inject file descriptor headroom and automatic recovery strategies inside the `[Service]` block:

    ```ini
    [Service]
    Type=simple
    User=your-system-user
    Group=your-system-group
    ExecStart=/usr/local/bin/your-go-binary

    # --- STABILITY & RESOURCE OVERRIDES ---
    # Prevent the low-privilege user from hitting the default 1024 open file socket ceiling
    LimitNOFILE=65535

    # Automatically restart the application if a network drop panics the runtime
    Restart=always
    # Wait 5 seconds before a restart attempt to prevent CPU spin-locking loops
    RestartSec=5s
    ```

3. Save, exit, and reload the systemd daemon:
    ```bash
    sudo systemctl daemon-reload
    ```

### Phase 3: Go Application Code Optimization

Align your connection pool constraints with the physical reality of a shared 1 vCPU environment.

1. Verify that your code instantiates **one single global `mongo.Client` as a singleton**. Never call `mongo.Connect` inside transient HTTP routers or short-lived helper functions.
2. Refactor your `options.Client()` logic to match the optimized pooling configuration:
    ```go
    // Optimized pool initialization for 1 vCPU environments and strict cloud firewalls
    opts := options.Client().
        ApplyURI(uri).
        // Keep background monitoring threads warm and active
        SetMinPoolSize(2).
        // Limit context-switching overhead on a fractional/shared vCPU instance
        SetMaxPoolSize(15).
        // Prune unused connections cleanly inside our 120s keep-alive window
        SetMaxConnIdleTime(2 * time.Minute).
        // Hard operational deadline stop to prevent application-wide stalling
        SetTimeout(5 * time.Second)
    ```
3. Re-compile your application binary and copy it to your server destination.
4. Restart the service to apply the environment adjustments:
    ```bash
    sudo systemctl restart your-app.service
    ```

---

## 4. Engineering Evaluation: Why the Fix Works

This multi-tiered architectural fix operates at three independent layers to guarantee continuous uptime:

[Go App: Pool Pruning (2 min)] -> [Linux Kernel: Keep-Alive Probes (120s)] -> [OCI VCN Firewall: (Allowed)] -> [Atlas Cluster]

- **At the Network Layer:** By changing `net.ipv4.tcp_keepalive_time` to 120 seconds, the Oracle OS sends miniature background network probes every two minutes. This traffic tricks the OCI firewall into believing the connection is actively processing data, preventing the 6-minute idle deletion rule from ever tripping.
- **At the System Layer:** Elevating `LimitNOFILE` to 65535 removes the invisible ceiling that stops low-privilege system users from maintaining complex connection pools. The addition of `Restart=always` ensures that even if Atlas drops connections during severe server updates, systemd acts as a hypervisor to reboot the app cleanly.
- **At the Application Layer:** Lowering `MaxPoolSize` from its default of 100 down to 15 stops your Go runtime from choking your 1 shared vCPU with massive multi-threaded context switching during request spikes. `SetMinPoolSize(2)` forces the driver to maintain warm pipelines that catch DNS and topology adjustments instantly.

---

## 5. Key Takeaways & Lessons Learned

1.  **Drivers Are Dependent on the Host OS:** Application-level database connection properties like `SetMaxConnIdleTime` only control how the client behaves inside its own sandbox. They cannot stop structural cloud infrastructure firewalls from silently dropping underlying network routes.
2.  **Compute Scale Controls App Parameters:** Default driver settings (like a max pool size of 100) are engineered for multi-core systems. On specialized free tiers or micro-shapes with fractional vCPUs, pool caps must be restricted to prevent thread thrashing.
3.  **Permissions Limit Infrastructure Settings:** When running systemd under specific system users, any global OS network tuning (like `sysctl`) must be strictly applied via root privileges, while process-specific constraints (`LimitNOFILE`) must be declared directly inside the service manager unit.

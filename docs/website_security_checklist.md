# 🔒 Website Security Checklist (SvelteKit + Go + Oracle Cloud)

This checklist is tailored for a **daily puzzle game** architecture built with **SvelteKit (frontend)** and **Go (backend)**, hosted on an **Oracle Cloud free-tier instance**.

---

## ✅ Do Now (High Impact, Low Friction)

### Perimeter & Hosting
- [ ] HTTPS via Let’s Encrypt + HSTS (min 6 months)
- [ ] Firewall: allow only ports 80 (ACME) and 443; deny all others
- [ ] SSH: key-only, no password auth, non-root user, fail2ban
- [ ] Enable automatic OS security updates

### Architecture & API Surface
- [ ] Use SvelteKit `.server` endpoints as a proxy to Go backend
- [ ] Keep Go backend on private network or bind to localhost
- [ ] Separate GET (reads) and POST (writes)
- [ ] Auth or anti-abuse controls for write routes

### Secrets & Config
- [ ] Store secrets in environment variables (never in repo)
- [ ] Rotate credentials periodically
- [ ] Separate prod/stage/dev credentials

### Input, Output, Errors
- [ ] Validate input (length, type, regex)
- [ ] Limit body size and request timeout
- [ ] Sanitize error messages (no internal info)
- [ ] Always set `Content-Type: application/json; charset=utf-8`

### Abuse & Rate Limiting
- [ ] Global, per-IP, and per-user rate limits
- [ ] Prevent replay attacks (nonces/timestamps)
- [ ] Limit guesses per puzzle/day

### Browser Security Headers (SvelteKit)
- [ ] Content-Security-Policy (CSP)
- [ ] X-Frame-Options: DENY
- [ ] Referrer-Policy: no-referrer
- [ ] X-Content-Type-Options: nosniff
- [ ] Permissions-Policy: disable unused APIs
- [ ] Cross-Origin-Opener-Policy / Cross-Origin-Resource-Policy

### Cookies, CSRF, CORS
- [ ] Cookies: HttpOnly, Secure, SameSite=Lax/Strict
- [ ] CSRF: use tokens or SameSite + Origin checks
- [ ] CORS: restrict to your domain only

### Logging & Monitoring
- [ ] Structured logs (req ID, user ID, IP, latency)
- [ ] Rotate logs regularly
- [ ] Health endpoints and 5xx alerts

---

## 🔜 Do Next (Secure Defaults → Resilience)

### Network & Edge
- [ ] Use CDN/WAF (Cloudflare, Fastly, Oracle WAF)
- [ ] Use mTLS or HMAC-signed requests between SvelteKit and Go

### Game/Puzzle Integrity (Anti-cheat)
- [ ] Verify answers server-side only
- [ ] Use opaque, non-sequential puzzle IDs
- [ ] Server-enforced attempt limits
- [ ] Honeypot or proof-of-work for abuse detection
- [ ] Validate leaderboard scores with signed receipts

### Data & Backups
- [ ] Daily encrypted backups (DB + assets)
- [ ] Test restores monthly
- [ ] TTL for telemetry / non-essential logs

### Build & Dependencies
- [ ] Run npm audit / govulncheck
- [ ] Pin versions, signed builds, SBOM if possible

### SLOs & Timeouts
- [ ] End-to-end request timeouts
- [ ] Cache puzzle reads where safe

---

## ✨ Nice-to-Have (Defense in Depth)

- [ ] Subresource Integrity (SRI) for 3rd-party assets
- [ ] Feature flags and kill switches
- [ ] Canary deploys and rollbacks
- [ ] Security.txt + vulnerability disclosure
- [ ] mTLS to DB/Redis, role-based DB users

---

## SvelteKit Best Practices
- [ ] Secrets and API calls in `.server.ts` or `.page.server.ts`
- [ ] Add CSP/HSTS headers via `handle` hook
- [ ] Manage auth via `locals.user` (in `hooks.server.ts`)
- [ ] Require POST + CSRF token for mutations

---

## Go Backend Best Practices
- [ ] Middleware: recoverer, request ID, timeout, rate limit
- [ ] Use `DisallowUnknownFields` in JSON decode
- [ ] Set context deadlines per request
- [ ] Constant-time token/sig comparisons
- [ ] Health/readiness probes

---

## Oracle Cloud Hardening
- [ ] UFW: default deny, allow 443 (and 80 temporarily for ACME)
- [ ] Non-root deploy user; minimal sudo access
- [ ] Enable unattended security updates
- [ ] Disable unused services (`ss -tulpn`)
- [ ] Set up logrotate and disk alerts

---

## “What Good Looks Like”
- Private Go API (localhost, VPC, or allowlist)
- SvelteKit proxy enforces auth, CSRF, rate limiting
- Puzzle verification server-side only
- CDN/WAF front-end with caching for puzzle assets
- Monitoring: submits/min, win rate, latency, rate-limit hits

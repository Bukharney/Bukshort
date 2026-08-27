# Bukshort - Homelab Linux ARM Deployment Guide

This guide provides instructions for deploying **Bukshort** (Go backend, SvelteKit frontend, and PostgreSQL database) on ARM-based homelab environments (e.g. Raspberry Pi 4/5, Rock 5, Odroid, ARM64 Proxmox LXC/VMs, or any Linux ARM system running Docker) behind **Cloudflare Tunnel (`cloudflared`)**.

---

## Architecture & Cloudflare Tunnel Flow

```
[ Public User ]
       |
       v (HTTPS / SSL via Cloudflare Edge)
[ Cloudflare Global Network ]
       |
       v (Encrypted Tunnel)
[ cloudflared Docker Container / Service ]
       |
       +---> Frontend (SvelteKit) : http://frontend:3000 (e.g., https://short.jirapat.tech)
       |
       +---> Backend (Go API)     : http://backend:8080  (e.g., https://api.short.jirapat.tech)
```

---

## Cloudflare Zero Trust Setup (Step-by-Step)

### Step 1: Create a Cloudflare Tunnel
1. Log in to [Cloudflare Zero Trust Dashboard](https://one.dash.cloudflare.com/).
2. Go to **Networks** -> **Tunnels**.
3. Click **Create a Tunnel** -> Select **Cloudflared**.
4. Name your tunnel (e.g., `bukshort-homelab`).
5. Copy your **Tunnel Token** (starts with `eyJh...`).

### Step 2: Add Public Hostname Ingress Rules
Under the **Public Hostnames** tab of your tunnel, add two rules:

1. **Frontend Rule:**
   - **Subdomain:** `short` (or `@` for root domain)
   - **Domain:** `jirapat.tech`
   - **Service Type:** `HTTP`
   - **URL:** `frontend:3000` (or `localhost:3000` if running outside Docker container)

2. **Backend API Rule:**
   - **Subdomain:** `api.short` (or `api-short`)
   - **Domain:** `jirapat.tech`
   - **Service Type:** `HTTP`
   - **URL:** `backend:8080` (or `localhost:8080` if running outside Docker container)

---

## Deployment Options

### Option A: Automated Cloudflare Tunnel via Docker Compose (Recommended)

Simply paste your `TUNNEL_TOKEN` in `.env` and `docker compose` will run `cloudflared` automatically alongside your app.

1. **Clone repository and create `.env`:**
   ```bash
   git clone https://github.com/Bukharney/Bukshort.git
   cd Bukshort
   cp .env.example .env
   ```

2. **Configure `.env`:**
   ```env
   POSTGRES_PASSWORD=your_secure_password
   API_URL=https://api.short.jirapat.tech/
   TUNNEL_TOKEN=eyJh...paste_your_cloudflare_tunnel_token_here...
   ```

3. **Start production stack with GHCR pre-built images:**
   ```bash
   docker compose -f docker-compose.prod.yml up -d
   ```
   *(Or build locally: `docker compose up -d --build`)*

---

### Option B: Using an Existing Host-Level `cloudflared` Daemon

If you already run `cloudflared` directly on your host machine or via another Docker stack, point your ingress rules to `http://localhost:3000` (frontend) and `http://localhost:8080` (backend).

---

## Access & Verification

- **Public Frontend Web App:** `https://short.jirapat.tech`
- **Public API Server:** `https://api.short.jirapat.tech`
- **Swagger Documentation:** `https://api.short.jirapat.tech/docs/index.html`

---

## Useful Maintenance Commands

- **Check Tunnel Status:**
  ```bash
  docker logs -f bukshort_cloudflared
  ```
- **Update Application & Images:**
  ```bash
  docker compose -f docker-compose.prod.yml pull
  docker compose -f docker-compose.prod.yml up -d
  ```
- **Backup Database:**
  ```bash
  docker exec -t bukshort_db pg_dump -U postgres bukshort > backup.sql
  ```

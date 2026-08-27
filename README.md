# Bukshort — Production-Ready URL Shortener

[![Go Version](https://img.shields.io/badge/Go-1.22-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![SvelteKit](https://img.shields.io/badge/SvelteKit-2.0-FF3E00?style=flat-square&logo=svelte)](https://kit.svelte.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14--Alpine-4169E1?style=flat-square&logo=postgresql)](https://www.postgresql.org/)
[![Docker Multi-Arch](https://img.shields.io/badge/Docker-ARM64%20%7C%20AMD64-2496ED?style=flat-square&logo=docker)](https://www.docker.com/)
[![CI/CD Pipeline](https://img.shields.io/badge/GitHub%20Actions-CI%2FCD-2088FF?style=flat-square&logo=githubactions)](https://github.com/features/actions)

**Bukshort** is a modern, high-performance, full-stack URL shortening web application. Built with a **Go (Gin framework)** REST API backend, a **SvelteKit** single-page application with SSR server proxying, **PostgreSQL** database persistence, and multi-architecture **Docker Compose** containerization designed for homelab ARM64 and cloud deployments.

---

## 📐 System Architecture

Bukshort is designed with a **single-domain reverse proxy architecture**. The SvelteKit SSR Node engine acts as the primary public gateway, forwarding API and redirection requests to the Go backend via Docker's internal container network (`app_net`).

```mermaid
graph TD
    User["Public User / Client Browser"] -->|HTTPS| CF["Cloudflare Tunnel (cloudflared)"]
    CF -->|Port 3000| FE["SvelteKit Frontend (Node.js Adapter)"]
    
    subgraph Docker Network: app_net
        FE -->|SSR Server Proxy /shorten| BE["Go Backend API (Gin Server :8080)"]
        FE -->|SSR Server Proxy /:shortURL| BE
        FE -->|SSR Server Proxy /docs/*| BE
        BE -->|GORM Driver / SQL| DB[("PostgreSQL Database :5432")]
    end
```

---

## 🔄 Request Sequence Flow

```mermaid
sequenceDiagram
    autonumber
    actor User as User Browser
    participant FE as SvelteKit Server (Node)
    participant BE as Go Backend (Gin)
    participant DB as PostgreSQL DB

    Note over User, DB: 1. URL Shortening Request
    User->>FE: POST /shorten (url)
    FE->>BE: Forward POST http://backend:8080/shorten
    BE->>DB: Query existing original_url
    alt URL exists in DB
        DB-->>BE: Return existing short_url
    else New URL
        BE->>BE: Generate Random Base62 Shortcode (6 chars)
        BE->>DB: INSERT into links (short_url, original_url)
        DB-->>BE: Success
    end
    BE-->>FE: JSON { url: short_url }
    FE-->>User: 200 OK + QR Code Data

    Note over User, DB: 2. Short URL Redirection
    User->>FE: GET /aB3xD9
    FE->>BE: Forward GET http://backend:8080/aB3xD9
    BE->>DB: Query short_url = aB3xD9
    DB-->>BE: Found original_url
    BE-->>FE: 302 Redirect (Location: destination_url)
    FE-->>User: 302 Redirect to Original Destination
```

---

## 🛠️ Key Engineering & Architectural Highlights

### 1. Zero-CORS & Single-Domain SSR Server Proxying
To prevent browser **Cross-Origin Resource Sharing (CORS)** and **Private Network Access (PNA loopback)** errors, SvelteKit server endpoints (`+server.ts`) proxy API calls (`/shorten`, `/:shortURL`, and `/docs/*any`) directly to `http://backend:8080` on Docker's internal network.
- **Result:** Users interact with a single domain (`https://short.jirapat.site`). No secondary API subdomains or open cross-origin policies are required.

### 2. High-Performance Go Backend & Collision Handling
- Built with **Gin** for minimal latency and **GORM** ORM targeting PostgreSQL.
- Shortcodes are generated using a Base62 character set (`a-z`, `A-Z`, `0-9`).
- Database constraints enforce unique indexes on both `original_url` and `short_url` with an automated retry loop to guarantee collision safety under high request concurrency.

### 3. Multi-Stage & Multi-Architecture Docker Containers
- Both **Backend** (`golang:1.22-alpine` -> `alpine:latest`) and **Frontend** (`node:20-alpine` with `pnpm`) use multi-stage Docker builds.
- Native support for **ARM64** (Raspberry Pi 4/5, Apple Silicon, SBCs, Proxmox ARM LXCs) and **AMD64** architectures.

### 4. Automated CI/CD & GHCR Publishing
- A custom **GitHub Actions workflow** (`.github/workflows/ci-cd.yml`) automatically builds and tests both Go backend and SvelteKit frontend on every pull request.
- Pushes to `main` use **Docker Buildx + QEMU** to publish multi-platform container images directly to **GitHub Container Registry (`ghcr.io`)**.

---

## 🧰 Tech Stack

| Layer | Technologies Used |
| :--- | :--- |
| **Frontend** | SvelteKit 2, Svelte 4, TypeScript, TailwindCSS, PostCSS, `@zerodevx/svelte-toast`, `qrcode` |
| **Backend** | Go 1.22, Gin Web Framework, GORM ORM, `swaggo/gin-swagger` (OpenAPI) |
| **Database** | PostgreSQL 14 Alpine |
| **DevOps & Infrastructure** | Docker, Docker Compose, GitHub Actions CI/CD, Cloudflare Tunnel (`cloudflared`) |

---

## 📡 API Specification

Interactive Swagger UI documentation is served automatically at `/docs/index.html`.

| Method | Endpoint | Description | Request Body / Params | Response |
| :--- | :--- | :--- | :--- | :--- |
| `POST` | `/shorten` | Shortens a long URL | `{ "url": "https://example.com" }` | `{ "url": "aB3xD9" }` |
| `GET` | `/:shortURL` | Redirects to original destination | Path param: `shortURL` | `302 Found` (Redirect) |
| `GET` | `/docs/*any` | Swagger API Documentation UI | None | `200 OK` (HTML/CSS/JS) |

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/) (v1.21+)
- [Node.js](https://nodejs.org/) (v20+) & [pnpm](https://pnpm.io/)
- [Docker](https://www.docker.com/) & Docker Compose (v2)

---

### Local Development Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Bukharney/Bukshort.git
   cd Bukshort
   ```

2. **Start Backend (Go):**
   ```bash
   cd backend
   cp .env.example .env # Set DB_URL and PORT
   go run main.go
   ```

3. **Start Frontend (SvelteKit):**
   ```bash
   cd frontend
   pnpm install
   pnpm dev
   ```

   Access SvelteKit dev server at `http://localhost:5173`.

---

### Production Deployment via Docker Compose

To deploy pre-built multi-architecture images directly from **GHCR**:

1. **Create environment file:**
   ```bash
   cp .env.example .env
   ```

2. **Run Production Containers:**
   ```bash
   docker compose -f docker-compose.prod.yml up -d
   ```

3. **Verify running services:**
   ```bash
   docker compose -f docker-compose.prod.yml ps
   ```

   - **Frontend App:** `http://localhost:3000`
   - **Backend API:** `http://localhost:8080` (Internal access)

---

## 🔒 Cloudflare Tunnel Setup (Homelab)

Bukshort runs securely behind Cloudflare Tunnel without exposing host ports to the internet:

1. **Create Tunnel** in [Cloudflare Zero Trust Dashboard](https://one.dash.cloudflare.com/).
2. **Add Public Hostname Ingress Rule:**
   - **Domain:** `short.jirapat.site`
   - **Service:** `HTTP` -> `localhost:3000` (or `frontend:3000`)
3. SvelteKit SSR will handle proxying API calls to backend internally.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

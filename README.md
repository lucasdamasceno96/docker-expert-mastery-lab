# Docker Expert Mastery Lab

This repository is a hands-on laboratory designed to bridge the gap between intermediate Docker usage and Senior/Specialist expertise. The goal is to master container internals, production-grade security, performance optimization, and advanced orchestration patterns.

## 🎯 Objectives

- Build production-ready, secure, and slim images.
- Understand Docker internals (Namespaces, Cgroups, Storage Drivers).
- Implement advanced networking and persistence strategies.
- Master troubleshooting and observability in containerized environments.
- Prepare for high-level technical interviews.

## 🛠 Project Structure

- All code, comments, and documentation are in **English**.
- Teaching and theoretical discussions are conducted in **Portuguese** (via Pair Programming with LLM).

---

## 🚀 The Mastery Roadmap

### Phase 1: Image Optimization & BuildKit

- **Exercise 01:** Multi-stage builds for Go. Focus on layer caching and distroless images.
- **Exercise 02:** Advanced BuildKit features (Secret mounts, SSH mounts, and Cache imports/exports).

### Phase 2: Docker Internals & Runtime isolation

- **Exercise 03:** Manual exploration of Namespaces and Cgroups. Process isolation validation.
- **Exercise 04:** Understanding the Container Lifecycle: Init processes, Signals (SIGTERM/SIGKILL), and Zombie processes.

### Phase 3: Advanced Networking

- **Exercise 05:** Custom Bridge networks, DNS resolution, and IPAM (IP Address Management).
- **Exercise 06:** Network Troubleshooting: Using `tcpdump` and `nsenter` to debug container traffic.

### Phase 4: Persistence & Storage Drivers

- **Exercise 07:** Deep dive into `overlay2` and storage driver performance.
- **Exercise 08:** Advanced Volume Management: Drivers, Backups, and Read-only mounts.

### Phase 5: Hardening & Security

- **Exercise 09:** Non-root execution, Capabilities (`cap-drop`/`cap-add`), and Seccomp profiles.
- **Exercise 10:** Supply Chain Security: Image signing, Vulnerability scanning (Trivy), and SBOMs.

### Phase 6: Observability & Troubleshooting

- **Exercise 11:** Healthcheck strategies and Logging Drivers (Fluentd/Splunk/Gelf).
- **Exercise 12:** Performance profiling: `docker stats` internals and memory/CPU limit testing.

### Final Project

- **🏗️ Exercise 13: Production-Ready Hardened Stack**

**What we built in Exercise 13:**

1. **Gateway (Nginx)**
   - Based on `nginx:1.25-bookworm` (Debian only).
   - Read-only filesystem and tmpfs for `/tmp`, `/var/run`, `/var/cache/nginx`.
   - Proxy to internal backend using secure networking.
   - Healthcheck ensured by container startup and proper proxy.

2. **Backend API (Python HTTP Server)**
   - Based on `python:3.11-slim-bookworm`.
   - Serves real content from `/app` folder (mounted read-only).
   - Internal network only (`backend_net`), cannot access external hosts.
   - Healthcheck via native Python socket connection.

3. **Database (Postgres)**
   - Based on `postgres:16-bookworm`.
   - Internal network only with healthcheck using `pg_isready`.
   - Credentials and database configured via environment variables.

4. **Networking**
   - Two isolated networks:
     - `frontend_net`: exposes gateway to host.
     - `backend_net`: internal only, shared between backend and database.

5. **Security & Hardening**
   - Read-only gateway filesystem.
   - tmpfs for runtime directories.
   - Non-root Nginx user (`nginx`).
   - Backend cannot reach external networks.
   - Healthchecks validate container readiness without extra tools.

---

## 👨‍💻 How to use this Lab

This is a **Pair Programming** journey. For each exercise:

1. Initialize the discussion with the LLM Mentor.
2. Implement the solution in a dedicated directory (e.g., `/exercises/ex01-multi-stage`).
3. Request a Senior Code Review.
4. Document the "Specialist Takeaways" in each exercise folder.

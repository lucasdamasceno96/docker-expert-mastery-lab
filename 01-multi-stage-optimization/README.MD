# Exercise 01: Multi-Stage Optimization

This module demonstrates the transition from a standard Docker setup to a **production-grade, high-performance container environment**. The goal was to build a secure, minimal Go API using SQLite without compromising on speed or security.

## 🚀 The Project

A Go-based REST API built with **Clean Architecture** principles, featuring:

- **Health Check Endpoint**: `GET /health`
- **Persistence Layer**: SQLite (using `modernc.org/sqlite` for a Pure Go, CGO-free implementation).
- **Log Management**: `POST /logs` to store request data.

---

## 🛠 Mastery Concepts Applied

### 1. Multi-Stage Build Strategy

We use the `AS builder` pattern to separate the **build environment** from the **runtime environment**.

- **Builder Stage**: Includes the Go compiler, git, and build caches.
- **Runner Stage**: A hardened Alpine image containing _only_ the final static binary.
- **Result**: Reduced image size from ~800MB (standard Go image) to ~15MB.

### 2. Advanced Layer Caching

By copying `go.mod` and `go.sum` before the rest of the source code, we ensure that:

- Dependencies are only re-downloaded if the modules change.
- Code changes (frequent) do not invalidate the dependency layer (slow).

### 3. Static Compilation & Binary Stripping

Using `CGO_ENABLED=0` and `-ldflags="-s -w"`, we created a **Static Binary**:

- **Portability**: It runs without any external library dependencies (no `libc`).
- **Security**: Stripping symbols makes reverse engineering harder and reduces size.

### 4. Hardened Security (Non-Root)

The container does not run as `root`. We implemented:

- **Dedicated User**: `appuser` with limited permissions.
- **Read-Only mindset**: The application only has write access to the `/data` volume for the SQLite database.

---

## 📦 How to Build & Run

### Build with BuildKit (IPv4 optimized)

```bash
docker build --network=host -t mastery-api:v1 .
```

````

### Run with Persistent Volume

```bash
docker run -d \
  --name mastery-app \
  -p 8080:8080 \
  -v sqlite_data:/data \
  mastery-api:v1

```
---

````

# Exercise 02: Advanced BuildKit - Secure Secrets & SSH Mounts

This module demonstrates how to use **Docker BuildKit** to handle sensitive data during the image construction phase without leaving any traces in the final image layers.

## ⚙️ Core Concepts

### 1. Secret Mounts (`--mount=type=secret`)

Traditional methods like `ARG` or `ENV` persist sensitive data in the image metadata. By using BuildKit Secrets, we mount a temporary **tmpfs** (RAM-based filesystem) that exists only during the execution of a specific `RUN` command.

- **Zero-Persistence:** The secret is never written to the container's disk.
- **Security:** Prevents accidental exposure via `docker history`.

### 2. SSH Mounts (`--mount=type=ssh`)

In enterprise environments, builds often require private dependencies (e.g., internal Git repos).

- **Identity Forwarding:** Instead of copying private keys into the Dockerfile, we "borrow" the host's SSH agent.
- **Isolation:** The build process can authenticate with GitHub/GitLab using the developer's or CI's identity, but the key never enters the container.

## 🛠️ Infrastructure Stack

- **CI/CD:** GitHub Actions.
- **Secret Management:** GitHub Secrets (encrypted at rest).
- **Engine:** Docker Buildx (BuildKit).

## 🚀 Impact on Production

Using this pattern reduces the **Attack Surface** of your containerized applications by ensuring that API Keys, Tokens, and Private Keys are strictly ephemeral.

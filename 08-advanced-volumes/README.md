# Exercise 08: Advanced Volume Management & Data Persistence

## 🎯 Objectives

- Master the difference between Bind Mounts and Named Volumes.
- Implement security hardening using Read-Only mounts.
- Perform "on-the-fly" backups using Sidecar containers.

## 🛠️ Key Commands

- **Create Volume:** `docker volume create my-vol`
- **Mount RO:** `-v /host/path:/container/path:ro`
- **Backup Strategy:** `docker run --rm --volumes-from <container> -v $(pwd):/backup alpine tar cvf /backup/data.tar /data-path`

## 💡 Specialist Takeaways

1. **Performance:** Volumes bypass the `overlay2` storage driver, providing native I/O speed.
2. **Decoupling:** Named Volumes allow you to upgrade or replace containers without worrying about the underlying host filesystem structure.
3. **Immutability:** Always mount configuration files as `:ro` (read-only) to prevent runtime tampering.
4. **Tmpfs:** For ultra-fast, non-persistent data (like secrets or cache), use `--tmpfs`. It stores data in the Host's RAM, never touching the disk.

## 🚨 Why This Matters in Production

- Data loss can happen instantly if containers are removed without proper volume usage.
- Misusing bind mounts can expose the host filesystem and create security risks.
- Incorrect volume strategy can severely impact performance and portability.

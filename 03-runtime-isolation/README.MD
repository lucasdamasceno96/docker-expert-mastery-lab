# Exercise 03: Linux Internals - Namespaces & Cgroups

## 📝 Overview

This exercise consisted of creating a **"manual container"** to validate that Docker is actually an abstraction layer built on top of three fundamental pillars of the Linux Kernel.

---

## 🏗️ The Three Pillars of Containerization

| Primitive               | Main Function            | What problem does it solve?                                                               |
| ----------------------- | ------------------------ | ----------------------------------------------------------------------------------------- |
| **Namespaces**          | **View Isolation**       | Ensures that a process cannot see the PIDs, networks, or hostnames of other processes.    |
| **Cgroups**             | **Resource Isolation**   | Limits how much CPU, memory, and I/O a process or group of processes can consume.         |
| **Chroot / Pivot_root** | **Filesystem Isolation** | Changes the process root directory (`/`), creating the illusion of a separate filesystem. |

---

## 🛠️ Implementation Details

In the `manual_isolation.sh` script, we used:

- **`unshare`**: Command used to create new namespaces.

- `--uts`: Isolates the hostname.

- `--pid`: Isolates the process tree (the shell becomes PID 1).

- `--fork`: Required so the new PID 1 process is properly managed.

- `--mount`: Creates a new mount namespace to isolate `/proc`.

- **Cgroups v2**: We directly manipulated the directory `/sys/fs/cgroup/` to define `memory.max`.

- **Procfs Mounting**: We learned that the `ps aux` command reads data from `/proc`. For full isolation, we must remount `/proc` inside the new namespace.

---

## 🎓 Specialist Takeaways

1. **The Container Illusion:** Containers are not “lightweight virtual machines”; they are simply processes with restrictions enforced by the Linux Kernel.

2. **Troubleshooting:** If a Docker container behaves unexpectedly, a specialist knows to inspect `/proc/[PID]/ns` and `/sys/fs/cgroup` on the host to diagnose the issue.

3. **Security:** Understanding namespaces is critical to understanding _Container Escape_ attacks. If the **user namespace** is not enabled, `root` inside the container is the same `root` on the host.

# Exercise 12: Performance Profiling & Resource Constraints

## 🎯 Objectives

- Differentiate between **Hard Limits** (enforced by Kernel) and **Soft Limits** (Reservations).
- Observe the **OOM (Out of Memory) Killer** in action.
- Master real-time monitoring using `docker stats` and `inspect`.

---

## 🛠️ The Stress Lab Configuration

In this exercise, we use `stress-ng` to intentionally exceed the memory limits defined in our `docker-compose.yml`.

### Resource Definition:

```yaml
deploy:
  resources:
    limits:
      memory: 50M # Hard Limit: Kernel will kill the process if exceeded
    reservations:
      memory: 20M # Soft Limit: Guaranteed memory for the container
```

### The Stress Command:

`--vm 1 --vm-bytes 100M`
We are forcing the container to allocate **100MB** of RAM, which is **double** its allowed limit.

---

## 🔍 Deep Analysis: The Death of a Container

When running this lab, we observed the following state via `docker inspect`:
`OOMKilled: true, ExitCode: 0` (or 137)

### 👨‍💻 Specialist Breakdown:

1. **OOMKilled: true**: This confirms the Linux Kernel (via Cgroups) intervened. When the container tried to hit 50.1MB, the Kernel identified a violation and sent a termination signal.
2. **Exit Code 137 vs 0**:
   - **137 (128+9):** The standard code for a `SIGKILL` sent by the OOM Killer.
   - **0:** If you see 0, it means the process managed to handle the signal or the container runtime wrapped the exit gracefully. **In production, always treat OOMKilled: true as a critical failure, regardless of the Exit Code.**

---

## 📊 Monitoring Tools

| Tool             | Usage                              | Specialist Context                                               |
| :--------------- | :--------------------------------- | :--------------------------------------------------------------- |
| `docker stats`   | Live streaming of CPU/MEM/Network. | Quick troubleshooting of a "noisy neighbor" container.           |
| `docker inspect` | Detailed state and error codes.    | Auditing why a container restarted in the middle of the night.   |
| **cAdvisor**     | Background metric collection.      | The standard for exporting container data to Prometheus/Grafana. |

---

## 💡 Specialist Takeaways

1. **The "Silent Killer":** Memory limits are "Hard". If you hit them, you die (OOM). CPU limits are "Soft" by default (Throttling); the container doesn't die, it just gets extremely slow.
2. **Swap Management:** By default, Docker allows a container to use swap memory if the host has it. To be a true specialist, limit swap using `--memory-swap` to ensure predictable performance.
3. **Avoid "Unlimited":** Never run containers in production without memory limits. One memory leak in a single container can trigger the OOM Killer to kill the **Docker Daemon** or the **Database** to save the Host.

# Exercise 04: Understanding the Container Lifecycle

This laboratory explores how the Linux Kernel manages process life cycles within a container, focusing on Signal Handling (SIGTERM/SIGKILL), the responsibilities of PID 1, and the creation of Zombie processes.

## 🎯 Learning Objectives

- Differentiate between `SIGTERM` (graceful) and `SIGKILL` (abrupt) signals.
- Understand the "PID 1 Problem" and why containers might ignore signals.
- Identify and simulate Zombie processes (`<defunct>`).
- Compare **Shell form** vs. **Exec form** in Dockerfiles.

## 🧬 Core Concepts

### 1. The PID 1 Responsibility

In Linux, the first process (PID 1) has two unique duties:

1. **Signal Forwarding:** It must pass signals from the Kernel (like `SIGTERM`) to its child processes.
2. **Reaping Orphans:** It must "wait" for child processes that have exited to prevent them from becoming Zombies.

### 2. Signal Handling

- **SIGTERM (15):** A request for the process to terminate gracefully. Our `signal_handler.sh` traps this to perform cleanup.
- **SIGKILL (9):** An immediate termination by the Kernel. No cleanup possible.

### 3. Zombie Processes (`<defunct>`)

A zombie is a process that has completed execution but still has an entry in the process table. This happens when the parent process fails to call `wait()`.

## 🛠️ Experiments in this Lab

### [A] Signal Handling Test

**Script:** `scripts/signal_handler.sh`

- **Goal:** Observe how a process can catch a signal and perform a "Graceful Shutdown" before exiting.
- **Key Command:** `docker stop <container>` (sends SIGTERM).

### [B] Zombie Creation Test

**Script:** `scripts/zombie_maker.py`

- **Goal:** Force the creation of a `<defunct>` process by making a parent ignore its child's exit status.
- **Key Observation:** Using `ps aux` to find the `Z` status.

## 🧠 Senior Takeaways

- Always use **Exec form** (`ENTRYPOINT ["/app"]`) in production to ensure signals reach your application.
- If your application spawns multiple sub-processes, consider using a minimalist init system like `--init` (Tini) to manage them properly.

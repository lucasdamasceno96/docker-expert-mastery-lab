# Exercise 09: Hardening - Linux Capabilities & Security Profiles

## 🎯 Objectives

- Understand and implement the **Principle of Least Privilege**.
- Master the use of `cap-drop` and `cap-add` to restrict container interaction with the Linux Kernel.
- Compare default security vs. hardened runtime environments.

## 🛠️ The Hardening Lab

We use a Docker Compose setup to demonstrate how to "lock down" a container while keeping it functional for its specific purpose.

### 1. The "Default" Risk

By default, Docker containers run with a subset of root capabilities (like `NET_RAW` for pinging). In a high-security environment, we start by dropping **everything**.

### 2. The "Drop-All" Strategy

In our `docker-compose.yml`, we use:

- `cap_drop: [ALL]`: Removes every single privilege.
- `cap_add: [NET_RAW]`: Specifically adds back only what is needed (e.g., for healthchecks or pings).

---

## 🔬 Practical Test (The Specialist Way)

1. **Attempting to change System Clock (Denied):**

   ```bash
   docker exec -it secure-app date -s "12:00:00"
   # Output: date: can't set date: Operation not permitted
   ```

2. **Verifying minimal networking (Allowed):**
   ```bash
   docker exec -it secure-app ping -c 1 8.8.8.8
   # Output: 1 packets transmitted, 1 received...
   ```

## 💡 Specialist Takeaways

1. **Attack Surface Reduction:** By dropping `CAP_SYS_ADMIN`, you prevent 90% of common container breakout exploits.
2. **Security-First Compose:** In production, every service should ideally have `cap_drop: [ALL]` as a starting point.
3. **No-New-Privileges:** Use `security_opt: [no-new-privileges:true]` to prevent processes from gaining more privileges than their parent.

## 🚨 Why This Matters in Production

- Containers running with excessive capabilities can lead to **container escape**.
- Attackers can abuse Linux capabilities to interact with the host kernel.
- Most applications require only a **minimal subset of privileges**.

👉 Real-world risk: A compromised container with `CAP_SYS_ADMIN` is almost equivalent to root on the host.

# Exercise 06: Network Troubleshooting (The Specialist Approach)

## 🎯 Overview

In production environments, security best practices dictate that containers should be **minimal** (Distroless or Slim). This means they lack tools like `ping`, `curl`, `tcpdump`, or `ip`.

This exercise demonstrates how to troubleshoot network issues from the **Host perspective** by entering the container's Network Namespace without installing any software inside the container.

---

## 🛠️ The "Ninja" Workflow

### 1. Identify the Target Process

Every Docker container is a process on the Host. To "enter" its network, we first need its **PID (Process ID)**.

```bash
# Get the PID of a running container
export TARGET_PID=$(docker inspect -f '{{.State.Pid}}' api-server)
echo "Target PID: $TARGET_PID"
```

### 2. Enter the Network Namespace

We use `nsenter` (namespace enter) to run host commands inside the container's network stack.

```bash
# View the container's internal interfaces from the Host
sudo nsenter -t $TARGET_PID -n ip addr
```

### 3. Sniffing Traffic (The Secret Weapon)

Since the container has no `tcpdump`, we use the Host's version to listen to the container's virtual interface.

```bash
# Run this on the Host to capture container traffic
sudo nsenter -t $TARGET_PID -n tcpdump -i eth0 -n icmp
```

---

## 🧠 What We Learned

1. **Namespace Isolation:** Containers don't have their own "OS"; they have isolated namespaces. The Network Namespace (`netns`) governs interfaces, routing tables, and iptables rules.
2. **Zero-Intrusion Debugging:** We can debug a container using 100% host-side tools. This keeps the image size small and the attack surface low.
3. **PID Mapping:** Understanding that `docker inspect` is the bridge between the Docker Engine abstraction and the Linux Kernel reality.

---

## 🚀 Real-World Scenarios (When to use this)

### 1. "It works on my machine, but not in the container"

When a container cannot reach an external API or database. Use `nsenter + tcpdump` to see if the SYN packets are leaving the container or if they are being dropped by a firewall/security group.

### 2. Debugging Latency

If an application is slow, you can capture a `.pcap` file using `nsenter` and analyze it in **Wireshark** to find TCP retransmissions or DNS resolution delays.

> `sudo nsenter -t $PID -n tcpdump -i eth0 -w debug.pcap`

### 3. Validating Service Mesh/Sidecars

In complex environments (like Istio or Linkerd), traffic is often redirected via `iptables`. Using `nsenter -n -t $PID iptables -L -t nat` allows you to see the hidden redirection rules inside the container.

### 4. Security Auditing

Verifying that a "locked-down" container truly cannot reach unauthorized internal subnets by attempting to sniff outgoing traffic during a penetration test.

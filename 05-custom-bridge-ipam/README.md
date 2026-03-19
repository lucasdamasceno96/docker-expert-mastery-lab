# Exercise 05: Custom Bridge Networks, DNS Resolution, and IPAM

## 🎯 Objective

Master the creation of isolated networks with granular addressing control (IPAM) and understand Docker’s native name resolution.

## 🛠️ Key Commands

- Create with Subnet and IP Range:
  `docker network create --driver bridge --subnet 10.200.0.0/24 --ip-range 10.200.0.128/25 --gateway 10.200.0.1 lab-secure-net`

- Container with Static IP:
  `docker run -d --name db-server --network lab-secure-net --ip 10.200.0.50 alpine sleep 3600`

## 💡 Specialist Takeaways

1. **Native DNS:** Only custom networks (User-Defined Bridges) have automatic name resolution. On the default `bridge` network, you must use IPs or `--link` (legacy).
2. **Layer 2 Isolation:** Containers on different networks cannot communicate, even on the same host, unless a container is connected to both.
3. **IPAM & Conflicts:** The error “Address already in use” is common when Docker subnets overlap with VPN routes or local networks (LAN). Planning IP ranges is part of infrastructure design.
4. **NAT/Masquerade:** External access for containers is handled via IPTables on the host, masking traffic from the internal network.

## 🔍 Inspection Evidence (Host)

- Bridge Interface: `brctl show` displays the creation of a `br-XXXX` bridge.
- NAT Rules: `sudo iptables -t nat -L -n` confirms MASQUERADE for the range 10.200.0.0/24.

# Exercise 10: Supply Chain Security & Vulnerability Scanning

## 🎯 Objectives

- Implement automated security scanning for Docker images.
- Understand and generate a **Software Bill of Materials (SBOM)**.
- Analyze real-world CVEs and prioritize remediation.

---

## 🛡️ What is Supply Chain Security?

Securing your runtime is not enough if your base image is already compromised. Supply Chain Security ensures that:

- Base images are free from known vulnerabilities (CVEs).
- Dependencies are tracked and auditable via **SBOM**.
- Security checks are integrated into the **CI/CD pipeline**.

---

## 🛠️ Tool: Trivy

Trivy is the industry-standard scanner. We use it to detect:

- **OS Vulnerabilities:** Alpine, Debian, etc.
- **Language Dependencies:** npm, pip, go modules.
- **Secrets:** Accidental private keys or tokens in layers.

---

## 🔍 Scanning in Action

### 1. Simple Vulnerability Scan

```bash
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
  aquasec/trivy:latest image alpine:latest
```

### 2. Generating an SBOM (CycloneDX Format)

An SBOM is a "bill of materials" for your container. It's essential for compliance (SOC2/ISO27001).

```bash
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
  aquasec/trivy:latest image --format cyclonedx --output sbom.json alpine:latest
```

---

## 🧪 Real-World Case Study: Analyzing `CVE-2026-22184`

During this lab, we identified a real vulnerability in the `zlib` package of an Alpine image:

- **Vulnerability:** `CVE-2026-22184`
- **Severity:** **HIGH** (Score: 8.6)
- **Package:** `zlib@1.3.1-r2`
- **Description:** Global buffer overflow in the `untgz` utility.
- **Remediation:** Upgrade `zlib` to version `1.3.2-r0`.

### 👨‍💻 Specialist Analysis

As a Senior Engineer, you must evaluate the **reachability** of the flaw. In this case, the vulnerability is in the `untgz` demonstration utility, not the core library.

- **Decision:** If your app doesn't call `untgz`, the risk is lower, but the **Best Practice** is still to update the base image or run `apk upgrade` in the Dockerfile.

---

## 💡 Specialist Takeaways

1. **Shift Left:** Failure to scan images in CI means you are deploying "blind".
2. **Ignore Unfixed:** Use `--ignore-unfixed` in automated pipelines to reduce noise from CVEs that have no available patch yet.
3. **PURL (Package URL):** Notice the `pkg:apk/alpine/zlib` format in the JSON. This is a standard way to identify software components across different tools.

---

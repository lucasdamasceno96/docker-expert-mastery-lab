# Exercise 07: Deep Dive into overlay2 and Storage Performance

## 🎯 Overview

Understand how Docker manages the filesystem using the `overlay2` storage driver and analyze the performance implications of the **Copy-on-Write (CoW)** strategy.

## 🔍 Key Concepts

- **LowerDir:** Read-only layers from the image.
- **UpperDir:** The "Diff" layer. Everything written or modified during container runtime lives here.
- **MergedDir:** The unified view of all layers. This is what the process inside the container sees.
- **Copy-on-Write (CoW):** When modifying an existing file from a lower layer, Docker copies it to the upper layer first. This causes I/O overhead for large files.

## 🛠️ Hands-on: Inspecting Layers

To find where your container's data actually lives on the host:

```bash
docker inspect <container_id> | jq '.[0].GraphDriver.Data'
```

## ⚡ Performance Specialist Takeaways

1. **Avoid Writing in Layers:** Heavy I/O (Databases, Logs, Temp files) should **never** happen in the container's writable layer. Use **Volumes** or **Bind Mounts** for native performance.
2. **Layer Deletion Myth:** Deleting a file in a new Dockerfile `RUN` command doesn't reduce image size if the file was created in a previous layer. It only adds a "whiteout" marker in the `UpperDir`.
3. **Storage Driver Choice:** While `overlay2` is the default and most efficient for most Linux distros, understanding its structure is crucial for debugging "Disk Full" issues where `/var/lib/docker/overlay2` consumes all space.

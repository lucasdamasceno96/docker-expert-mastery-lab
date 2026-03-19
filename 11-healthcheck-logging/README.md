# Exercise 11: Healthcheck Strategies & Logging Drivers

## 🎯 Objectives

- Implement advanced **Healthchecks** to monitor application liveness.
- Configure **Logging Drivers** to prevent disk exhaustion (Log Rotation).
- Understand the "Unhealthy" state and auto-recovery.

## 🔍 Key Concepts

### Healthcheck

A container can be `Up` but useless. The `HEALTHCHECK` instruction tells Docker how to test the application.

- `interval`: How often to run the test.
- `retries`: Number of consecutive failures before the container is considered `unhealthy`.
- `start_period`: Grace period for the app to initialize.

### Log Rotation

By default, Docker logs can grow indefinitely. In production, always limit log size:

```yaml
logging:
  driver: "json-file"
  options:
    max-size: "10m"
    max-file: "3"
```

## 🛠️ How to Test

1. Start the stack: `docker compose up -d`
2. Monitor health: `docker ps` (Wait for the status to change from `health: starting` to `healthy`).
3. **Simulate Failure:** Enter the container and stop Nginx:
   `docker exec -it web-api nginx -s stop`
4. Observe: Run `docker ps` again. After 30s (3 retries), the status will become `unhealthy`.

## 💡 Specialist Takeaways

1. **Never use generic healthchecks:** Don't just check if the port is open. Check a `/health` endpoint that validates DB connections and cache.
2. **External Logging:** In massive scales, use drivers like `fluentd`, `gelf` (Graylog), or `splunk` to send logs to a central server instead of storing them locally.

## 🧪 Proof of Concept: Simulating Failure

1. **Stop the service:** `docker exec -it web-api nginx -s stop`
2. **Detection:** Docker monitored the failure via the `HEALTHCHECK` instruction.
3. **Recovery:** Due to `restart: always` and the `unhealthy` state, the container was automatically restarted by the Docker Engine.

## 📊 Results from `docker inspect`

The health logs showed the following failure sequence:

- Exit code 7 (Failed to connect to localhost)
- 3 consecutive failures reached.
- Status transitioned to `unhealthy`.

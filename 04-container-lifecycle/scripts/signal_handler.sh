#!/bin/bash

# A função que será executada quando o sinal chegar
cleanup() {
    echo ""
    echo "[SIGNAL] SIGTERM received!"
    echo "[CLEANUP] Closing database connections..."
    sleep 2
    echo "[CLEANUP] Done. Exiting gracefully."
    exit 0
}

# O comando 'trap' captura o sinal (SIGTERM) e chama a função cleanup
trap 'cleanup' SIGTERM

echo "[INIT] Process started with PID: $$"
echo "[INFO] Waiting for SIGTERM... (Run 'docker stop' or 'kill -15 $$')"

# Loop infinito para manter o processo vivo
while true; do
    sleep 1 & 
    wait $!
done
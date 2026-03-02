# syntax=docker/dockerfile:1
FROM alpine:3.21

# Using a more robust shell configuration
SHELL ["/bin/sh", "-e", "-u", "-o", "pipefail", "-c"]

# 1. Install dependencies
RUN apk add --no-cache git openssh-client

# 2. Setup SSH known_hosts for GitHub
RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

# 3. USE THE SECRET (The "Zero-Leak" method)
RUN --mount=type=secret,id=api_token \
    echo "Using secret to configure build..." && \
    # Logic: Read secret and do something (simulated)
    TOKEN_VALUE=$(cat /run/secrets/api_token) && \
    echo "Secret size is: ${#TOKEN_VALUE} characters."

# 4. USE THE SSH MOUNT
RUN --mount=type=ssh \
    git clone git@github.com:lucasdamasceno96/docker-expert-mastery-lab.git /tmp/internal-lib

# 5. Final validation
RUN ls -la /tmp/internal-lib && \
    (ls /run/secrets/api_token 2>/dev/null || echo "Success: Secret is unmounted")